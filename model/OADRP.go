package model

import "gorm.io/gorm"

//Office办事处 Area地区 Department部门 Role角色 Permission权限

type Office struct {
	gorm.Model
	Name  string
	Areas []Area
}

type Area struct {
	gorm.Model
	OfficeID uint
	Name     string
}

type Department struct {
	gorm.Model
	OfficeID uint
	Name     string
}

type Role struct {
	gorm.Model
	Name        string
	Permissions []Permission `gorm:"many2many:role_permission;"`
}

type Permission struct {
	gorm.Model
	Name string
	Text string
}
