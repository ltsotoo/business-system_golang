package model

import (
	"business-system_golang/utils/msg"
	"business-system_golang/utils/uid"

	"gorm.io/gorm"
)

// 客户 Model
type Customer struct {
	gorm.Model
	UID           string `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	CompanyUID    string `gorm:"type:varchar(32);comment:公司UID;not null" json:"companyUID"`
	Name          string `gorm:"type:varchar(20);comment:姓名;not null" json:"name"`
	ResearchGroup string `gorm:"type:varchar(20);comment:课题组" json:"researchGroup"`
	Phone         string `gorm:"type:varchar(20);comment:电话" json:"phone"`
	WechatID      string `gorm:"type:varchar(20);comment:微信号" json:"wechatID"`
	Email         string `gorm:"type:varchar(20);comment:邮箱" json:"email"`
	Status        int    `gorm:"type:int;comment:状态(0:未审核,1:通过审核)" json:"status"`

	Company CustomerCompany `gorm:"foreignKey:CompanyUID;references:UID" json:"company"`
}

type CustomerQuery struct {
	AreaUID       string `json:"areaUID"`
	CompanyUID    string `json:"companyUID"`
	ResearchGroup string `json:"researchGroup"`
	Name          string `json:"name"`
}

// 客户公司 Model
type CustomerCompany struct {
	gorm.Model
	UID     string `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	AreaUID string `gorm:"type:varchar(32);comment:地区UID;not null" json:"areaUID"`
	Name    string `gorm:"type:varchar(20);comment:名称;not null" json:"name"`
	Address string `gorm:"type:varchar(20);comment:地址" json:"address"`

	Area Area `gorm:"foreignKey:AreaUID;references:UID" json:"area"`
}

func InsertCustomer(customer *Customer) (code int) {
	customer.UID = uid.Generate()
	err = db.Create(&customer).Error
	if err != nil {
		return msg.ERROR_CUSTOMER_INSERT
	}
	return msg.SUCCESS
}

func DeleteCustomer(uid string) (code int) {
	err = db.Where("uid = ?", uid).Delete(&Customer{}).Error
	if err != nil {
		return msg.ERROR_CUSTOMER_DELETE
	}
	return msg.SUCCESS
}

func UpdateCustomer(customer *Customer) (code int) {
	var maps = make(map[string]interface{})
	maps["ResearchGroup"] = customer.ResearchGroup
	maps["Phone"] = customer.Phone
	maps["WechatID"] = customer.WechatID
	maps["Email"] = customer.Email
	err = db.Model(&Customer{}).Where("uid = ?", customer.UID).Updates(maps).Error
	if err != nil {
		return msg.ERROR_CUSTOMER_UPDATE
	}
	return msg.SUCCESS
}

func SelectCustomer(uid string) (customer Customer, code int) {
	err = db.Preload("Company").First(&customer, "uid = ?", uid).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return customer, msg.ERROR_CUSTOMER_NOT_EXIST
		} else {
			return customer, msg.ERROR_CUSTOMER_SELECT
		}
	}
	return customer, msg.SUCCESS
}

func SelectCustomers(pageSize int, pageNo int, customerQuery *CustomerQuery) (customers []Customer, code int, total int64) {
	var maps = make(map[string]interface{})
	if customerQuery.CompanyUID != "" {
		maps["customer.company_uid"] = customerQuery.CompanyUID
	}
	if customerQuery.AreaUID != "" {
		maps["Company.area_uid"] = customerQuery.AreaUID
	}

	err = db.Joins("Company").Where(maps).
		Where("customer.research_group LIKE ? AND customer.name LIKE ?",
			"%"+customerQuery.ResearchGroup+"%", "%"+customerQuery.Name+"%").
		Find(&customers).Count(&total).
		Limit(pageSize).Offset((pageNo - 1) * pageSize).
		Find(&customers).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return customers, msg.ERROR_CUSTOMER_SELECT, total
	}
	return customers, msg.SUCCESS, total
}

func SelectCompanys(areaUID string) (CustomerCompanys []CustomerCompany, code int) {
	err = db.Find(&CustomerCompanys).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return CustomerCompanys, msg.ERROR
	}
	return CustomerCompanys, msg.SUCCESS
}
