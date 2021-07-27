package model

import (
	"business-system_golang/utils/msg"

	"gorm.io/gorm"
)

// 客户 Model
type Customer struct {
	gorm.Model
	Name          string `gorm:"type:varchar(20);comment:姓名;not null" json:"name"`
	Address       string `gorm:"type:varchar(20);comment:地址;not null" json:"address"`
	Company       string `gorm:"type:varchar(20);comment:公司;not null" json:"company"`
	ResearchGroup string `gorm:"type:varchar(20);comment:课题组;not null" json:"researchGroup"`
	Phone         string `gorm:"type:varchar(20);comment:电话;not null" json:"phone"`
	WechatID      string `gorm:"type:varchar(20);comment:微信号" json:"wechatID"`
	Email         string `gorm:"type:varchar(20);comment:邮箱" json:"email"`
}

func CreateCustomer(customer *Customer) (code int) {
	err = db.Create(&customer).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func DeleteCustomer(id int) (code int) {
	err = db.Where("id = ?", id).Delete(&Customer{}).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func UpdateCustomer(customer *Customer) (code int) {
	var maps = make(map[string]interface{})
	maps["WechatID"] = customer.WechatID
	maps["Email"] = customer.Email
	err = db.Model(&customer).Updates(maps).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func SelectCustomer(id int) (customer Customer, code int) {
	err = db.First(&customer, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return customer, msg.ERROR_CUSTOMER_NOT_EXIST
		} else {
			return customer, msg.ERROR
		}
	}
	return customer, msg.SUCCESS
}

func SelectCustomers(pageSize int, pageNo int) (customers []Customer, code int, total int64) {
	err = db.Limit(pageSize).Offset((pageNo - 1) * pageSize).Find(&customers).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, msg.ERROR, 0
	}
	db.Model(&customers).Count(&total)
	return customers, msg.SUCCESS, total
}
