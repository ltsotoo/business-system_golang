package model

import (
	"business-system_golang/utils/msg"

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
	NewTargetLoad       float64 `gorm:"type:decimal(20,6);comment:新今年目标量(元)" json:"newTargetLoad"`

	Remarks string `gorm:"type:varchar(200);comment:备注" json:"remarks"`

	Office   Office   `gorm:"foreignKey:OfficeUID;references:UID" json:"office"`
	Employee Employee `gorm:"foreignKey:EmployeeUID;references:UID" json:"employee"`

	StartDate string `gorm:"-" json:"startDate"`
	EndDate   string `gorm:"-" json:"endDate"`
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

	OfficeUID string `gorm:"-" json:"officeUID"`
	StartDate string `gorm:"-" json:"startDate"`
	EndDate   string `gorm:"-" json:"endDate"`
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

func SelectHistoryOffices(pageSize int, pageNo int, historyOffice *HistoryOffice) (historyOffices []HistoryOffice, code int, total int64) {
	var maps = make(map[string]interface{})
	if historyOffice.OfficeUID != "" {
		maps["office_uid"] = historyOffice.OfficeUID
	}

	tDb := db.Where(maps)
	if historyOffice.Remarks != "" {
		tDb = tDb.Where("remarks LIKE ?", "%"+historyOffice.Remarks+"%")
	}

	if historyOffice.StartDate != "" && historyOffice.EndDate != "" {
		tDb = tDb.Where("created_at BETWEEN ? AND ?", historyOffice.StartDate, historyOffice.EndDate)
	} else {
		if historyOffice.StartDate != "" {
			tDb = tDb.Where("created_at >= ?", historyOffice.StartDate)
		}
		if historyOffice.EndDate != "" {
			tDb = tDb.Where("created_at <= ?", historyOffice.EndDate)
		}
	}

	err = tDb.Find(&historyOffices).Count(&total).
		Order("id desc").
		Preload("Office").Preload("Employee").
		Limit(pageSize).Offset((pageNo - 1) * pageSize).
		Find(&historyOffices).Error
	if err != nil {
		return historyOffices, msg.ERROR, total
	}
	return historyOffices, msg.SUCCESS, total
}

func SelectHistoryEmployees(pageSize int, pageNo int, historyEmployee *HistoryEmployee) (historyEmployees []HistoryEmployee, code int, total int64) {
	var maps = make(map[string]interface{})
	if historyEmployee.UserUID != "" {
		maps["history_employee.user_uid"] = historyEmployee.UserUID
	}

	tDb := db.Where(maps)

	if historyEmployee.OfficeUID != "" && historyEmployee.UserUID == "" {
		tDb = tDb.Joins("User").Joins("left join office on User.office_uid = office.uid").Where("office.uid = ?", historyEmployee.OfficeUID)
	}

	if historyEmployee.Remarks != "" {
		tDb = tDb.Where("history_employee.remarks LIKE ?", "%"+historyEmployee.Remarks+"%")
	}

	if historyEmployee.StartDate != "" && historyEmployee.EndDate != "" {
		tDb = tDb.Where("history_employee.created_at BETWEEN ? AND ?", historyEmployee.StartDate, historyEmployee.EndDate)
	} else {
		if historyEmployee.StartDate != "" {
			tDb = tDb.Where("history_employee.created_at >= ?", historyEmployee.StartDate)
		}
		if historyEmployee.EndDate != "" {
			tDb = tDb.Where("history_employee.created_at <= ?", historyEmployee.EndDate)
		}
	}

	err = tDb.Find(&historyEmployees).Count(&total).
		Preload("User").Preload("Employee").
		Limit(pageSize).Offset((pageNo - 1) * pageSize).
		Order("history_employee.id desc").
		Find(&historyEmployees).Error
	if err != nil {
		return historyEmployees, msg.ERROR, total
	}
	return historyEmployees, msg.SUCCESS, total
}
