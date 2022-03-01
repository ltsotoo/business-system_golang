package model

import (
	"gorm.io/gorm"
)

type History struct {
}

type HistoryOffice struct {
	BaseModel
	OfficeUID           string  `gorm:"type:varchar(32);comment:办事处UID;default:(-)" json:"officeUID"`
	EmployeeUID         string  `gorm:"type:varchar(32);comment:操作人UID;default:(-)" json:"employeeUID"`
	OldBusinessMoney    float64 `gorm:"type:decimal(20,6);comment:旧业务费用(元)" json:"oldBusinessMoney"`
	OldMoney            float64 `gorm:"type:decimal(20,6);comment:旧办事处目前可报销额度(元)" json:"oldMoney"`
	OldMoneyCold        float64 `gorm:"type:decimal(20,6);comment:旧办事处今年冻结报销额度(元)" json:"oldMoneyCold"`
	OldTargetLoad       float64 `gorm:"type:decimal(20,6);comment:旧今年完成量(元)" json:"oldTargetLoad"`
	ChangeBusinessMoney float64 `gorm:"type:decimal(20,6);comment:修改业务费用(元)" json:"changeBusinessMoney"`
	ChangeMoney         float64 `gorm:"type:decimal(20,6);comment:修改办事处目前可报销额度(元)" json:"changeMoney"`
	ChangeMoneyCold     float64 `gorm:"type:decimal(20,6);comment:修改办事处今年冻结报销额度(元)" json:"changeMoneyCold"`
	ChangeTargetLoad    float64 `gorm:"type:decimal(20,6);comment:修改今年目标量(元)" json:"changeTargetLoad"`
	NewBusinessMoney    float64 `gorm:"type:decimal(20,6);comment:新业务费用(元)" json:"newBusinessMoney"`
	NewMoney            float64 `gorm:"type:decimal(20,6);comment:新办事处目前可报销额度(元)" json:"newMoney"`
	NewMoneyCold        float64 `gorm:"type:decimal(20,6);comment:新办事处今年冻结报销额度(元)" json:"newMoneyCold"`
	NewTargetLoad       float64 `gorm:"type:decimal(20,6);comment:新今年目标量(元)" json:"newtTargetLoad"`

	Remarks string `gorm:"type:varchar(200);comment:备注" json:"remarks"`

	Office   Office   `gorm:"foreignKey:OfficeUID;references:UID" json:"office"`
	Employee Employee `gorm:"foreignKey:EmployeeUID;references:UID" json:"employee"`
}

type HistoryEmployee struct {
	BaseModel
	EmployeeUID string  `gorm:"type:varchar(32);comment:操作人UID;default:(-)" json:"employeeUID"`
	UserUID     string  `gorm:"type:varchar(32);comment:员工UID;default:(-)" json:"userUID"`
	OldMoney    float64 `gorm:"type:decimal(20,6);comment:原补助额度(元)" json:"oldMoney"`
	ChangeMoney float64 `gorm:"type:decimal(20,6);comment:修改补助额度(元)" json:"changeMoney"`
	NewMoney    float64 `gorm:"type:decimal(20,6);comment:新补助额度(元)" json:"newMoney"`

	Remarks string `gorm:"type:varchar(200);comment:备注" json:"remarks"`

	Employee Employee `gorm:"foreignKey:EmployeeUID;references:UID" json:"employee"`
	User     Employee `gorm:"foreignKey:UserUID;references:UID" json:"user"`
}

func InsertHistoryOffice(historyOffice *HistoryOffice, tdb *gorm.DB) (err error) {
	var office Office
	err = tdb.First(&office, "uid = ?", historyOffice.OfficeUID).Error
	if err != nil {
		return err
	}
	historyOffice.OldBusinessMoney = office.BusinessMoney
	historyOffice.OldMoney = office.Money
	historyOffice.OldMoneyCold = office.MoneyCold
	historyOffice.OldTargetLoad = office.TargetLoad

	historyOffice.NewBusinessMoney = historyOffice.OldBusinessMoney + historyOffice.ChangeBusinessMoney
	historyOffice.NewMoney = historyOffice.OldMoney + historyOffice.ChangeMoney
	historyOffice.NewMoneyCold = historyOffice.OldMoneyCold + historyOffice.ChangeMoneyCold
	historyOffice.NewTargetLoad = historyOffice.OldTargetLoad + historyOffice.ChangeTargetLoad

	err = tdb.Create(&historyOffice).Error
	return err
}

func InsertHistoryEmployee(historyEmployee *HistoryEmployee, tdb *gorm.DB) (err error) {
	var user Employee
	err = tdb.First(&user, "uid = ?", historyEmployee.UserUID).Error
	if err != nil {
		return err
	}
	historyEmployee.OldMoney = user.Money

	historyEmployee.NewMoney = historyEmployee.OldMoney + historyEmployee.ChangeMoney

	err = tdb.Create(&historyEmployee).Error
	return err
}

//直接修改
func InsertHistoryEmployee2(historyEmployee *HistoryEmployee, tdb *gorm.DB) (err error) {
	var user Employee
	err = tdb.First(&user, "uid = ?", historyEmployee.UserUID).Error
	if err != nil {
		return err
	}
	historyEmployee.OldMoney = user.Money
	err = tdb.Create(&historyEmployee).Error
	return err
}
