package model

import (
	"business-system_golang/utils/msg"
	"business-system_golang/utils/uid"

	"gorm.io/gorm"
)

// 供应商 Model
type Supplier struct {
	BaseModel
	UID         string `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	Name        string `gorm:"type:varchar(100);comment:名称;not null" json:"name"`
	Address     string `gorm:"type:varchar(200);comment:地址" json:"address"`
	Web         string `gorm:"type:varchar(100);comment:网站" json:"web"`
	Linkman     string `gorm:"type:varchar(50);comment:联系人" json:"linkman"`
	Phone       string `gorm:"type:varchar(100);comment:联系电话" json:"phone"`
	WechatID    string `gorm:"type:varchar(50);comment:微信号" json:"wechatID"`
	Email       string `gorm:"type:varchar(50);comment:邮箱" json:"email"`
	Description string `gorm:"type:varchar(600);comment:主营产品概述" json:"description"`
	Remark      string `gorm:"type:varchar(600);comment:备注" json:"remark"`
	IsDelete    bool   `gorm:"type:boolean;comment:是否删除" json:"isDelete"`
}

type SupplierQuery struct {
	Name    string `json:"name"`
	Linkman string `json:"linkman"`
	Phone   string `json:"phone"`
}

func InsertSupplier(supplier *Supplier) (code int) {
	supplier.UID = uid.Generate()
	err = db.Create(&supplier).Error
	if err != nil {
		return msg.ERROR_SUPPLIER_INSERT
	}
	return msg.SUCCESS
}

func DeleteSupplier(uid string) (code int) {
	// err = db.Where("uid = ?", uid).Delete(&Supplier{}).Error
	err = db.Model(&Supplier{}).Where("uid = ?", uid).Update("is_delete", true).Error
	if err != nil {
		return msg.ERROR_SUPPLIER_DELETE
	}
	return msg.SUCCESS
}

func UpdateSupplier(supplier *Supplier) (code int) {
	var maps = make(map[string]interface{})
	maps["name"] = supplier.Name
	maps["web"] = supplier.Web
	maps["address"] = supplier.Address
	maps["linkman"] = supplier.Linkman
	maps["phone"] = supplier.Phone
	maps["wechat_id"] = supplier.WechatID
	maps["email"] = supplier.Email

	err = db.Model(&Supplier{}).Where("uid = ?", supplier.UID).Updates(supplier).Error

	if err != nil {
		return msg.ERROR_SUPPLIER_UPDATE
	}
	return msg.SUCCESS
}

func SelectSupplier(uid string) (supplier Supplier, code int) {
	err = db.Where("is_delete = ?", false).First(&supplier, "uid = ?", uid).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return supplier, msg.ERROR_SUPPLIER_NOT_EXIST
		} else {
			return supplier, msg.ERROR_SUPPLIER_SELECT
		}
	}
	return supplier, msg.SUCCESS
}

func SelectSuppliers(pageSize int, pageNo int, supplierQuery *SupplierQuery) (suppliers []Supplier, code int, total int64) {

	tDb := db.Where("is_delete = ?", false)

	if supplierQuery.Name != "" {
		tDb = tDb.Where("name LIKE ?", "%"+supplierQuery.Name+"%")
	}
	if supplierQuery.Linkman != "" {
		tDb = tDb.Where("linkman LIKE ?", "%"+supplierQuery.Linkman+"%")
	}
	if supplierQuery.Phone != "" {
		tDb = tDb.Where("phone LIKE ?", "%"+supplierQuery.Phone+"%")
	}

	err = tDb.Find(&suppliers).Count(&total).
		Limit(pageSize).Offset((pageNo - 1) * pageSize).
		Find(&suppliers).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return suppliers, msg.ERROR_SUPPLIER_SELECT, total
	}
	return suppliers, msg.SUCCESS, total
}
