package model

import (
	"business-system_golang/utils/msg"
	"business-system_golang/utils/uid"

	"gorm.io/gorm"
)

// 供应商 Model
type Supplier struct {
	gorm.Model
	UID      string `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	Name     string `gorm:"type:varchar(20);comment:名称;not null" json:"name"`
	Address  string `gorm:"type:varchar(20);comment:地址;not null" json:"address"`
	Linkman  string `gorm:"type:varchar(20);comment:联系人;not null" json:"linkman"`
	Phone    string `gorm:"type:varchar(20);comment:联系电话;not null" json:"phone"`
	WechatID string `gorm:"type:varchar(20);comment:微信号" json:"wechatID"`
	Email    string `gorm:"type:varchar(20);comment:邮箱" json:"email"`
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
	err = db.Where("uid = ?", uid).Delete(&Supplier{}).Error
	if err != nil {
		return msg.ERROR_SUPPLIER_DELETE
	}
	return msg.SUCCESS
}

func UpdateSupplier(supplier *Supplier) (code int) {
	var maps = make(map[string]interface{})
	maps["Linkman"] = supplier.Linkman
	maps["Phone"] = supplier.Phone
	maps["WechatID"] = supplier.WechatID
	maps["Email"] = supplier.Email

	err = db.Model(&supplier).Updates(maps).Error

	if err != nil {
		return msg.ERROR_SUPPLIER_UPDATE
	}
	return msg.SUCCESS
}

func SelectSupplier(uid string) (supplier Supplier, code int) {
	err = db.First(&supplier, "uid = ?", uid).Error
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
	err = db.Where("name LIKE ? AND linkman LIKE ? AND phone LIKE ?", "%"+supplierQuery.Name+"%", "%"+supplierQuery.Linkman+"%", "%"+supplierQuery.Phone+"%").
		Find(&suppliers).Count(&total).
		Limit(pageSize).Offset((pageNo - 1) * pageSize).
		Find(&suppliers).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return suppliers, msg.ERROR_SUPPLIER_SELECT, total
	}
	return suppliers, msg.SUCCESS, total
}
