package model

import (
	"business-system_golang/utils/msg"
	"business-system_golang/utils/pwd"

	"gorm.io/gorm"
)

// 员工 Model
type Employee struct {
	gorm.Model
	Phone        string `gorm:"type:varchar(20);comment:电话;not null" json:"phone"`
	Name         string `gorm:"type:varchar(20);comment:姓名;not null" json:"name"`
	Password     string `gorm:"type:varchar(20);comment:密码;not null" json:"password"`
	WechatID     string `gorm:"type:varchar(20);comment:微信号" json:"wechatID"`
	Email        string `gorm:"type:varchar(20);comment:邮箱" json:"email"`
	OfficeID     uint   `gorm:"type:int;comment:办事处ID;default:(-)" json:"officeID"`
	DepartmentID uint   `gorm:"type:int;comment:部门ID;default:(-)" json:"departmentID"`
	RoleID       uint   `gorm:"type:int;comment:角色ID;default:(-)" json:"roleID"`

	Office     Office     `gorm:"foreignKey:OfficeID" json:"office"`
	Department Department `gorm:"foreignKey:DepartmentID" json:"department"`
	Role       Role       `gorm:"foreignKey:RoleID" json:"role"`
}

//查询员工(手机号)是否录入
func CheckEmployeePhone(phone string) (code int) {
	var employee Employee
	db.Where("phone = ?", phone).First(&employee)
	if employee.ID > 0 {
		return msg.ERROR_EMPLOYEE_EXIST
	}
	return msg.ERROR_EMPLOYEE_NOT_EXIST
}

func CheckEmployeeByPhoneAndPwd(employee *Employee) (code int) {
	employee.Password, _ = pwd.ScryptPwd(employee.Password)
	db.Where("phone = ? AND password = ?", employee.Phone, employee.Password).First(&employee)
	if employee.ID > 0 {
		return msg.SUCCESS
	}
	return msg.ERROR_EMPLOYEE_LOGIN_FAIL
}

func CreateEmployee(employee *Employee) (code int) {
	err = db.Create(&employee).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func DeleteEmployee(id int) (code int) {
	err = db.Where("id = ?", id).Delete(&Employee{}).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func UpdateEmployee(employee *Employee) (code int) {
	var maps = make(map[string]interface{})
	maps["WechatID"] = employee.WechatID
	maps["Email"] = employee.Email
	err = db.Model(&employee).Updates(maps).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func SelectEmployee(id int) (employee Employee, code int) {
	err = db.Preload("Office").Preload("Department").Preload("Role").First(&employee, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return employee, msg.ERROR_EMPLOYEE_NOT_EXIST
		} else {
			return employee, msg.ERROR
		}
	}
	return employee, msg.SUCCESS
}

func SelectEmployees(pageSize int, pageNo int) (employees []Employee, code int, total int64) {
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

func (employee *Employee) AfterFind(tx *gorm.DB) (err error) {
	employee.Password = ""
	return
}
