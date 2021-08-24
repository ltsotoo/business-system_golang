package model

import (
	"business-system_golang/utils/msg"

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

type Department struct {
	gorm.Model
	UID      string `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	OfficeID uint   `gorm:"type:int;comment:办事处ID;not null" json:"officeID"`
	Name     string `gorm:"type:varchar(20);comment:名称;not null" json:"name"`
}

type Role struct {
	gorm.Model
	UID         string       `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	Name        string       `gorm:"type:varchar(20);comment:名称;not null" json:"name"`
	Permissions []Permission `gorm:"many2many:role_permission;" json:"rermissions"`
}

type Permission struct {
	gorm.Model
	UID  string `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
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

func DeleteOffice(id int) (code int) {
	err = db.Where("id = ?", id).Delete(&Office{}).Error
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

func DeleteArea(id int) (code int) {
	err = db.Where("id = ?", id).Delete(&Area{}).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func UpdateArea(area *Area) (code int) {
	// err = db.Model(&area).Updates(Area{Name: area.Name, OfficeID: area.OfficeID}).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func SelectAreas(area *Area) (areas []Area, code int) {
	// err = db.Preload("Office").Joins("Office").Where("area.name LIKE ? AND Office.name LIKE ?", "%"+area.Name+"%", "%"+area.Office.Name+"%").Find(&areas).Error
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

func DeleteDepartment(id int) (code int) {
	err = db.Where("id = ?", id).Delete(&Department{}).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func SelectDepartments(department *Department) (departments []Department, code int) {
	err = db.Where("office_id = ? AND name LIKE ?", department.OfficeID, "%"+department.Name+"%").Find(&departments).Error
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
