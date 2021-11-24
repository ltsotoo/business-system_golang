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
	CompanyUID    string `gorm:"type:varchar(32);comment:公司UID;not null" json:"companyUID"`
	Name          string `gorm:"type:varchar(50);comment:姓名;not null" json:"name"`
	ResearchGroup string `gorm:"type:varchar(50);comment:课题组" json:"researchGroup"`
	Phone         string `gorm:"type:varchar(50);comment:电话" json:"phone"`
	WechatID      string `gorm:"type:varchar(50);comment:微信号" json:"wechatID"`
	Email         string `gorm:"type:varchar(50);comment:邮箱" json:"email"`
	Status        int    `gorm:"type:int;comment:状态(0:未审核,1:通过审核)" json:"status"`
	IsDelete      bool   `gorm:"type:boolean;comment:是否删除" json:"isDelete"`

	Company CustomerCompany `gorm:"foreignKey:CompanyUID;references:UID" json:"company"`
}

type CustomerQuery struct {
	AreaUID       string `json:"areaUID"`
	CompanyUID    string `json:"companyUID"`
	CompanyName   string `json:"companyName"`
	ResearchGroup string `json:"researchGroup"`
	Name          string `json:"name"`
}

// 客户公司 Model
type CustomerCompany struct {
	BaseModel
	UID      string `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	AreaUID  string `gorm:"type:varchar(32);comment:区域UID;not null" json:"areaUID"`
	Name     string `gorm:"type:varchar(50);comment:名称;not null" json:"name"`
	Address  string `gorm:"type:varchar(50);comment:地址" json:"address"`
	IsDelete bool   `gorm:"type:boolean;comment:是否删除" json:"isDelete"`

	Area Area `gorm:"foreignKey:AreaUID;references:UID" json:"area"`
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
	// err = db.Where("uid = ?", uid).Delete(&Customer{}).Error
	err = db.Model(&Customer{}).Where("uid = ?", uid).Update("is_delete", true).Error
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
	if customerQuery.AreaUID != "" {
		tDb = tDb.Where("Company.area_uid = ?", customerQuery.AreaUID)
	}
	if customerQuery.CompanyUID != "" {
		tDb = tDb.Where("customer.company_uid = ?", customerQuery.CompanyUID)
	}
	if customerQuery.CompanyName != "" {
		tDb = tDb.Where("Company.name LIKE ?", "%"+customerQuery.CompanyName+"%")
	}
	if customerQuery.ResearchGroup != "" {
		tDb = tDb.Where("customer.research_group LIKE ?", "%"+customerQuery.ResearchGroup+"%")
	}
	if customerQuery.Name != "" {
		tDb = tDb.Where("customer.name LIKE ?", "%"+customerQuery.Name+"%")
	}

	err = tDb.Find(&customers).Count(&total).
		Preload("Company").Preload("Company.Area").
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
	// err = db.Where("uid = ?", uid).Delete(&CustomerCompany{}).Error
	err = db.Model(&CustomerCompany{}).Where("uid = ?", uid).Update("is_delete", true).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func SelectCustomerCompanys(customerCompany *CustomerCompany) (CustomerCompanys []CustomerCompany, code int) {
	var maps = make(map[string]interface{})
	if customerCompany.AreaUID != "" {
		maps["area_uid"] = customerCompany.AreaUID
		maps["is_delete"] = false
	}
	err = db.Preload("Area").Where(maps).Where("name LIKE ?", "%"+customerCompany.Name+"%").Find(&CustomerCompanys).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return CustomerCompanys, msg.ERROR
	}
	return CustomerCompanys, msg.SUCCESS
}
