package model

import (
	"business-system_golang/utils/msg"
	"business-system_golang/utils/pwd"
	"business-system_golang/utils/uid"

	"gorm.io/gorm"
)

// 员工 Model
type Employee struct {
	gorm.Model
	UID           string `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	Phone         string `gorm:"type:varchar(20);comment:电话;not null" json:"phone"`
	Name          string `gorm:"type:varchar(20);comment:姓名;not null" json:"name"`
	Password      string `gorm:"type:varchar(20);comment:密码;not null" json:"password"`
	WechatID      string `gorm:"type:varchar(20);comment:微信号" json:"wechatID"`
	Email         string `gorm:"type:varchar(20);comment:邮箱" json:"email"`
	OfficeUID     string `gorm:"type:varchar(32);comment:办事处UID;default:(-)" json:"officeUID"`
	DepartmentUID string `gorm:"type:varchar(32);comment:部门UID;default:(-)" json:"departmentUID"`

	Office     Office     `gorm:"foreignKey:OfficeUID;references:UID" json:"office"`
	Department Department `gorm:"foreignKey:DepartmentUID;references:UID" json:"department"`
	Roles      []Role     `gorm:"many2many:employee_role;foreignKey:UID;references:UID" json:"roles"`
}

type EmployeeQuery struct {
	Name          string `json:"name"`
	Phone         string `json:"phone"`
	OfficeUID     string `json:"officeUID"`
	DepartmentUID string `json:"departmentUID"`
}

//查询员工(手机号)是否录入
func CheckEmployee(phone string) (code int) {
	var employee Employee
	db.Where("phone = ?", phone).First(&employee)
	if employee.ID > 0 {
		return msg.ERROR_EMPLOYEE_EXIST
	}
	return msg.ERROR_EMPLOYEE_NOT_EXIST
}

func CheckLogin(phone string, password string) (employee Employee, code int) {
	db.Where("phone = ?", phone).First(&employee)
	if employee.ID == 0 {
		return employee, msg.ERROR_EMPLOYEE_LOGIN_FAIL
	}
	password, err = pwd.ScryptPwd(password)
	if err != nil || employee.Password != password {
		return employee, msg.ERROR_EMPLOYEE_LOGIN_FAIL
	}
	return employee, msg.SUCCESS
}

func SelectAllPermission(employeeUID string, departmentUID string) (permissions []string) {
	//查出所有的权限(去重)
	db.Raw("SELECT distinct permission_uid FROM role_permission WHERE role_uid IN "+
		"(SELECT role_uid AS uid FROM employee_role WHERE employee_uid = ? UNION "+
		"SELECT uid FROM role WHERE department_uid IN "+
		"(SELECT type_uid FROM department WHERE	uid = ?))", employeeUID, departmentUID).Scan(&permissions)
	return
}

func InsertEmployee(employee *Employee) (code int) {
	employee.UID = uid.Generate()
	employee.Password, err = pwd.ScryptPwd(employee.Password)
	err = db.Create(&employee).Error
	if err != nil {
		return msg.ERROR_CUSTOMER_INSERT
	}
	return msg.SUCCESS
}

func DeleteEmployee(uid string) (code int) {
	err = db.Where("uid = ?", uid).Delete(&Employee{}).Error
	if err != nil {
		return msg.ERROR_CUSTOMER_DELETE
	}
	return msg.SUCCESS
}

func UpdateEmployee(employee *Employee) (code int) {
	err = db.Transaction(func(tdb *gorm.DB) error {
		if txErr := tdb.Model(&Employee{}).Where("uid = ?", employee.UID).Updates(employee).Error; err != nil {
			return txErr
		}
		if txErr := tdb.Model(&employee).Association("Roles").Replace(employee.Roles); err != nil {
			return txErr
		}
		return nil
	})
	if err != nil {
		return msg.ERROR_CUSTOMER_UPDATE
	}
	return msg.SUCCESS
}

func SelectEmployee(uid string) (employee Employee, code int) {
	err = db.Preload("Office").Preload("Department").Preload("Roles").
		First(&employee, "uid = ?", uid).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return employee, msg.ERROR_EMPLOYEE_NOT_EXIST
		} else {
			return employee, msg.ERROR
		}
	}
	return employee, msg.SUCCESS
}

func SelectEmployees(pageSize int, pageNo int, employeeQuery *EmployeeQuery) (employees []Employee, code int, total int64) {
	var maps = make(map[string]interface{})
	if employeeQuery.OfficeUID != "" {
		maps["office_uid"] = employeeQuery.OfficeUID
	}
	if employeeQuery.DepartmentUID != "" {
		maps["department_uid"] = employeeQuery.DepartmentUID
	}

	err = db.Model(&employees).Where(maps).
		Where("name LIKE ? AND phone LIKE ?", "%"+employeeQuery.Name+"%", "%"+employeeQuery.Phone+"%").
		Count(&total).
		Preload("Office").Preload("Department").
		Limit(pageSize).Offset((pageNo - 1) * pageSize).
		Find(&employees).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return employees, msg.ERROR_CUSTOMER_SELECT, total
	}
	return employees, msg.SUCCESS, total
}
