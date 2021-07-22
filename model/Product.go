package model

import "gorm.io/gorm"

// 产品 Model
type Product struct {
	gorm.Model
	Name           string `gorm:"type:varchar(20);comment:名称;not null" json:"name"`
	Brand          string `gorm:"type:varchar(20);comment:品牌;not null" json:"brand"`
	Specification  string `gorm:"type:varchar(50);comment:规格;not null" json:"specification"`
	SupplierId     int    `gorm:"type:int;comment:供应商ID;not null" json:"supplierId"`
	Number         int    `gorm:"type:int;comment:数量;not null" json:"number"`
	Unit           string `gorm:"type:varchar(20);comment:单位;not null" json:"unit"`
	PurchasedPrice int    `gorm:"type:int;comment:采购价格(元);not null" json:"purchasedPrice"`
	StandardPrice  int    `gorm:"type:int;comment:标准价格(元);not null" json:"standardPrice"`
	DeliveryCycle  string `gorm:"type:varchar(20);comment:供货周期;not null" json:"deliveryCycle"`
	Remarks        string `gorm:"type:varchar(200);comment:备注" json:"remarks"`
}
