package model

import (
	"business-system_golang/utils/msg"
	"business-system_golang/utils/uid"

	"gorm.io/gorm"
)

// 供应商 Model
type Supplier struct {
	BaseModel
	UID      string `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	Name     string `gorm:"type:varchar(50);comment:名称;not null" json:"name"`
	Address  string `gorm:"type:varchar(50);comment:地址;not null" json:"address"`
	Web      string `gorm:"type:varchar(50);comment:网站;not null" json:"web"`
	Linkman  string `gorm:"type:varchar(50);comment:联系人;not null" json:"linkman"`
	Phone    string `gorm:"type:varchar(50);comment:联系电话;not null" json:"phone"`
	WechatID string `gorm:"type:varchar(50);comment:微信号" json:"wechatID"`
	Email    string `gorm:"type:varchar(50);comment:邮箱" json:"email"`
	IsDelete bool   `gorm:"type:boolean;comment:是否删除" json:"isDelete"`
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
	maps["Linkman"] = supplier.Linkman
	maps["Phone"] = supplier.Phone
	maps["WechatID"] = supplier.WechatID
	maps["Email"] = supplier.Email
	maps["Web"] = supplier.Web

	err = db.Model(&Supplier{}).Where("uid = ?", supplier.UID).Updates(maps).Error

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
	err = db.Where("is_delete = ?", false).Where("name LIKE ? AND linkman LIKE ? AND phone LIKE ?", "%"+supplierQuery.Name+"%", "%"+supplierQuery.Linkman+"%", "%"+supplierQuery.Phone+"%").
		Find(&suppliers).Count(&total).
		Limit(pageSize).Offset((pageNo - 1) * pageSize).
		Find(&suppliers).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return suppliers, msg.ERROR_SUPPLIER_SELECT, total
	}
	return suppliers, msg.SUCCESS, total
}
