package model

import "gorm.io/gorm"

type Expense struct {
	gorm.Model
	UID         string `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	EmployeeUID string `gorm:"type:varchar(32);comment:申请员工UID;default:(-)" json:"employeeUID"`
	TypeUID     string `gorm:"type:varchar(32);comment:部门类型;not null" json:"typeUID"`
	Text        string `gorm:"type:varchar(200);comment:申请理由" json:"text"`
	Amount      int    `gorm:"type:int;comment:金额(元)" json:"totalAmount"`
	Status      int    `gorm:"type:int;comment:状态(-1:拒绝,1:通过);not null" json:"status"`
	ApproverUID string `gorm:"type:varchar(32);comment:审批财务员工UID;default:(-)" json:"approverUID"`

	Employee Employee `gorm:"foreignKey:EmployeeUID;references:UID" json:"employee"`
	Approver Employee `gorm:"foreignKey:ApproverUID;references:UID" json:"approver"`
}
