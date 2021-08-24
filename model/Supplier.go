package model

import (
	"business-system_golang/utils/msg"

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

func CreateSupplier(supplier *Supplier) (code int) {
	err = db.Create(&supplier).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func DeleteSupplier(id int) (code int) {
	err = db.Where("id = ?", id).Delete(&Supplier{}).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func UpdateSupplier(supplier *Supplier) (code int) {
	err = db.Omit("name").Updates(&supplier).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func SelectSupplier(id int) (supplier Supplier, code int) {
	err = db.First(&supplier, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return supplier, msg.ERROR_SUPPLIER_NOT_EXIST
		} else {
			return supplier, msg.ERROR
		}
	}
	return supplier, msg.SUCCESS
}

func SelectSuppliers(pageSize int, pageNo int, supplierQuery SupplierQuery) (suppliers []Supplier, code int, total int64) {
	err = db.Limit(pageSize).Offset((pageNo-1)*pageSize).Where("name LIKE ? AND linkman LIKE ? AND phone LIKE ?", "%"+supplierQuery.Name+"%", "%"+supplierQuery.Linkman+"%", "%"+supplierQuery.Phone+"%").Find(&suppliers).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, msg.ERROR, 0
	}
	db.Model(&suppliers).Where("name LIKE ? AND linkman LIKE ? AND phone LIKE ?", "%"+supplierQuery.Name+"%", "%"+supplierQuery.Linkman+"%", "%"+supplierQuery.Phone+"%").Count(&total)
	return suppliers, msg.SUCCESS, total
}
