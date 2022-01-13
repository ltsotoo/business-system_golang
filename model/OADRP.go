package model

import (
	"business-system_golang/utils/msg"
	"business-system_golang/utils/uid"

	"gorm.io/gorm"
)

//Office办事处 Department部门 Role角色 Permission权限
type Office struct {
	ID            uint    `gorm:"primary_key" json:"ID"`
	UID           string  `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	Name          string  `gorm:"type:varchar(50);comment:名称;not null" json:"name"`
	Number        string  `gorm:"type:varchar(50);comment:编号" json:"number"`
	BusinessMoney float64 `gorm:"type:decimal(20,6);comment:业务费用(元)" json:"businessMoney"`
	Money         float64 `gorm:"type:decimal(20,6);comment:办事处目前可报销额度(元)" json:"money"`
	MoneyCold     float64 `gorm:"type:decimal(20,6);comment:办事处今年冻结报销额度(元)" json:"moneyCold"`
	TaskLoad      float64 `gorm:"type:decimal(20,6);comment:今年目标量(元)" json:"taskLoad"`
	TargetLoad    float64 `gorm:"type:decimal(20,6);comment:今年完成量(元)" json:"targetLoad"`

	NewBusinessMoney float64 `gorm:"type:decimal(20,6);comment:业务费用(元)" json:"newBusinessMoney"`
	NewMoney         float64 `gorm:"type:decimal(20,6);comment:办事处目前可报销额度(元)" json:"newMoney"`
	NewMoneyCold     float64 `gorm:"type:decimal(20,6);comment:办事处今年冻结报销额度(元)" json:"newMoneyCold"`
	NewTaskLoad      float64 `gorm:"type:decimal(20,6);comment:今年目标量(元)" json:"newTaskLoad"`
	NewTargetLoad    float64 `gorm:"type:decimal(20,6);comment:今年完成量(元)" json:"newTargetLoad"`

	IsSubmit bool `gorm:"type:boolean;comment:今年结算是否提交" json:"isSubmit"`

	IsDelete bool `gorm:"type:boolean;comment:是否删除" json:"isDelete"`

	FinalPercentages float64 `gorm:"-" json:"finalPercentages"`
	NotPayment       float64 `gorm:"-" json:"notPayment"`
}

type Department struct {
	ID        uint   `gorm:"primary_key" json:"ID"`
	UID       string `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	OfficeUID string `gorm:"type:varchar(32);comment:办事处ID;not null" json:"officeUID"`
	Name      string `gorm:"type:varchar(50);comment:名称;not null" json:"name"`
	RoleUID   string `gorm:"type:varchar(32);comment:部门员工默认拥有职位;default:(-)" json:"roleUID"`
	IsDelete  bool   `gorm:"type:boolean;comment:是否删除" json:"isDelete"`

	Office Office `gorm:"foreignKey:OfficeUID;references:UID" json:"office"`
	Role   Role   `gorm:"foreignKey:RoleUID;references:UID" json:"role"`
}

type Role struct {
	ID   uint   `gorm:"primary_key" json:"ID"`
	UID  string `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	Name string `gorm:"type:varchar(50);comment:名称;not null" json:"name"`

	Permissions []Permission `gorm:"many2many:role_permission;foreignKey:UID;References:UID" json:"permissions"`
}

type Permission struct {
	ID     uint   `gorm:"primary_key" json:"ID"`
	UID    string `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	Text   string `gorm:"type:varchar(99);comment:文本;not null" json:"text"`
	No     string `gorm:"type:varchar(9);comment:序号" json:"no"`
	UrlUID string `gorm:"type:varchar(32);comment:Url_UID;default:(-)" json:"urlUID"`
}

type Url struct {
	ID    uint   `gorm:"primary_key" json:"ID"`
	UID   string `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	Title string `gorm:"type:varchar(20);comment:标题;not null" json:"title"`
	Icon  string `gorm:"type:varchar(20);comment:图标" json:"icon"`
	Url   string `gorm:"type:varchar(20);comment:url" json:"url"`
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
	var maps = make(map[string]interface{})
	maps["number"] = office.Number
	maps["name"] = office.Name
	maps["task_load"] = office.TaskLoad
	maps["target_load"] = office.TargetLoad
	err = db.Model(&Office{}).Where("uid = ?", office.UID).Updates(maps).Error
	if err != nil {
		return msg.ERROR_OFFICE_UPDATE
	}
	return msg.SUCCESS
}

func UpdateOfficeMoney(office *Office) (code int) {
	err = db.Exec("UPDATE office SET money = money + ?, money_cold = money_cold + ?, business_money = business_money + ? WHERE uid = ?", office.Money, office.MoneyCold, office.BusinessMoney, office.UID).Error
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

	if office.Name != "" {
		err = db.Where("name LIKE ? AND is_delete is false", "%"+office.Name+"%").Find(&offices).Error
	} else {
		err = db.Where("is_delete is false").Find(&offices).Error
	}

	if err != nil {
		return offices, msg.ERROR_OFFICE_SELECT
	}
	return offices, msg.SUCCESS
}

func SelectTopList() (offices1 []Office, offices2 []Office) {
	db.Raw("SELECT office_uid uid,sum(total_amount - payment_total_amount) money FROM contract WHERE is_delete IS FALSE AND is_pre_deposit IS FALSE AND STATUS > 1 GROUP BY office_uid").Scan(&offices1)
	db.Raw("SELECT office_uid uid,sum(pre_deposit_record - payment_total_amount) money FROM contract WHERE is_delete IS FALSE AND is_pre_deposit IS TRUE AND STATUS > 1 GROUP BY office_uid").Scan(&offices2)
	return
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
	var maps = make(map[string]interface{})
	maps["name"] = department.Name
	if department.RoleUID != "" {
		maps["role_uid"] = department.RoleUID
	} else {
		maps["role_uid"] = nil
	}
	err = db.Model(&Department{}).Where("uid = ?", department.UID).Updates(maps).Error
	if err != nil {
		return msg.ERROR_DEPARTMENT_UPDATE
	}
	return msg.SUCCESS
}

func SelectDepartments(department *Department) (departments []Department, code int) {

	var maps = make(map[string]interface{})
	if department.OfficeUID != "" {
		maps["office_uid"] = department.OfficeUID
	}

	tdb := db.Preload("Office").Preload("Role").Where(maps)

	if department.Name != "" {
		tdb = tdb.Where("name LIKE ?", "%"+department.Name+"%")
	}

	err = tdb.Find(&departments).Error
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

	err = db.Transaction(func(tdb *gorm.DB) error {
		if tErr := tdb.Model(&Role{}).Where("uid = ?", role.UID).Update("name", role.Name).Error; tErr != nil {
			return tErr
		}
		if tErr := tdb.Model(&role).Association("Permissions").Replace(role.Permissions); tErr != nil {
			return tErr
		}
		return nil
	})

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
	err = db.Raw("SELECT * From role").Scan(&roles).Error
	if err != nil {
		return roles, msg.ERROR_ROLE_SELECT
	}
	return roles, msg.SUCCESS
}

func SelectAllRoles(name string) (roles []Role, code int) {
	if name != "" {
		err = db.Where("name LIKE ?", "%"+name+"%").Find(&roles).Error
	} else {
		err = db.Find(&roles).Error
	}

	if err != nil {
		return roles, msg.ERROR_ROLE_SELECT
	}
	return roles, msg.SUCCESS
}

func SelectPermissions() (permissions []Permission, code int) {
	err = db.Find(&permissions).Error
	if err != nil {
		return permissions, msg.ERROR_PERMISSION_SELECT
	}
	return permissions, msg.SUCCESS
}

func SelectUrls(uids []string) (urls []Url) {
	db.Raw("SELECT distinct url.* FROM url LEFT JOIN permission "+
		"ON url.uid = permission.url_uid WHERE permission.uid IN (?) order by url.id", uids).
		Scan(&urls)
	return
}

func SelectPermissionsNo(uids []string) (nos []string) {
	db.Raw("SELECT no FROM permission WHERE uid IN (?)", uids).Scan(&nos)
	return
}
