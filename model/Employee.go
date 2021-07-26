package model

import (
	"business-system_golang/utils/msg"
	"business-system_golang/utils/pwd"

	"gorm.io/gorm"
)

// 员工 Model
type Employee struct {
	gorm.Model
	Phone    string `gorm:"type:varchar(20);comment:电话;not null" json:"phone"`
	Name     string `gorm:"type:varchar(20);comment:姓名;not null" json:"name"`
	Password string `gorm:"type:varchar(20);comment:密码;not null" json:"password"`
	AreaId   int    `gorm:"type:varchar(20);comment:所属区域ID;not null" json:"aread"`
	WechatID string `gorm:"type:varchar(20);comment:微信号" json:"wechatId"`
	Email    string `gorm:"type:varchar(20);comment:邮箱" json:"email"`
}

//查询手机号是否注册
func CheckEmployeePhone(phone string) int {
	var employee Employee
	db.Where("phone = ?", phone).First(&employee)
	if employee.ID > 0 {
		return msg.ERROR_EMPLOYEE_EXIST
	}
	return msg.ERROR_EMPLOYEE_NOT_EXIST
}

//增加员工
func CreateEmployee(employee *Employee) int {
	err = db.Create(&employee).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

//删除员工
func DeleteEmployee(id int) int {
	var employee Employee
	err = db.Where("id = ?", id).Delete(&employee).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

//编辑员工
func UpdateEmployee(employee *Employee) int {
	var maps = make(map[string]interface{})
	maps["WechatID"] = employee.WechatID
	maps["Email"] = employee.Email
	err = db.Model(&employee).Updates(maps).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

//查询员工
func SelectEmployee(id int) (Employee, int) {
	var employee Employee
	err = db.First(&employee, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return employee, msg.ERROR_EMPLOYEE_NOT_EXIST
		} else {
			return employee, msg.ERROR
		}
	}
	return employee, msg.SUCCESS
}

//查询员工列表
func SelectEmployees(pageSize int, pageNo int) ([]Employee, int, int64) {
	var employees []Employee
	var total int64
	err = db.Limit(pageSize).Offset((pageNo - 1) * pageSize).Find(&employees).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, msg.ERROR, 0
	}
	db.Model(&employees).Count(&total)
	return employees, msg.SUCCESS, total
}

func (employee *Employee) BeforeCreate(tx *gorm.DB) (err error) {
	employee.Password, err = pwd.ScryptPwd(employee.Password)
	return err
}
