package model

import (
	"business-system_golang/utils/msg"
	"business-system_golang/utils/uid"

	"gorm.io/gorm"
)

type Expense struct {
	BaseModel
	UID         string `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	EmployeeUID string `gorm:"type:varchar(32);comment:申请员工UID;default:(-)" json:"employeeUID"`
	TypeUID     string `gorm:"type:varchar(32);comment:部门类型;not null" json:"typeUID"`
	Text        string `gorm:"type:varchar(200);comment:申请理由" json:"text"`
	Amount      int    `gorm:"type:int;comment:金额(元)" json:"amount"`
	Status      int    `gorm:"type:int;comment:状态(-1:拒绝,1:通过);not null" json:"status"`
	ApproverUID string `gorm:"type:varchar(32);comment:审批财务员工UID;default:(-)" json:"approverUID"`

	Type     Dictionary `gorm:"foreignKey:TypeUID;references:UID" json:"type"`
	Employee Employee   `gorm:"foreignKey:EmployeeUID;references:UID" json:"employee"`
	Approver Employee   `gorm:"foreignKey:ApproverUID;references:UID" json:"approver"`
}

type ExpenseQuery struct {
	TypeUID       string
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

func SelectExpense(uid string) (expense Expense, code int) {
	err = db.Preload("Employee").Preload("Approver").Preload("Type").
		First(&expense, "uid = ?", uid).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return expense, msg.ERROR_EXPENSE_SELECT
	}
	return expense, msg.SUCCESS
}

func SelectExpenses(pageSize int, pageNo int, expenseQuery *ExpenseQuery) (expenses []Expense, code int, total int64) {
	var maps = make(map[string]interface{})
	if expenseQuery.TypeUID != "" {
		maps["expense.type_uid"] = expenseQuery.TypeUID
	}
	if expenseQuery.OfficeUID != "" {
		maps["Employee.office_uid"] = expenseQuery.OfficeUID
	}

	err = db.Joins("Employee").Where(maps).
		Where("Employee.name LIKE ? AND Employee.phone LIKE ?",
			"%"+expenseQuery.EmployeeName+"%", "%"+expenseQuery.EmployeePhone+"%").
		Find(&expenses).Count(&total).
		Preload("Employee").Preload("Employee.Office").Preload("Approver").Preload("Type").
		Limit(pageSize).Offset((pageNo - 1) * pageSize).
		Find(&expenses).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return expenses, msg.ERROR, total
	}
	return expenses, msg.SUCCESS, total
}
