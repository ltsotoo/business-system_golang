package model

import (
	"business-system_golang/utils/magic"
	"business-system_golang/utils/msg"
	"business-system_golang/utils/uid"

	"gorm.io/gorm"
)

type Expense struct {
	BaseModel
	UID         string  `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	EmployeeUID string  `gorm:"type:varchar(32);comment:申请员工UID;default:(-)" json:"employeeUID"`
	Type        int     `gorm:"type:int;comment:部门类型(1:个人,2:办事处)" json:"type"`
	Text        string  `gorm:"type:varchar(600);comment:申请理由" json:"text"`
	Amount      float64 `gorm:"type:decimal(20,6);comment:金额(元)" json:"amount"`
	Status      int     `gorm:"type:int;comment:状态(-1:拒绝,1:待办事处审批,2:待财务审批,3:通过)" json:"status"`
	ApproverUID string  `gorm:"type:varchar(32);comment:审批财务员工UID;default:(-)" json:"approverUID"`
	IsDelete    bool    `gorm:"type:boolean;comment:是否删除" json:"isDelete"`

	Employee Employee `gorm:"foreignKey:EmployeeUID;references:UID" json:"employee"`
	Approver Employee `gorm:"foreignKey:ApproverUID;references:UID" json:"approver"`
}

type ExpenseQuery struct {
	Type          int
	OfficeUID     string
	EmployeeName  string
	EmployeePhone string
}

func InsertExpense(expense *Expense) (code int) {
	expense.UID = uid.Generate()
	err = db.Create(&expense).Error
	if err != nil {
		return msg.ERROR_EXPENSE_INSERT
	}
	return msg.SUCCESS
}

func UpdateExpense(expense *Expense) (code int) {
	var maps = make(map[string]interface{})
	maps["approver_uid"] = expense.ApproverUID
	maps["status"] = expense.Status
	err = db.Model(&Expense{}).Where("uid = ?", expense.UID).Updates(maps).Error
	if err != nil {
		return msg.ERROR_EXPENSE_UPDATE
	}
	return msg.SUCCESS
}

func UpdateMoneyExpense(expense *Expense) (code int) {
	var maps = make(map[string]interface{})
	maps["approver_uid"] = expense.ApproverUID
	maps["status"] = expense.Status
	if expense.Status == magic.EXPENSE_STATUS_FAIL ||
		expense.Status == magic.EXPENSE_STATUS_NOT_APPROVAL_2 {
		err = db.Model(&Expense{}).Where("uid = ?", expense.UID).Updates(maps).Error
	}
	if expense.Status == magic.EXPENSE_STATUS_PASS {
		if expense.Type == magic.EXPENSE_TYPE_OFFICE {
			office, _ := SelectOffice(expense.Employee.OfficeUID)
			office.Money = office.Money - expense.Amount
			err = db.Transaction(func(tdb *gorm.DB) error {
				if tErr := tdb.Model(&Expense{}).Where("uid = ?", expense.UID).Updates(maps).Error; tErr != nil {
					return tErr
				}
				if tErr := tdb.Model(&Office{}).Where("uid = ?", office.UID).Update("money", office.Money).Error; tErr != nil {
					return tErr
				}
				return nil
			})
		}
		if expense.Type == magic.EXPENSE_TYPE_EMPLOYEE {
			employee, _ := SelectEmployee(expense.EmployeeUID)
			if employee.Money > expense.Amount {
				employee.Money = employee.Money - expense.Amount
				err = db.Transaction(func(tdb *gorm.DB) error {
					if tErr := tdb.Model(&Expense{}).Where("uid = ?", expense.UID).Updates(maps).Error; tErr != nil {
						return tErr
					}
					if tErr := tdb.Model(&Employee{}).Where("uid = ?", employee.UID).Update("money", employee.Money).Error; tErr != nil {
						return tErr
					}
					return nil
				})
			} else {
				return msg.ERROR_EXPENSE_EMPLOYEE_MONEY_LESS
			}
		}
	}
	if err != nil {
		return msg.ERROR_EXPENSE_UPDATE
	}
	return msg.SUCCESS
}

func SelectExpense(uid string) (expense Expense, code int) {
	err = db.Preload("Employee").Preload("Approver").
		Where("uid = ?", uid).First(&expense).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return expense, msg.ERROR_EXPENSE_SELECT
	}
	return expense, msg.SUCCESS
}

func SelectBaseExpense(uid string) (expense Expense, code int) {
	err = db.Where("uid = ?", uid).First(&expense).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return expense, msg.ERROR_EXPENSE_SELECT
	}
	return expense, msg.SUCCESS
}

func SelectExpenses(pageSize int, pageNo int, expenseQuery *ExpenseQuery) (expenses []Expense, code int, total int64) {
	var maps = make(map[string]interface{})
	if expenseQuery.Type != 0 {
		maps["expense.type"] = expenseQuery.Type
	}
	if expenseQuery.OfficeUID != "" {
		maps["Employee.office_uid"] = expenseQuery.OfficeUID
	}

	err = db.Joins("Employee").Where(maps).
		Where("Employee.name LIKE ? AND Employee.phone LIKE ?",
			"%"+expenseQuery.EmployeeName+"%", "%"+expenseQuery.EmployeePhone+"%").
		Find(&expenses).Count(&total).
		Preload("Employee.Office").Preload("Approver").
		Limit(pageSize).Offset((pageNo - 1) * pageSize).
		Find(&expenses).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return expenses, msg.ERROR, total
	}
	return expenses, msg.SUCCESS, total
}

func SelectMyExpenses(pageSize int, pageNo int, expense *Expense) (expenses []Expense, code int, total int64) {
	var maps = make(map[string]interface{})
	maps["employee_uid"] = expense.EmployeeUID
	if expense.Type != 0 {
		maps["type"] = expense.Type
	}
	if expense.Status != 0 {
		maps["status"] = expense.Status
	}
	err = db.Joins("Employee").Where(maps).
		Find(&expenses).Count(&total).
		Preload("Approver").
		Limit(pageSize).Offset((pageNo - 1) * pageSize).
		Find(&expenses).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return expenses, msg.ERROR, total
	}
	return expenses, msg.SUCCESS, total
}
