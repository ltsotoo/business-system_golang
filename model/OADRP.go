package model

import (
	"business-system_golang/utils/msg"

	"gorm.io/gorm"
)

//Office办事处 Area地区 Department部门 Role角色 Permission权限

type Office struct {
	gorm.Model
	Name  string `gorm:"type:varchar(20);comment:名称;not null" json:"name"`
	Areas []Area `gorm:"foreignKey:OfficeID" json:"areas"`
}

type Area struct {
	gorm.Model
	OfficeID uint   `gorm:"type:int;comment:办事处ID;not null" json:"officeID"`
	Name     string `gorm:"type:varchar(20);comment:名称;not null" json:"name"`

	Office Office `gorm:"foreignKey:OfficeID" json:"office"`
}

type Department struct {
	gorm.Model
	OfficeID uint   `gorm:"type:int;comment:办事处ID;not null" json:"officeID"`
	Name     string `gorm:"type:varchar(20);comment:名称;not null" json:"name"`
}

type Role struct {
	gorm.Model
	Name        string       `gorm:"type:varchar(20);comment:名称;not null" json:"name"`
	Permissions []Permission `gorm:"many2many:role_permission;" json:"rermissions"`
}

type Permission struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);comment:名称;not null" json:"name"`
	Text string `gorm:"type:varchar(20);comment:描述;not null" json:"text"`
}

func CreateOffice(office *Office) (code int) {
	err = db.Create(&office).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func SelectOffice(id int) (office Office, code int) {
	err = db.First(&office, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return office, msg.ERROR_Office_NOT_EXIST
		} else {
			return office, msg.ERROR
		}
	}
	return office, msg.SUCCESS
}

func SelectOffices(name string) (offices []Office, code int) {
	if name == "" {
		err = db.Find(&offices).Error
	} else {
		err = db.Where("name LIKE ?", "%"+name+"%").Find(&offices).Error
	}
	if err != nil {
		return nil, msg.ERROR
	}
	return offices, msg.SUCCESS
}

func CreateArea(area *Area) (code int) {
	err = db.Create(&area).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func SelectAreas(area *Area) (areas []Area, code int) {
	err = db.Where(&area).Find(&areas).Error
	if err != nil {
		return nil, msg.ERROR
	}
	return areas, msg.SUCCESS
}

func CreateDepartment(department *Department) (code int) {
	err = db.Create(&department).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func SelectDepartments(department *Department) (departments []Department, code int) {
	err = db.Where(&department).Find(&departments).Error
	if err != nil {
		return nil, msg.ERROR
	}
	return departments, msg.SUCCESS
}

func SelectRoles() (roles []Role, code int) {
	err = db.Find(&roles).Error
	if err != nil {
		return nil, msg.ERROR
	}
	return roles, msg.SUCCESS
}

func SelectPermissions() (permissions []Permission, code int) {
	err = db.Find(&permissions).Error
	if err != nil {
		return nil, msg.ERROR
	}
	return permissions, msg.SUCCESS
}
