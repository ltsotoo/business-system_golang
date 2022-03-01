package model

import (
	"business-system_golang/utils/msg"
	"business-system_golang/utils/uid"

	"gorm.io/gorm"
)

// 产品 Model
type Product struct {
	BaseModel
	UID              string  `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	Name             string  `gorm:"type:varchar(100);comment:名称;not null" json:"name"`
	Brand            string  `gorm:"type:varchar(100);comment:品牌" json:"brand"`
	Specification    string  `gorm:"type:varchar(200);comment:规格" json:"specification"`
	SupplierUID      string  `gorm:"type:varchar(32);comment:供应商ID;default:(-)" json:"supplierUID"`
	Number           int     `gorm:"type:int;comment:可售数量(库存数量-订单锁定但未出库的数量)" json:"number"`
	NumberCount      int     `gorm:"type:int;comment:库存数量" json:"numberCount"`
	Unit             string  `gorm:"type:varchar(50);comment:单位" json:"unit"`
	StandardPrice    float64 `gorm:"type:decimal(20,6);comment:标准价格(元)" json:"standardPrice"`
	StandardPriceUSD float64 `gorm:"type:decimal(20,6);comment:标准价格(美元)" json:"standardPriceUSD"`
	DeliveryCycle    string  `gorm:"type:varchar(50);comment:供货周期" json:"deliveryCycle"`
	Remarks          string  `gorm:"type:varchar(600);comment:备注" json:"remarks"`
	TypeUID          string  `gorm:"type:varchar(32);comment:类型;default:(-)" json:"typeUID"`
	IsDelete         bool    `gorm:"type:boolean;comment:是否删除" json:"isDelete"`

	Type     ProductType `gorm:"foreignKey:TypeUID;references:UID" json:"type"`
	Supplier Supplier    `gorm:"foreignKey:SupplierUID;references:UID" json:"supplier"`
}

type ProductType struct {
	BaseModel
	UID                        string  `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	Name                       string  `gorm:"type:varchar(20);comment:名称;not null" json:"name"`
	PushMoneyPercentages       float64 `gorm:"type:decimal(20,6);comment:标准提成百分比" json:"pushMoneyPercentages"`
	MinPushMoneyPercentages    float64 `gorm:"type:decimal(20,6);comment:标准提成百分比" json:"minPushMoneyPercentages"`
	PushMoneyPercentagesUp     float64 `gorm:"type:decimal(20,6);comment:提成上涨百分比" json:"pushMoneyPercentagesUp"`
	PushMoneyPercentagesDown   float64 `gorm:"type:decimal(20,6);comment:提成下降百分比" json:"pushMoneyPercentagesDown"`
	BusinessMoneyPercentages   float64 `gorm:"type:decimal(20,6);comment:标准业务费用百分比" json:"businessMoneyPercentages"`
	BusinessMoneyPercentagesUp float64 `gorm:"type:decimal(20,6);comment:业务费用上涨百分比" json:"businessMoneyPercentagesUp"`
	IsTaskLoad                 bool    `gorm:"type:boolean;comment:是否计算任务量" json:"isTaskLoad"`

	IsDelete bool `gorm:"type:boolean;comment:是否删除" json:"isDelete"`
}

type ProductQuery struct {
	TypeUID       string `json:"typeUID"`
	Name          string `json:"name"`
	Specification string `json:"specification"`
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
	err = db.Model(&Product{}).Where("uid = ?", uid).Update("is_delete", true).Error
	if err != nil {
		return msg.ERROR_PRODUCT_DELETE
	}
	return msg.SUCCESS
}

func UpdateProductBase(product *Product) (code int) {
	var maps = make(map[string]interface{})
	if product.TypeUID != "" {
		maps["type_uid"] = product.TypeUID
	} else {
		maps["type_uid"] = nil
	}
	if product.SupplierUID != "" {
		maps["supplier_uid"] = product.SupplierUID
	} else {
		maps["supplier_uid"] = nil
	}
	maps["name"] = product.Name
	maps["brand"] = product.Brand
	maps["specification"] = product.Specification
	maps["unit"] = product.Unit
	maps["delivery_cycle"] = product.DeliveryCycle
	maps["remarks"] = product.Remarks
	err = db.Model(&Product{}).Where("uid = ?", product.UID).Updates(maps).Error
	if err != nil {
		return msg.ERROR_PRODUCT_UPDATE
	}
	return msg.SUCCESS
}

func UpdateProductNumber(product *Product) (code int) {
	err = db.Exec("update product set number = number + ? , number_count = number_count + ? where uid = ?", product.Number, product.Number, product.UID).Error
	if err != nil {
		return msg.ERROR_PRODUCT_UPDATE
	}
	return msg.SUCCESS
}

func UpdateProductPrice(product *Product) (code int) {
	var maps = make(map[string]interface{})
	maps["standard_price"] = product.StandardPrice
	maps["standard_price_usd"] = product.StandardPriceUSD
	err = db.Model(&Product{}).Where("uid = ?", product.UID).Updates(maps).Error
	if err != nil {
		return msg.ERROR_PRODUCT_UPDATE
	}
	return msg.SUCCESS
}

func SelectProduct(uid string) (product Product, code int) {
	err = db.Preload("Supplier").Preload("Type").Where("is_delete = ?", false).First(&product, "uid = ?", uid).Error
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
	maps["is_delete"] = false
	if productQuery.TypeUID != "" {
		maps["type_uid"] = productQuery.TypeUID
	}

	tDb := db.Where(maps)

	if productQuery.Name != "" {
		tDb = tDb.Where("name LIKE ?", "%"+productQuery.Name+"%")
	}
	if productQuery.Specification != "" {
		tDb = tDb.Where("specification LIKE ?", "%"+productQuery.Specification+"%")
	}

	err = tDb.Find(&products).Count(&total).
		Preload("Supplier").Preload("Type").
		Limit(pageSize).Offset((pageNo - 1) * pageSize).
		Find(&products).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return products, msg.ERROR, 0
	}
	return products, msg.SUCCESS, total
}

func InsertProductType(productType *ProductType) (code int) {
	productType.UID = uid.Generate()
	err = db.Create(&productType).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func DeleteProductType(uid string) (code int) {
	// err = db.Where("uid = ?", uid).Delete(&ProductType{}).Error
	err = db.Model(&ProductType{}).Where("uid = ?", uid).Update("is_delete", true).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func UpdateProductType(productType *ProductType) (code int) {
	var maps = make(map[string]interface{})
	maps["name"] = productType.Name
	maps["push_money_percentages"] = productType.PushMoneyPercentages
	maps["min_push_money_percentages"] = productType.MinPushMoneyPercentages
	maps["push_money_percentages_up"] = productType.PushMoneyPercentagesUp
	maps["push_money_percentages_down"] = productType.PushMoneyPercentagesDown
	maps["business_money_percentages"] = productType.BusinessMoneyPercentages
	maps["business_money_percentages_up"] = productType.BusinessMoneyPercentagesUp
	maps["is_task_load"] = productType.IsTaskLoad
	err = db.Model(&ProductType{}).Where("uid = ?", productType.UID).Updates(maps).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func SelectProductTypes(productType *ProductType) (productTypes []ProductType, code int) {
	err = db.Where("is_delete = ? AND name LIKE ?", false, "%"+productType.Name+"%").Find(&productTypes).Error
	if err != nil {
		return productTypes, msg.ERROR
	}
	return productTypes, msg.SUCCESS
}
