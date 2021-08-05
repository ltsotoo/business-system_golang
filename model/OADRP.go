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

func SelectOffices() (offices []Office, code int) {
	err = db.Find(&offices).Error
	if err != nil {
		return nil, msg.ERROR
	}
	return offices, msg.SUCCESS
}

func SelectAreas() (areas []Area, code int) {
	err = db.Find(&areas).Error
	if err != nil {
		return nil, msg.ERROR
	}
	return areas, msg.SUCCESS
}

func SelectAreasByOfficeID(officeID int) (areas []Area, code int) {
	err = db.Where("office_id = ?", officeID).Find(&areas).Error
	if err != nil {
		return nil, msg.ERROR
	}
	return areas, msg.SUCCESS
}

func SelectDepartmentsByOfficeID(officeID int) (departments []Department, code int) {
	err = db.Where("office_id = ?", officeID).Find(&departments).Error
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
