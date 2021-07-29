package model

import (
	"business-system_golang/utils/msg"

	"gorm.io/gorm"
)

// 客户 Model
type Customer struct {
	gorm.Model
	Name            string `gorm:"type:varchar(20);comment:姓名;not null" json:"name"`
	CompanyID       uint   `gorm:"type:int;comment:公司ID;not null" json:"companyID"`
	ResearchGroupID uint   `gorm:"type:int;comment:课题组ID;not null" json:"researchGroupID"`
	Phone           string `gorm:"type:varchar(20);comment:电话;not null" json:"phone"`
	WechatID        string `gorm:"type:varchar(20);comment:微信号" json:"wechatID"`
	Email           string `gorm:"type:varchar(20);comment:邮箱" json:"email"`
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

func SelectCustomers(pageSize int, pageNo int, customerQuery CustomerQuery) (customers []Customer, code int, total int64) {

	var maps = make(map[string]interface{})
	if customerQuery.CompanyID != 0 {
		maps["company_id"] = customerQuery.CompanyID
	}
	if customerQuery.ResearchGroupID != 0 {
		maps["research_group_id"] = customerQuery.ResearchGroupID
	}

	err = db.Limit(pageSize).Offset((pageNo-1)*pageSize).Where(maps).Where("name LIKE ?", "%"+customerQuery.Name+"%").Find(&customers).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, msg.ERROR, 0
	}
	db.Model(&customers).Where(maps).Where("name LIKE ?", "%"+customerQuery.Name+"%").Count(&total)
	return customers, msg.SUCCESS, total
}
