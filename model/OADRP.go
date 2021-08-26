package model

import (
	"business-system_golang/utils/msg"
	"business-system_golang/utils/uid"

	"gorm.io/gorm"
)

//Office办事处 Area地区 Department部门 Role角色 Permission权限

type Office struct {
	gorm.Model
	UID  string `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	Name string `gorm:"type:varchar(20);comment:名称;not null" json:"name"`

	Areas []Area `gorm:"foreignKey:OfficeUID;references:UID" json:"areas"`
}

type Area struct {
	gorm.Model
	UID       string `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	OfficeUID string `gorm:"type:varchar(32);comment:办事处UID;not null" json:"officeUID"`
	Name      string `gorm:"type:varchar(20);comment:名称;not null" json:"name"`

	Office Office `gorm:"foreignKey:OfficeUID;references:UID" json:"office"`
}

type AreaQuery struct {
	Name       string `json:"name"`
	OfficeName string `json:"officeName"`
}

type Department struct {
	gorm.Model
	UID       string `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	Type      string `gorm:"type:varchar(32);comment:部门类型;not null" json:"type"`
	OfficeUID string `gorm:"type:varchar(32);comment:办事处ID;not null" json:"officeUID"`
	Name      string `gorm:"type:varchar(20);comment:名称;not null" json:"name"`
}

type Role struct {
	gorm.Model
	UID           string       `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	Name          string       `gorm:"type:varchar(20);comment:名称;not null" json:"name"`
	DepartmentUID string       `gorm:"type:varchar(32);comment:部门UID" json:"departmentUID"`
	Permissions   []Permission `gorm:"many2many:role_permission;foreignKey:UID;References:UID" json:"rermissions"`
}

type Permission struct {
	gorm.Model
	UID    string `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	Module string `gorm:"type:varchar(20);comment:模块;not null" json:"module"`
	Name   string `gorm:"type:varchar(20);comment:名称;not null" json:"name"`
	Text   string `gorm:"type:varchar(20);comment:描述;not null" json:"text"`
}

func InsertOffice(office *Office) (code int) {
	office.UID = uid.Generate()
	err = db.Create(&office).Error
	if err != nil {
		return msg.ERROR_OFFICE_INSERT
	}
	return msg.SUCCESS
}

func DeleteOffice(uid string) (code int) {
	err = db.Where("uid = ?", uid).Delete(&Office{}).Error
	if err != nil {
		return msg.ERROR_OFFICE_DELETE
	}
	return msg.SUCCESS
}

func SelectOffice(uid string) (office Office, code int) {
	err = db.First(&office, "uid = ?", uid).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return office, msg.ERROR_OFFICE_NOT_EXIST
		} else {
			return office, msg.ERROR_OFFICE_SELECT
		}
	}
	return office, msg.SUCCESS
}

func SelectOffices(name string) (offices []Office, code int) {
	err = db.Where("name LIKE ?", "%"+name+"%").Find(&offices).Error
	if err != nil {
		return offices, msg.ERROR_OFFICE_SELECT
	}
	return offices, msg.SUCCESS
}

func InsertArea(area *Area) (code int) {
	area.UID = uid.Generate()
	err = db.Create(&area).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func DeleteArea(uid string) (code int) {
	err = db.Where("uid = ?", uid).Delete(&Area{}).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func UpdateArea(area *Area) (code int) {
	err = db.Model(&Area{}).Where("uid = ?", area.UID).Updates(Area{Name: area.Name, OfficeUID: area.OfficeUID}).Error
	if err != nil {
		return msg.ERROR_AREA_UPDATE
	}
	return msg.SUCCESS
}

func SelectAreas(areaQuery *AreaQuery) (areas []Area, code int) {
	err = db.Preload("Office").Joins("Office").
		Where("area.name LIKE ? AND Office.name LIKE ?", "%"+areaQuery.Name+"%", "%"+areaQuery.OfficeName+"%").
		Find(&areas).Error
	if err != nil {
		return areas, msg.ERROR_AREA_SELECT
	}
	return areas, msg.SUCCESS
}

func InsertDepartment(department *Department) (code int) {
	department.UID = uid.Generate()
	err = db.Create(&department).Error
	if err != nil {
		return msg.ERROR_DEPARTMENT_INSERT
	}
	return msg.SUCCESS
}

func DeleteDepartment(uid string) (code int) {
	err = db.Where("uid = ?", uid).Delete(&Department{}).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func SelectDepartments(department *Department) (departments []Department, code int) {
	err = db.Where("office_uid = ? AND name LIKE ?", department.OfficeUID, "%"+department.Name+"%").
		Find(&departments).Error
	if err != nil {
		return nil, msg.ERROR_DEPARTMENT_SELECT
	}
	return departments, msg.SUCCESS
}

func SelectRoles() (roles []Role, code int) {
	err = db.Find(&roles).Error
	if err != nil {
		return nil, msg.ERROR_ROLE_SELECT
	}
	return roles, msg.SUCCESS
}

func SelectPermissions() (permissions []Permission, code int) {
	err = db.Find(&permissions).Error
	if err != nil {
		return nil, msg.ERROR_PERMISSION_SELECT
	}
	return permissions, msg.SUCCESS
}

func (office *Office) BeforeCreate(tx *gorm.DB) (err error) {
	office.UID = uid.Generate()
	return err
}

func (area *Area) BeforeCreate(tx *gorm.DB) (err error) {
	area.UID = uid.Generate()
	return err
}

func (department *Department) BeforeCreate(tx *gorm.DB) (err error) {
	department.UID = uid.Generate()
	return err
}

func (role *Role) BeforeCreate(tx *gorm.DB) (err error) {
	role.UID = uid.Generate()
	return err
}

func (permission *Permission) BeforeCreate(tx *gorm.DB) (err error) {
	permission.UID = uid.Generate()
	return err
}
