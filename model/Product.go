package model

import (
	"business-system_golang/utils/msg"

	"gorm.io/gorm"
)

// 产品 Model
type Product struct {
	gorm.Model
	Name           string `gorm:"type:varchar(20);comment:名称;not null" json:"name"`
	Brand          string `gorm:"type:varchar(20);comment:品牌" json:"brand"`
	Specification  string `gorm:"type:varchar(50);comment:规格" json:"specification"`
	SupplierID     uint   `gorm:"type:int;comment:供应商ID;default:(-)" json:"supplierID"`
	Number         int    `gorm:"type:int;comment:数量" json:"number"`
	Unit           string `gorm:"type:varchar(20);comment:单位" json:"unit"`
	PurchasedPrice int    `gorm:"type:int;comment:采购价格(元)" json:"purchasedPrice"`
	StandardPrice  int    `gorm:"type:int;comment:标准价格(元)" json:"standardPrice"`
	DeliveryCycle  string `gorm:"type:varchar(20);comment:供货周期" json:"deliveryCycle"`
	Remarks        string `gorm:"type:varchar(200);comment:备注" json:"remarks"`
	SourceTypeID   uint   `gorm:"type:int;comment:来源类型;default:(-)" json:"sourceTypeID"`
	SubtypeID      uint   `gorm:"type:int;comment:子类型;default:(-)" json:"subtypeID"`

	Supplier   Supplier   `gorm:"foreignKey:SupplierID" json:"supplier"`
	SourceType Dictionary `gorm:"foreignKey:SourceTypeID" json:"sourceType"`
	Subtype    Dictionary `gorm:"foreignKey:SubtypeID" json:"subtype"`
}

func CreateProduct(product *Product) (code int) {
	err = db.Create(&product).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func DeleteProduct(id int) (code int) {
	err = db.Where("id = ?", id).Delete(&Product{}).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func UpdateProduct(product *Product) (code int) {
	var maps = make(map[string]interface{})
	maps["Remarks"] = product.Remarks
	err = db.Model(&product).Updates(maps).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func SelectProduct(id int) (product Product, code int) {
	err = db.Preload("Supplier").Preload("SourceType").Preload("Subtype").First(&product, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return product, msg.ERROR_PRODUCT_NOT_EXIST
		} else {
			return product, msg.ERROR
		}
	}
	return product, msg.SUCCESS
}

func SelectProducts(pageSize int, pageNo int, productQuery ProductQuery) (products []Product, code int, total int64) {
	var maps = make(map[string]interface{})
	if productQuery.SourceTypeID != 0 {
		maps["source_type_id"] = productQuery.SourceTypeID
	}
	if productQuery.SubtypeID != 0 {
		maps["subtype_id"] = productQuery.SubtypeID
	}

	err = db.Limit(pageSize).Offset((pageNo-1)*pageSize).Where(maps).Where("name LIKE ? AND specification LIKE ?", "%"+productQuery.Name+"%", "%"+productQuery.Specification+"%").Find(&products).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, msg.ERROR, 0
	}
	db.Model(&products).Where(maps).Where("name LIKE ? AND specification LIKE ?", "%"+productQuery.Name+"%", "%"+productQuery.Specification+"%").Count(&total)
	return products, msg.SUCCESS, total
}
