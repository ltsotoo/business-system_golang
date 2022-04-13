package model

import (
	"business-system_golang/utils/magic"
	"business-system_golang/utils/msg"
	"business-system_golang/utils/uid"

	"gorm.io/gorm"
)

type Expense struct {
	BaseModel
	UID          string  `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	EmployeeUID  string  `gorm:"type:varchar(32);comment:申请员工UID;default:(-)" json:"employeeUID"`
	Type         string  `gorm:"type:varchar(32);comment:报销类型(1:补助,2:提成,3:业务费,4:差旅费)" json:"type"`
	Text         string  `gorm:"type:varchar(600);comment:申请理由" json:"text"`
	Amount       float64 `gorm:"type:decimal(20,6);comment:金额(元)" json:"amount"`
	Status       int     `gorm:"type:int;comment:状态(-1:拒绝,1:待办事处审批,2:待财务审批,3:带出纳付款,4:完成)" json:"status"`
	ApproverUID1 string  `gorm:"type:varchar(32);comment:办事处审批员工UID;default:(-)" json:"approverUID1"`
	ApproverUID2 string  `gorm:"type:varchar(32);comment:财务审批财务员工UID;default:(-)" json:"approverUID2"`
	ApproverUID3 string  `gorm:"type:varchar(32);comment:出纳员工UID;default:(-)" json:"approverUID3"`
	IsDelete     bool    `gorm:"type:boolean;comment:是否删除" json:"isDelete"`

	ExpenseType Dictionary `gorm:"foreignKey:Type;references:UID" json:"expenseType"`
	Employee    Employee   `gorm:"foreignKey:EmployeeUID;references:UID" json:"employee"`
	Approver1   Employee   `gorm:"foreignKey:ApproverUID1;references:UID" json:"approver1"`
	Approver2   Employee   `gorm:"foreignKey:ApproverUID2;references:UID" json:"approver2"`
	Approver3   Employee   `gorm:"foreignKey:ApproverUID3;references:UID" json:"approver3"`

	IsPass bool `gorm:"-" json:"isPass"`
}

type ExpenseQuery struct {
	Type          string
	Status        int
	OfficeUID     string
	EmployeeUID   string
	EmployeeName  string
	EmployeePhone string
	StartDate     string `json:"startDate"`
	EndDate       string `json:"endDate"`
}

func InsertExpense(expense *Expense) (code int) {
	expense.UID = uid.Generate()
	err = db.Create(&expense).Error
	if err != nil {
		return msg.ERROR_EXPENSE_INSERT
	}
	return msg.SUCCESS
}

func DeleteExpense(uid string) (code int) {
	err = db.Model(&Expense{}).Where("uid = ? AND status = ?", uid, -1).Update("is_delete", true).Error
	if err != nil {
		return msg.ERROR_EXPENSE_DELETE
	}
	return msg.SUCCESS
}

func ApprovalExpense(oldStatus int, expense *Expense, maps map[string]interface{}, employeeUID string, expenseDB Expense) (code int) {
	if expense.Status == magic.EXPENSE_STATUS_FAIL {
		if oldStatus == magic.EXPENSE_STATUS_NOT_PAYMENT {
			switch expense.Type {
			case magic.EXPENSE_TYPE_1:
				err = db.Transaction(func(tdb *gorm.DB) error {
					var historyEmployee HistoryEmployee
					historyEmployee.UserUID = expense.EmployeeUID
					historyEmployee.ChangeMoney = expense.Amount
					historyEmployee.Remarks = "[" + expenseDB.Employee.Office.Name + "]的员工[" + expenseDB.Employee.Name + "]发起的[补助]在待出纳付款时被驳回"
					historyEmployee.EmployeeUID = employeeUID
					if tErr := InsertHistoryEmployee(&historyEmployee, tdb); tErr != nil {
						return tErr
					}
					if tErr := tdb.Exec("UPDATE employee SET money = money + ? WHERE uid = ?", expense.Amount, expense.EmployeeUID).Error; tErr != nil {
						return tErr
					}
					if tErr := tdb.Model(&Expense{}).Where("uid = ?", expense.UID).Updates(maps).Error; tErr != nil {
						return tErr
					}
					return nil
				})
			case magic.EXPENSE_TYPE_2:
				err = db.Transaction(func(tdb *gorm.DB) error {
					var historyOffice HistoryOffice
					historyOffice.OfficeUID = expense.Employee.Office.UID
					historyOffice.ChangeMoney = expense.Amount
					historyOffice.Remarks = "[" + expenseDB.Employee.Office.Name + "]的员工[" + expenseDB.Employee.Name + "]发起的[提成]在待出纳付款时被驳回"
					historyOffice.EmployeeUID = employeeUID
					if tErr := InsertHistoryOffice(&historyOffice, tdb); tErr != nil {
						return tErr
					}
					if tErr := tdb.Exec("UPDATE office SET money = money + ? WHERE uid = ?", expense.Amount, expense.Employee.Office.UID).Error; tErr != nil {
						return tErr
					}
					if tErr := tdb.Model(&Expense{}).Where("uid = ?", expense.UID).Updates(maps).Error; tErr != nil {
						return tErr
					}
					return nil
				})
			case magic.EXPENSE_TYPE_3:
				err = db.Transaction(func(tdb *gorm.DB) error {
					var historyOffice HistoryOffice
					historyOffice.OfficeUID = expense.Employee.Office.UID
					historyOffice.ChangeBusinessMoney = expense.Amount
					historyOffice.Remarks = "[" + expenseDB.Employee.Office.Name + "]的员工[" + expenseDB.Employee.Name + "]发起的[业务费]在待出纳付款时被驳回"
					historyOffice.EmployeeUID = employeeUID
					if tErr := InsertHistoryOffice(&historyOffice, tdb); tErr != nil {
						return tErr
					}
					if tErr := tdb.Exec("UPDATE office SET business_money = business_money + ? WHERE uid = ?", expense.Amount, expense.Employee.Office.UID).Error; tErr != nil {
						return tErr
					}
					if tErr := tdb.Model(&Expense{}).Where("uid = ?", expense.UID).Updates(maps).Error; tErr != nil {
						return tErr
					}
					return nil
				})
			case magic.EXPENSE_TYPE_4:
				err = db.Model(&Expense{}).Where("uid = ?", expense.UID).Updates(maps).Error
			}
		} else {
			err = db.Model(&Expense{}).Where("uid = ?", expense.UID).Updates(maps).Error
		}
	} else if expense.Status == magic.EXPENSE_STATUS_NOT_APPROVAL_2 ||
		expense.Status == magic.EXPENSE_STATUS_FINISH {
		err = db.Model(&Expense{}).Where("uid = ?", expense.UID).Updates(maps).Error
	} else if expense.Status == magic.EXPENSE_STATUS_NOT_PAYMENT {
		switch expense.Type {
		case magic.EXPENSE_TYPE_1:
			err = db.Transaction(func(tdb *gorm.DB) error {
				var historyEmployee HistoryEmployee
				historyEmployee.UserUID = expense.EmployeeUID
				historyEmployee.ChangeMoney = -expense.Amount
				historyEmployee.Remarks = "[" + expenseDB.Employee.Office.Name + "]的员工[" + expenseDB.Employee.Name + "]发起的[补助]财务审批通过"
				historyEmployee.EmployeeUID = employeeUID
				if tErr := InsertHistoryEmployee(&historyEmployee, tdb); tErr != nil {
					return tErr
				}
				if tErr := tdb.Exec("UPDATE employee SET money = money - ? WHERE uid = ?", expense.Amount, expense.EmployeeUID).Error; tErr != nil {
					return tErr
				}
				if tErr := tdb.Model(&Expense{}).Where("uid = ?", expense.UID).Updates(maps).Error; tErr != nil {
					return tErr
				}
				return nil
			})
		case magic.EXPENSE_TYPE_2:
			err = db.Transaction(func(tdb *gorm.DB) error {
				var historyOffice HistoryOffice
				historyOffice.OfficeUID = expense.Employee.Office.UID
				historyOffice.ChangeMoney = -expense.Amount
				historyOffice.Remarks = "[" + expenseDB.Employee.Office.Name + "]的员工[" + expenseDB.Employee.Name + "]发起的[提成]财务审批通过"
				historyOffice.EmployeeUID = employeeUID
				if tErr := InsertHistoryOffice(&historyOffice, tdb); tErr != nil {
					return tErr
				}
				if tErr := tdb.Exec("UPDATE office SET money = money - ? WHERE uid = ?", expense.Amount, expense.Employee.Office.UID).Error; tErr != nil {
					return tErr
				}
				if tErr := tdb.Model(&Expense{}).Where("uid = ?", expense.UID).Updates(maps).Error; tErr != nil {
					return tErr
				}
				return nil
			})
		case magic.EXPENSE_TYPE_3:
			err = db.Transaction(func(tdb *gorm.DB) error {
				var historyOffice HistoryOffice
				historyOffice.OfficeUID = expense.Employee.Office.UID
				historyOffice.ChangeBusinessMoney = -expense.Amount
				historyOffice.Remarks = "[" + expenseDB.Employee.Office.Name + "]的员工[" + expenseDB.Employee.Name + "]发起的[业务费]财务审批通过"
				historyOffice.EmployeeUID = employeeUID
				if tErr := InsertHistoryOffice(&historyOffice, tdb); tErr != nil {
					return tErr
				}
				if tErr := tdb.Exec("UPDATE office SET business_money = business_money - ? WHERE uid = ?", expense.Amount, expense.Employee.Office.UID).Error; tErr != nil {
					return tErr
				}
				if tErr := tdb.Model(&Expense{}).Where("uid = ?", expense.UID).Updates(maps).Error; tErr != nil {
					return tErr
				}
				return nil
			})
		case magic.EXPENSE_TYPE_4:
			err = db.Model(&Expense{}).Where("uid = ?", expense.UID).Updates(maps).Error
		}
	}
	if err != nil {
		return msg.ERROR_EXPENSE_UPDATE
	}
	return msg.SUCCESS
}

func SelectExpense(uid string) (expense Expense, code int) {
	err = db.Preload("ExpenseType").Preload("Employee.Office").Preload("Approver1").Preload("Approver2").Preload("Approver3").
		Where("uid = ?", uid).Where("is_delete = ?", false).First(&expense).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return expense, msg.ERROR_EXPENSE_SELECT
	}
	return expense, msg.SUCCESS
}

func SelectExpenses(pageSize int, pageNo int, expenseQuery *ExpenseQuery) (expenses []Expense, code int, total int64) {
	var maps = make(map[string]interface{})
	if expenseQuery.Type != "" {
		maps["expense.type"] = expenseQuery.Type
	}
	if expenseQuery.Status != 0 {
		maps["expense.status"] = expenseQuery.Status
	}
	if expenseQuery.EmployeeUID != "" {
		maps["Employee.uid"] = expenseQuery.EmployeeUID
	}
	if expenseQuery.OfficeUID != "" {
		maps["Employee.office_uid"] = expenseQuery.OfficeUID
	}

	tDb := db.Joins("Employee").Where(maps).Where("expense.is_delete = ?", false)

	if expenseQuery.StartDate != "" && expenseQuery.EndDate != "" {
		tDb = tDb.Where("expense.created_at BETWEEN ? AND ?", expenseQuery.StartDate, expenseQuery.EndDate)
	} else {
		if expenseQuery.StartDate != "" {
			tDb = tDb.Where("expense.created_at >= ?", expenseQuery.StartDate)
		}
		if expenseQuery.EndDate != "" {
			tDb = tDb.Where("expense.created_at <= ?", expenseQuery.EndDate)
		}
	}

	if expenseQuery.EmployeeName != "" {
		tDb = tDb.Where("Employee.name LIKE ?", "%"+expenseQuery.EmployeeName+"%")
	}
	err = tDb.Find(&expenses).Count(&total).
		Preload("ExpenseType").Preload("Employee.Office").
		Preload("Approver1").Preload("Approver2").Preload("Approver3").
		Limit(pageSize).Offset((pageNo - 1) * pageSize).
		Order("expense.created_at desc").
		Find(&expenses).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return expenses, msg.ERROR, total
	}
	return expenses, msg.SUCCESS, total
}
