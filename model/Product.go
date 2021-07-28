package model

import (
	"business-system_golang/utils/msg"

	"gorm.io/gorm"
)

// 产品 Model
type Product struct {
	gorm.Model
	Name           string `gorm:"type:varchar(20);comment:名称;not null" json:"name"`
	Brand          string `gorm:"type:varchar(20);comment:品牌;not null" json:"brand"`
	Specification  string `gorm:"type:varchar(50);comment:规格;not null" json:"specification"`
	SupplierId     uint   `gorm:"type:int;comment:供应商ID;not null" json:"supplierId"`
	Number         int    `gorm:"type:int;comment:数量;not null" json:"number"`
	Unit           string `gorm:"type:varchar(20);comment:单位;not null" json:"unit"`
	PurchasedPrice int    `gorm:"type:int;comment:采购价格(元);not null" json:"purchasedPrice"`
	StandardPrice  int    `gorm:"type:int;comment:标准价格(元);not null" json:"standardPrice"`
	DeliveryCycle  string `gorm:"type:varchar(20);comment:供货周期;not null" json:"deliveryCycle"`
	Remarks        string `gorm:"type:varchar(200);comment:备注" json:"remarks"`
	SourceType     uint   `gorm:"type:int;comment:来源类型;not null" json:"sourceType"`
	Subtype        uint   `gorm:"type:int;comment:子类型;not null" json:"subtype"`
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
	err = db.First(&product, id).Error
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
	err = db.Limit(pageSize).Offset((pageNo-1)*pageSize).Where("source_type = ? AND subtype = ? AND name LIKE ? AND specification LIKE ?", productQuery.SourceType, productQuery.Subtype, "%"+productQuery.Name+"%", "%"+productQuery.Specification+"%").Find(&products).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, msg.ERROR, 0
	}
	db.Model(&products).Where("source_type = ? AND subtype = ? AND name LIKE ? AND specification LIKE ?", productQuery.SourceType, productQuery.Subtype, "%"+productQuery.Name+"%", "%"+productQuery.Specification+"%").Count(&total)
	return products, msg.SUCCESS, total

}
