package model

import (
	"business-system_golang/utils/msg"
	"business-system_golang/utils/uid"

	"gorm.io/gorm"
)

// 客户 Model
type Customer struct {
	BaseModel
	UID           string `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	CompanyUID    string `gorm:"type:varchar(32);comment:公司UID;default:(-)" json:"companyUID"`
	Name          string `gorm:"type:varchar(50);comment:姓名;not null" json:"name"`
	ResearchGroup string `gorm:"type:varchar(100);comment:课题组" json:"researchGroup"`
	Phone         string `gorm:"type:varchar(100);comment:电话" json:"phone"`
	WechatID      string `gorm:"type:varchar(50);comment:微信号" json:"wechatID"`
	Email         string `gorm:"type:varchar(50);comment:邮箱" json:"email"`
	Status        int    `gorm:"type:int;comment:状态(0:未审核,1:通过审核)" json:"status"`
	IsDelete      bool   `gorm:"type:boolean;comment:是否删除" json:"isDelete"`

	Company CustomerCompany `gorm:"foreignKey:CompanyUID;references:UID" json:"company"`
}

type CustomerQuery struct {
	RegionUID     string `json:"regionUID"`
	CompanyUID    string `json:"companyUID"`
	ResearchGroup string `json:"researchGroup"`
	Name          string `json:"name"`
}

// 客户公司 Model
type CustomerCompany struct {
	BaseModel
	UID       string `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	RegionUID string `gorm:"type:varchar(32);comment:省份UID;default:(-)" json:"regionUID"`
	Name      string `gorm:"type:varchar(100);comment:名称;not null" json:"name"`
	Address   string `gorm:"type:varchar(200);comment:地址" json:"address"`
	IsDelete  bool   `gorm:"type:boolean;comment:是否删除" json:"isDelete"`

	Region Dictionary `gorm:"foreignKey:RegionUID;references:UID" json:"region"`
}

func InsertCustomer(customer *Customer) (code int) {
	customer.UID = uid.Generate()
	customer.Status = 1
	err = db.Create(&customer).Error
	if err != nil {
		return msg.ERROR_CUSTOMER_INSERT
	}
	return msg.SUCCESS
}

func DeleteCustomer(uid string) (code int) {
	err = db.Model(&Customer{}).Where("uid = ?", uid).Update("is_delete", true).Error
	if err != nil {
		return msg.ERROR_CUSTOMER_DELETE
	}
	return msg.SUCCESS
}

func UpdateCustomer(customer *Customer) (code int) {
	var maps = make(map[string]interface{})
	maps["name"] = customer.Name
	maps["research_group"] = customer.ResearchGroup
	maps["phone"] = customer.Phone
	maps["wechat_id"] = customer.WechatID
	maps["email"] = customer.Email
	err = db.Model(&Customer{}).Where("uid = ?", customer.UID).Updates(maps).Error
	if err != nil {
		return msg.ERROR_CUSTOMER_UPDATE
	}
	return msg.SUCCESS
}

func SelectCustomer(uid string) (customer Customer, code int) {
	err = db.Preload("Company").Where("is_delete = ? AND status = ?", false, 1).First(&customer, "uid = ?", uid).Error
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
	tDb := db.Joins("Company").Where("customer.is_delete = ? AND status = ?", false, 1)

	if customerQuery.RegionUID != "" {
		tDb = tDb.Where("Company.region_uid = ?", customerQuery.RegionUID)
	}
	if customerQuery.CompanyUID != "" {
		tDb = tDb.Where("customer.company_uid = ?", customerQuery.CompanyUID)
	}

	if customerQuery.Name != "" {
		tDb = tDb.Where("customer.name LIKE ?", "%"+customerQuery.Name+"%")
	}
	if customerQuery.ResearchGroup != "" {
		tDb = tDb.Where("customer.research_group LIKE ?", "%"+customerQuery.ResearchGroup+"%")
	}

	err = tDb.Find(&customers).Count(&total).
		Limit(pageSize).Offset((pageNo - 1) * pageSize).
		Find(&customers).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return customers, msg.ERROR_CUSTOMER_SELECT, total
	}
	return customers, msg.SUCCESS, total
}

func InsertCustomerCompany(customerCompany *CustomerCompany) (code int) {
	customerCompany.UID = uid.Generate()
	err = db.Create(&customerCompany).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func DeleteCustomerCompany(uid string) (code int) {
	err = db.Model(&CustomerCompany{}).Where("uid = ?", uid).Update("is_delete", true).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func UpdateCustomerCompany(customerCompany *CustomerCompany) (code int) {
	var maps = make(map[string]interface{})
	maps["region_uid"] = customerCompany.RegionUID
	maps["name"] = customerCompany.Name
	maps["address"] = customerCompany.Address

	err = db.Model(&CustomerCompany{}).Where("uid = ?", customerCompany.UID).Updates(maps).Error

	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func SelectCustomerCompanys(pageSize int, pageNo int, customerCompany *CustomerCompany) (CustomerCompanys []CustomerCompany, code int, total int64) {
	var maps = make(map[string]interface{})
	maps["is_delete"] = false
	if customerCompany.RegionUID != "" {
		maps["region_uid"] = customerCompany.RegionUID
	}

	tdb := db.Where(maps)

	if customerCompany.Name != "" {
		tdb = tdb.Where("name LIKE ?", "%"+customerCompany.Name+"%")
	}

	err = tdb.Find(&CustomerCompanys).Count(&total).
		Preload("Region").
		Limit(pageSize).Offset((pageNo - 1) * pageSize).
		Find(&CustomerCompanys).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return CustomerCompanys, msg.ERROR, total
	}
	return CustomerCompanys, msg.SUCCESS, total
}
