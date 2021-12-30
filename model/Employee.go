package model

import (
	"business-system_golang/utils/msg"
	"business-system_golang/utils/pwd"
	"business-system_golang/utils/uid"
	"fmt"

	"gorm.io/gorm"
)

// 员工 Model
type Employee struct {
	BaseModel
	UID           string  `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	Phone         string  `gorm:"type:varchar(50);comment:手机号/登录凭证;not null" json:"phone"`
	Name          string  `gorm:"type:varchar(50);comment:姓名;not null" json:"name"`
	Password      string  `gorm:"type:varchar(50);comment:密码;not null" json:"password"`
	WechatID      string  `gorm:"type:varchar(50);comment:微信号" json:"wechatID"`
	Email         string  `gorm:"type:varchar(50);comment:邮箱" json:"email"`
	OfficeUID     string  `gorm:"type:varchar(32);comment:办事处UID;default:(-)" json:"officeUID"`
	DepartmentUID string  `gorm:"type:varchar(32);comment:部门UID;default:(-)" json:"departmentUID"`
	Number        string  `gorm:"type:varchar(50);comment:编号" json:"number"`
	ContractCount int     `gorm:"type:int;comment:被审批合同总数" json:"contractCount"`
	Money         float64 `gorm:"type:decimal(20,6);comment:补助额度(元)" json:"money"`
	Credit        float64 `gorm:"type:decimal(20,6);comment:每月总部补助额度(元)" json:"credit"`
	OfficeCredit  float64 `gorm:"type:decimal(20,6);comment:每月办事处补助额度(元)" json:"officeCredit"`
	IsDelete      bool    `gorm:"type:boolean;comment:是否删除" json:"isDelete"`

	Roles      []Role     `gorm:"many2many:employee_role;foreignKey:UID;references:UID" json:"roles"`
	Office     Office     `gorm:"foreignKey:OfficeUID;references:UID" json:"office"`
	Department Department `gorm:"foreignKey:DepartmentUID;references:UID" json:"department"`

	EditType int `gorm:"-" json:"editType"`
}

type EmployeeQuery struct {
	Name          string `json:"name"`
	Phone         string `json:"phone"`
	OfficeUID     string `json:"officeUID"`
	DepartmentUID string `json:"departmentUID"`
	PermissionUID string `son:"permissionUID"`

	UID    string `json:"uid"`
	OldPWd string `json:"oldPwd"`
	NewPwd string `json:"newPwd"`
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
	db.Where("phone = ? AND is_delete = ?", phone, false).First(&employee)
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
	db.Raw("SELECT distinct permission_uid FROM role_permission WHERE role_uid IN (SELECT role_uid AS uid FROM employee_role WHERE employee_uid = ? UNION SELECT role_uid AS uid FROM department WHERE uid = ?)", employeeUID, departmentUID).Scan(&permissions)
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
	// err = db.Where("uid = ?", uid).Delete(&Employee{}).Error
	err = db.Model(&Employee{}).Where("uid = ?", uid).Update("is_delete", true).Error
	if err != nil {
		return msg.ERROR_CUSTOMER_DELETE
	}
	return msg.SUCCESS
}

func UpdateEmployee(employee *Employee) (code int) {
	var maps = make(map[string]interface{})
	if employee.EditType == 1 {
		maps["name"] = employee.Name
		maps["phone"] = employee.Phone
		maps["wechat_id"] = employee.WechatID
		maps["email"] = employee.Email
	} else if employee.EditType == 2 {
		maps["money"] = employee.Money
		maps["credit"] = employee.Credit
		maps["office_credit"] = employee.OfficeCredit
	} else if employee.EditType == 3 {
		maps["office_uid"] = employee.OfficeUID
		maps["department_uid"] = employee.DepartmentUID
		maps["number"] = employee.Number
	}

	if employee.EditType != 3 {
		err = db.Model(&Employee{}).Where("uid = ?", employee.UID).Updates(maps).Error
	} else {
		err = db.Transaction(func(tdb *gorm.DB) error {
			if tErr := tdb.Model(&Employee{}).Where("uid = ?", employee.UID).Updates(maps).Error; tErr != nil {
				return tErr
			}
			if tErr := tdb.Model(&employee).Association("Roles").Replace(employee.Roles); tErr != nil {
				return tErr
			}
			return nil
		})
	}

	if err != nil {
		return msg.ERROR_CUSTOMER_UPDATE
	}
	return msg.SUCCESS
}

func SelectEmployee(uid string) (employee Employee, code int) {
	err = db.Preload("Office").Preload("Department").Preload("Roles").
		Where("is_delete = ?", false).
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

func SelectSPEmployees(employeeQuery *EmployeeQuery) (employees []Employee, code int) {
	err = db.Raw("SELECT distinct * FROM employee WHERE (uid IN ("+
		"SELECT employee_uid FROM employee_role WHERE role_uid IN ("+
		"SELECT role_uid FROM role_permission WHERE permission_uid = ?"+
		")) OR department_uid IN (SELECT uid FROM department WHERE role_uid IN ("+
		"SELECT role_uid FROM role_permission WHERE permission_uid = ?))) AND "+
		"office_uid = ?", employeeQuery.PermissionUID, employeeQuery.PermissionUID, employeeQuery.OfficeUID).Scan(&employees).Error
	if err != nil {
		code = msg.ERROR
	} else {
		code = msg.SUCCESS
	}
	return
}

func SelectEmployees(pageSize int, pageNo int, employeeQuery *EmployeeQuery) (employees []Employee, code int, total int64) {
	var maps = make(map[string]interface{})
	maps["is_delete"] = false
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

func UpdatePwd(employeeQuery *EmployeeQuery) (code int) {
	employeeQuery.OldPWd, err = pwd.ScryptPwd(employeeQuery.OldPWd)
	if err == nil {
		var employee Employee
		db.Where("uid = ? AND password = ?", employeeQuery.UID, employeeQuery.OldPWd).First(&employee)
		employeeQuery.NewPwd, err = pwd.ScryptPwd(employeeQuery.NewPwd)
		if employee.ID > 0 {
			err = db.Model(&Employee{}).Where("uid = ?", employee.UID).Update("password", employeeQuery.NewPwd).Error
			if err == nil {
				return msg.SUCCESS
			}
		}
	}
	return msg.ERROR_EMPLOYEE_PASSWORD_FAIL
}

func ResetPwd(uid string) (code int) {
	var employee Employee
	var tempPwd string
	db.Preload("Office").First(&employee, "uid = ?", uid)
	if employee.ID > 0 {
		tempPwd = employee.Phone + employee.Office.Number + employee.Number
		tempPwd, err = pwd.ScryptPwd(tempPwd)
		err = db.Model(&Employee{}).Where("uid = ?", uid).Update("password", tempPwd).Error
		if err != nil {
			return msg.ERROR
		}
		return msg.SUCCESS
	} else {
		return msg.ERROR
	}
}

func ResetAllPwd() (code int) {
	var employees []Employee
	err = db.Preload("Office").Find(&employees).Error
	if err == nil {
		for i := range employees {
			tempPwd := employees[i].Phone + employees[i].Office.Number + employees[i].Number
			tempPwd, _ = pwd.ScryptPwd(tempPwd)
			err = db.Model(&Employee{}).Where("uid = ?", employees[i].UID).Update("password", tempPwd).Error
			if err != nil {
				return msg.ERROR_EMPLOYEE_PASSWORD_FAIL
			}
		}
		return msg.SUCCESS
	} else {
		return msg.ERROR
	}
}

func UpdateAllAddMoney() {
	err = db.Transaction(func(tdb *gorm.DB) error {
		if tErr := db.Exec("UPDATE employee SET money = money + credit WHERE is_delete = ?", false).Error; tErr != nil {
			return tErr
		}
		if tErr := db.Exec("UPDATE employee SET money = money + office_credit WHERE is_delete = ?", false).Error; tErr != nil {
			return tErr
		}
		if tErr := db.Exec("UPDATE office,(SELECT office_uid,sum(office_credit) AS sum FROM employee WHERE office_uid IS NOT NULL GROUP BY office_uid) AS a SET office.money = office.money - a.sum WHERE office.uid = a.office_uid").Error; tErr != nil {
			return tErr
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}
