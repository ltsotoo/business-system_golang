package model

import (
	"business-system_golang/utils/msg"

	"gorm.io/gorm"
)

// 合同 Model
type Contract struct {
	gorm.Model
	No                    string `gorm:"type:varchar(20);comment:合同编号" json:"no"`
	AreaID                uint   `gorm:"type:int;comment:所属区域ID;default:(-)" json:"areaID"`
	EmployeeID            uint   `gorm:"type:int;comment:业务员ID;default:(-)" json:"employeeID"`
	IsEntryCustomer       bool   `gorm:"type:boolean;comment:客户是否录入" json:"isEntryCustomer"`
	CustomerID            uint   `gorm:"type:int;comment:客户ID;default:(-)" json:"customerID"`
	ContractDate          string `gorm:"type:varchar(20);comment:签订日期" json:"contractDate"`
	ContractUnitID        uint   `gorm:"type:int;comment:签订单位;default:(-)" json:"contractUnitID"`
	EstimatedDeliveryDate string `gorm:"type:varchar(20);comment:合同交货日期" json:"estimatedDeliveryDate"`
	EndDeliveryDate       string `gorm:"type:varchar(20);comment:实际交货日期" json:"endDeliveryDate"`
	InvoiceType           int    `gorm:"type:int;comment:开票类型" json:"invoiceType"`
	InvoiceContent        string `gorm:"type:varchar(20);comment:开票内容" json:"invoiceContent"`
	IsSpecial             bool   `gorm:"type:boolean;comment:特殊合同?" json:"isSpecial"`
	TotalAmount           int    `gorm:"type:int;comment:总金额(元)" json:"totalAmount"`
	Remarks               string `gorm:"type:varchar(20);comment:备注" json:"remarks"`
	Status                int    `gorm:"type:int;comment:状态;not null" json:"status"`

	Area         Area       `gorm:"foreignKey:AreaID" json:"area"`
	Employee     Employee   `gorm:"foreignKey:EmployeeID" json:"employee"`
	Customer     Customer   `gorm:"foreignKey:CustomerID" json:"customer"`
	ContractUnit Dictionary `gorm:"foreignKey:ContractUnitID" json:"contractUnit"`
	Tasks        []Task     `gorm:"foreignKey:ContractID" json:"tasks"`
}

func CreateContract(contract *Contract) (code int) {
	for _, task := range contract.Tasks {
		contract.TotalAmount += task.TotalPrice
	}
	if contract.IsEntryCustomer {
		contract.Customer = Customer{}
	}
	err = db.Create(&contract).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func DeleteContract(id int) (code int) {
	var contract Contract
	contract.ID = uint(id)
	err = db.Select("Tasks").Delete(&contract).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func UpdateContract(contract *Contract) (code int) {
	var maps = make(map[string]interface{})
	maps["Remarks"] = contract.Remarks
	err = db.Model(&contract).Updates(maps).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func SelectContract(id int) (contract Contract, code int) {
	err = db.Preload("Area").Preload("Employee").Preload("Customer").Preload("ContractUnit").First(&contract, id).Error
	contract.Area.Office, _ = SelectOffice(int(contract.Area.OfficeID))
	contract.Customer.Company, _ = SelectCompany(int(contract.Customer.CompanyID))
	contract.Tasks, _ = SelectTaskByContractID(int(contract.ID))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return contract, msg.ERROR_CONTRACT_NOT_EXIST
		} else {
			return contract, msg.ERROR
		}
	}
	return contract, msg.SUCCESS
}

func SelectContracts(pageSize int, pageNo int, contractQuery ContractQuery) (contracts []Contract, code int, total int64) {
	err = db.Limit(pageSize).Offset((pageNo-1)*pageSize).Preload("Area").Preload("Employee").Preload("Customer").Where("no LIKE ?", "%"+contractQuery.No+"%").Find(&contracts).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, msg.ERROR, 0
	}
	db.Model(&contracts).Count(&total)
	return contracts, msg.SUCCESS, total
}
