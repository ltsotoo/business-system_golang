package model

import (
	"business-system_golang/utils/msg"
	"business-system_golang/utils/uid"

	"gorm.io/gorm"
)

//Office办事处 Area地区 Department部门 Role角色 Permission权限
type Office struct {
	BaseModel
	UID   string  `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	Name  string  `gorm:"type:varchar(20);comment:名称;not null" json:"name"`
	Money float64 `gorm:"type:decimal(20,6);comment:办事处总报销额度(元)" json:"money"`

	Areas []Area `gorm:"foreignKey:OfficeUID;references:UID" json:"areas"`
}

type Area struct {
	BaseModel
	UID       string `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	OfficeUID string `gorm:"type:varchar(32);comment:办事处UID;default:(-)" json:"officeUID"`
	Name      string `gorm:"type:varchar(20);comment:名称;not null" json:"name"`
	Number    string `gorm:"type:varchar(20);comment:编号;unique" json:"number"`

	Office Office `gorm:"foreignKey:OfficeUID;references:UID" json:"office"`
}

type AreaQuery struct {
	Name       string `json:"name"`
	OfficeName string `json:"officeName"`
}

type Department struct {
	BaseModel
	UID       string `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	TypeUID   string `gorm:"type:varchar(32);comment:部门类型;not null" json:"typeUID"`
	OfficeUID string `gorm:"type:varchar(32);comment:办事处ID;not null" json:"officeUID"`
	Name      string `gorm:"type:varchar(20);comment:名称;not null" json:"name"`
	RoleUID   string `gorm:"type:varchar(32);comment:部门员工默认拥有职位" json:"roleUID"`

	Type Dictionary `gorm:"foreignKey:TypeUID;references:UID" json:"type"`
	Role Role       `gorm:"foreignKey:RoleUID;references:UID" json:"role"`
}

type Role struct {
	BaseModel
	UID         string       `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	Name        string       `gorm:"type:varchar(20);comment:名称;not null" json:"name"`
	Permissions []Permission `gorm:"many2many:role_permission;foreignKey:UID;References:UID" json:"permissions"`
}

type Permission struct {
	BaseModel
	UID    string `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	Module string `gorm:"type:varchar(20);comment:模块;not null" json:"module"`
	Name   string `gorm:"type:varchar(20);comment:名称;not null" json:"name"`
	Text   string `gorm:"type:varchar(20);comment:描述;not null" json:"text"`
	No     string `gorm:"type:varchar(3);comment:序号;default:(-)" json:"no"`
	UrlUID string `gorm:"type:varchar(32);comment:Url_UID;default:(-)" json:"urlUID"`
}

type Url struct {
	BaseModel
	UID   string `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	Title string `gorm:"type:varchar(20);comment:标题;not null" json:"title"`
	Icon  string `gorm:"type:varchar(20);comment:图标;not null" json:"icon"`
	Url   string `gorm:"type:varchar(20);comment:url;not null" json:"url"`
	No    int    `gorm:"type:int;comment:序号" json:"no"`
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

func UpdateOffice(office *Office) (code int) {
	err = db.Model(&Office{}).Where("uid = ?", office.UID).Updates(Office{Name: office.Name, Money: office.Money}).Error
	if err != nil {
		return msg.ERROR_OFFICE_UPDATE
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

func SelectOffices(office *Office) (offices []Office, code int) {
	err = db.Where("name LIKE ?", "%"+office.Name+"%").Find(&offices).Error
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
	err = db.Model(&Area{}).Where("uid = ?", area.UID).Updates(Area{Name: area.Name, OfficeUID: area.OfficeUID, Number: area.Number}).Error
	if err != nil {
		return msg.ERROR_AREA_UPDATE
	}
	return msg.SUCCESS
}

func SelectArea(uid string) (area Area, code int) {
	err = db.Where("uid = ?", uid).First(&area).Error
	if err != nil {
		return area, msg.ERROR_AREA_SELECT
	}
	return area, msg.SUCCESS
}

func SelectAreas(areaQuery *AreaQuery) (areas []Area, code int) {
	tDb := db.Preload("Office").Joins("Office")
	if areaQuery.Name != "" {
		tDb = tDb.Where("area.name LIKE ?", "%"+areaQuery.Name+"%")
	}
	if areaQuery.OfficeName != "" {
		tDb = tDb.Where("Office.name LIKE ?", "%"+areaQuery.OfficeName+"%")
	}
	err = tDb.Find(&areas).Error
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

func UpdateDepartment(department *Department) (code int) {
	err = db.Model(&Department{}).Where("uid = ?", department.UID).Updates(Department{Name: department.Name}).Error
	if err != nil {
		return msg.ERROR_DEPARTMENT_UPDATE
	}
	return msg.SUCCESS
}

func SelectDepartments(department *Department) (departments []Department, code int) {
	err = db.Preload("Type").Where("office_uid = ? AND name LIKE ?", department.OfficeUID, "%"+department.Name+"%").
		Find(&departments).Error
	if err != nil {
		return departments, msg.ERROR_DEPARTMENT_SELECT
	}
	return departments, msg.SUCCESS
}

func InsertRole(role *Role) (code int) {
	role.UID = uid.Generate()
	err = db.Create(&role).Error
	if err != nil {
		return msg.ERROR_OFFICE_INSERT
	}
	return msg.SUCCESS
}

func UpdateRole(role *Role) (code int) {
	err = db.Model(&role).Association("Permissions").Replace(role.Permissions)
	if err != nil {
		return msg.ERROR_ROLE_UPDATE
	}
	return msg.SUCCESS
}

func SelectRole(uid string) (role Role, code int) {
	err = db.Preload("Permissions").Where("uid = ?", uid).First(&role).Error
	if err != nil {
		return role, msg.ERROR_ROLE_SELECT
	}
	return role, msg.SUCCESS
}

func SelectRoles() (roles []Role, code int) {
	err = db.Raw("SELECT * From role Where department_uid is NULL").Scan(&roles).Error
	if err != nil {
		return roles, msg.ERROR_ROLE_SELECT
	}
	return roles, msg.SUCCESS
}

func SelectAllRoles(name string) (roles []Role, code int) {
	err = db.Preload("Department").Where("name LIKE ?", "%"+name+"%").Find(&roles).Error
	if err != nil {
		return roles, msg.ERROR_ROLE_SELECT
	}
	return roles, msg.SUCCESS
}

func SelectPermission(module string, name string) (permission Permission, code int) {
	err = db.Where("module = ? AND name = ?", module, name).First(&permission).Error
	if err != nil {
		return permission, msg.ERROR_PERMISSION_SELECT
	}
	return permission, msg.SUCCESS
}

func SelectPermissions() (permissions []Permission, code int) {
	err = db.Find(&permissions).Error
	if err != nil {
		return permissions, msg.ERROR_PERMISSION_SELECT
	}
	return permissions, msg.SUCCESS
}

func SelectUrls(uids []string) (urls []Url) {
	// db.Raw("SELECT distinct url.* FROM url LEFT JOIN permission "+
	// 	"ON url.uid = permission.url_uid WHERE permission.uid IN (?) or url.id = 1 order by url.no", uids).
	// 	Scan(&urls)
	db.Raw("SELECT distinct url.* FROM url LEFT JOIN permission "+
		"ON url.uid = permission.url_uid WHERE permission.uid IN (?) order by url.no", uids).
		Scan(&urls)
	return
}

func SelectPermissionsNo(uids []string) (nos []string) {
	db.Raw("SELECT no FROM permission WHERE uid IN (?)", uids).Scan(&nos)
	return
}
