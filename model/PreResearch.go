package model

import (
	"business-system_golang/utils/msg"
	"business-system_golang/utils/uid"
	"time"

	"gorm.io/gorm"
)

type PreResearch struct {
	BaseModel
	UID         string `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	EmployeeUID string `gorm:"type:varchar(32);comment:业务员UID;default:(-)" json:"employeeUID"`
	AuditorUID  string `gorm:"type:varchar(32);comment:审核员ID;default:(-)" json:"auditorUID"`
	Remarks     string `gorm:"type:varchar(600);comment:设计需求" json:"remarks"`
	Status      int    `gorm:"type:int;comment:状态(-1:驳回 1:未审批 2:未完成 3:已完成);not null" json:"status"`
	IsDelete    bool   `gorm:"type:boolean;comment:是否删除" json:"isDelete"`

	PreResearchTasks []PreResearchTask `gorm:"-" json:"preResearchTasks"`
	Employee         Employee          `gorm:"foreignKey:EmployeeUID;references:UID" json:"employee"`
	Auditor          Employee          `gorm:"foreignKey:AuditorUID;references:UID" json:"auditor"`
}

type PreResearchTask struct {
	BaseModel
	UID            string `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	PreResearchUID string `gorm:"type:varchar(32);comment:预研UID;not null" json:"preResearchUID"`
	EmployeeUID    string `gorm:"type:varchar(32);comment:技术负责人UID;default:(-)" json:"employeeUID"`
	AuditorUID     string `gorm:"type:varchar(32);comment:审核员ID;default:(-)" json:"auditorUID"`
	Requirement    string `gorm:"type:varchar(600);comment:设计要求" json:"requirement"`
	Remarks        string `gorm:"type:varchar(600);comment:备注" json:"remarks"`
	Days           int    `gorm:"type:int;comment:分配工作天数" json:"days"`
	StartDate      XTime  `gorm:"type:datetime;comment:开始工作日期" json:"startDate"`
	EndDate        XTime  `gorm:"type:datetime;comment:预计结束工作日期" json:"endDate"`
	RealEndDate    XTime  `gorm:"type:datetime;comment:实际结束工作日期;default:(-)" json:"realEndDate"`
	Status         int    `gorm:"type:int;comment:状态( 1:未完成 2:未审核 3:未通过 4:已通过)" json:"status"`

	Employee Employee `gorm:"foreignKey:EmployeeUID;references:UID" json:"employee"`
	Auditor  Employee `gorm:"foreignKey:AuditorUID;references:UID" json:"auditor"`
}

type PreResearchQuery struct {
	UID          string `json:"UID"`
	AuditorUID   string `json:"auditorUID"`
	EmployeeUID  string `json:"employeeUID"`
	EmployeeName string `json:"employeeName"`
	Status       int    `json:"status"`
	Days         int    `json:"days"`
	Requirement  string `json:"requirement"`
}

func InsertPreResearch(preResearch *PreResearch) (code int) {
	preResearch.UID = uid.Generate()
	preResearch.Status = 1
	err = db.Create(&preResearch).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func DeletePreResearch(uid string) (code int) {
	err = db.Model(&PreResearch{}).Where("uid = ? And status = ?", uid, 1).Update("is_delete", true).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func SelectPreReasearch(uid string) (preResearch PreResearch, code int) {
	err = db.Preload("Employee.Office").Preload("Auditor").
		Where("is_delete = ?", false).First(&preResearch, "uid = ?", uid).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return preResearch, msg.ERROR
	}
	var preResearchTasks []PreResearchTask
	err = db.Preload("Employee").Preload("Auditor").Where("pre_research_uid = ?", uid).Find(&preResearchTasks).Error
	if err == nil {
		preResearch.PreResearchTasks = preResearchTasks
	}
	return preResearch, msg.SUCCESS
}

func SelectPreReasearchs(pageSize int, pageNo int, preResearchQuery *PreResearchQuery) (preResearchs []PreResearch, code int, total int64) {
	var maps = make(map[string]interface{})
	maps["pre_research.is_delete"] = false
	if preResearchQuery.Status != 0 {
		maps["pre_research.status"] = preResearchQuery.Status
	}
	if preResearchQuery.EmployeeUID != "" {
		maps["pre_research.employee_uid"] = preResearchQuery.EmployeeUID
	}

	tDb := db.Joins("Employee")
	if preResearchQuery.EmployeeName != "" {
		tDb = tDb.Where("Employee.name Like ?", "%"+preResearchQuery.EmployeeName+"%")
	}

	err = tDb.Where(maps).
		Preload("Employee.Office").Preload("Auditor").Find(&preResearchs).Count(&total).
		Limit(pageSize).Offset((pageNo - 1) * pageSize).
		Order("pre_research.created_at desc").
		Find(&preResearchs).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return preResearchs, msg.ERROR, total
	}
	return preResearchs, msg.SUCCESS, total
}

func SelectPreReasearchTask(uid string) (preResearchTask PreResearchTask, code int) {
	err = db.Preload("Employee").Preload("Auditor").First(&preResearchTask, "uid = ?", uid).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return preResearchTask, msg.ERROR
	}
	return preResearchTask, msg.SUCCESS
}

func SelectPreReasearchTasks(pageSize int, pageNo int, preResearchQuery *PreResearchQuery) (preResearchTasks []PreResearchTask, code int, total int64) {
	var maps = make(map[string]interface{})
	if preResearchQuery.Status != 0 {
		maps["status"] = preResearchQuery.Status
	}
	if preResearchQuery.EmployeeUID != "" {
		maps["pre_research_task.employee_uid"] = preResearchQuery.EmployeeUID
	}
	tDb := db.Joins("Employee")
	if preResearchQuery.EmployeeName != "" {
		tDb = tDb.Where("Employee.name Like ?", "%"+preResearchQuery.EmployeeName+"%")
	}
	err = tDb.Where(maps).
		Preload("Auditor").Find(&preResearchTasks).Count(&total).
		Limit(pageSize).Offset((pageNo - 1) * pageSize).
		Order("pre_research_task.created_at desc").
		Find(&preResearchTasks).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return preResearchTasks, msg.ERROR, total
	}
	return preResearchTasks, msg.SUCCESS, total
}

func UpdatePreResearchStatus(preResearchQuery *PreResearchQuery) (code int) {
	var maps = make(map[string]interface{})
	maps["status"] = preResearchQuery.Status
	maps["auditor_uid"] = preResearchQuery.AuditorUID

	if preResearchQuery.Status == 2 {

		var preResearchTask PreResearchTask
		preResearchTask.UID = uid.Generate()
		preResearchTask.Status = 1
		preResearchTask.PreResearchUID = preResearchQuery.UID
		preResearchTask.AuditorUID = preResearchQuery.AuditorUID
		preResearchTask.EmployeeUID = preResearchQuery.EmployeeUID
		preResearchTask.Days = preResearchQuery.Days
		preResearchTask.Requirement = preResearchQuery.Requirement

		t1 := time.Now()
		t2 := t1.AddDate(0, 0, preResearchTask.Days)

		xt1 := &XTime{t1}
		xt2 := &XTime{t2}

		preResearchTask.StartDate = *xt1
		preResearchTask.EndDate = *xt2

		err = db.Transaction(func(tdb *gorm.DB) error {
			if tErr := tdb.Model(&PreResearch{}).Where("uid = ?", preResearchQuery.UID).Updates(maps).Error; tErr != nil {
				return tErr
			}
			if tErr := tdb.Create(&preResearchTask).Error; tErr != nil {
				return tErr
			}
			return nil
		})
	} else {
		err = db.Model(&PreResearch{}).Where("uid = ?", preResearchQuery.UID).Updates(maps).Error
	}

	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func UpdatePreResearchTaskStatus(preResearchTask *PreResearchTask) (code int) {
	var maps = make(map[string]interface{})
	maps["status"] = preResearchTask.Status
	maps["auditor_uid"] = preResearchTask.AuditorUID
	if preResearchTask.Status == 2 {
		maps["remarks"] = preResearchTask.Remarks
		t := time.Now().Format("2006-01-02 15:04:05")
		maps["real_end_date"] = t

		err = db.Model(&PreResearchTask{}).Where("uid = ?", preResearchTask.UID).Updates(maps).Error
	} else if preResearchTask.Status == 3 {

		var newPreResearchTask PreResearchTask
		newPreResearchTask.UID = uid.Generate()
		newPreResearchTask.Status = 1
		newPreResearchTask.PreResearchUID = preResearchTask.PreResearchUID
		newPreResearchTask.AuditorUID = preResearchTask.AuditorUID
		newPreResearchTask.EmployeeUID = preResearchTask.EmployeeUID
		newPreResearchTask.Days = preResearchTask.Days
		newPreResearchTask.Requirement = preResearchTask.Requirement

		t1 := time.Now()
		t2 := t1.AddDate(0, 0, preResearchTask.Days)

		xt1 := &XTime{t1}
		xt2 := &XTime{t2}

		newPreResearchTask.StartDate = *xt1
		newPreResearchTask.EndDate = *xt2

		err = db.Transaction(func(tdb *gorm.DB) error {
			if tErr := tdb.Debug().Model(&PreResearchTask{}).Where("uid = ?", preResearchTask.UID).Updates(maps).Error; tErr != nil {
				return tErr
			}
			if tErr := tdb.Debug().Create(&newPreResearchTask).Error; tErr != nil {
				return tErr
			}
			return nil
		})
	} else if preResearchTask.Status == 4 {
		err = db.Transaction(func(tdb *gorm.DB) error {
			if tErr := tdb.Debug().Model(&PreResearchTask{}).Where("uid = ?", preResearchTask.UID).Updates(maps).Error; tErr != nil {
				return tErr
			}
			if tErr := tdb.Debug().Model(&PreResearch{}).Where("uid = ?", preResearchTask.PreResearchUID).Update("status", 3).Error; tErr != nil {
				return tErr
			}
			return nil
		})
	}
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}
