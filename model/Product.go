package model

import (
	"business-system_golang/utils/msg"
	"business-system_golang/utils/uid"

	"gorm.io/gorm"
)

// 产品 Model
type Product struct {
	gorm.Model
	UID            string `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	Name           string `gorm:"type:varchar(20);comment:名称;not null" json:"name"`
	Brand          string `gorm:"type:varchar(20);comment:品牌" json:"brand"`
	Specification  string `gorm:"type:varchar(50);comment:规格" json:"specification"`
	SupplierUID    string `gorm:"type:varchar(32);comment:供应商ID;default:(-)" json:"supplierUID"`
	Number         int    `gorm:"type:int;comment:数量" json:"number"`
	Unit           string `gorm:"type:varchar(20);comment:单位" json:"unit"`
	PurchasedPrice int    `gorm:"type:int;comment:采购价格(元)" json:"purchasedPrice"`
	StandardPrice  int    `gorm:"type:int;comment:标准价格(元)" json:"standardPrice"`
	DeliveryCycle  string `gorm:"type:varchar(20);comment:供货周期" json:"deliveryCycle"`
	Remarks        string `gorm:"type:varchar(200);comment:备注" json:"remarks"`
	SourceTypeUID  string `gorm:"type:varchar(32);comment:来源类型;default:(-)" json:"sourceTypeUID"`
	SubtypeUID     string `gorm:"type:varchar(32);comment:子类型;default:(-)" json:"subtypeUID"`

	Supplier   Supplier   `gorm:"foreignKey:SupplierUID;references:UID" json:"supplier"`
	SourceType Dictionary `gorm:"foreignKey:SourceTypeUID;references:UID" json:"sourceType"`
	Subtype    Dictionary `gorm:"foreignKey:SubtypeUID;references:UID" json:"subtype"`
}

func InsertProduct(product *Product) (code int) {
	product.UID = uid.Generate()
	err = db.Create(&product).Error
	if err != nil {
		return msg.ERROR_PRODUCT_INSERT
	}
	return msg.SUCCESS
}

func DeleteProduct(uid string) (code int) {
	err = db.Where("uid = ?", uid).Delete(&Product{}).Error
	if err != nil {
		return msg.ERROR_PRODUCT_DELETE
	}
	return msg.SUCCESS
}

func UpdateProduct(product *Product) (code int) {
	var maps = make(map[string]interface{})
	maps["Remarks"] = product.Remarks
	err = db.Model(&product).Updates(maps).Error
	if err != nil {
		return msg.ERROR_PRODUCT_UPDATE
	}
	return msg.SUCCESS
}

func SelectProduct(uid string) (product Product, code int) {
	err = db.Preload("Supplier").Preload("SourceType").Preload("Subtype").
		First(&product, "uid = ?", uid).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return product, msg.ERROR_PRODUCT_NOT_EXIST
		} else {
			return product, msg.ERROR_PRODUCT_SELECT
		}
	}
	return product, msg.SUCCESS
}

func SelectProducts(pageSize int, pageNo int, productQuery *ProductQuery) (products []Product, code int, total int64) {
	var maps = make(map[string]interface{})
	if productQuery.SourceTypeUID != "" {
		maps["source_type_uid"] = productQuery.SourceTypeUID
	}
	if productQuery.SubtypeUID != "" {
		maps["subtype_uid"] = productQuery.SubtypeUID
	}

	err = db.Where(maps).Where("name LIKE ? AND specification LIKE ?", "%"+productQuery.Name+"%", "%"+productQuery.Specification+"%").
		Find(&products).Count(&total).
		Preload("Supplier").Preload("SourceType").Preload("Subtype").
		Limit(pageSize).Offset((pageNo - 1) * pageSize).Find(&products).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, msg.ERROR, 0
	}
	return products, msg.SUCCESS, total
}
