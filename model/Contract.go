package model

import (
	"business-system_golang/utils/msg"
	"business-system_golang/utils/uid"

	"gorm.io/gorm"
)

// 合同 Model
type Contract struct {
	gorm.Model
	No                    string `gorm:"type:varchar(32);comment:合同编号" json:"no"`
	UID                   string `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	AreaUID               string `gorm:"type:varchar(32);comment:所属区域UID;default:(-)" json:"areaUID"`
	EmployeeUID           string `gorm:"type:varchar(32);comment:业务员UID;default:(-)" json:"employeeUID"`
	IsEntryCustomer       bool   `gorm:"type:boolean;comment:客户是否录入" json:"isEntryCustomer"`
	CustomerUID           string `gorm:"type:varchar(32);comment:客户ID;default:(-)" json:"customerUID"`
	ContractDate          string `gorm:"type:varchar(20);comment:签订日期" json:"contractDate"`
	ContractUnitUID       string `gorm:"type:varchar(32);comment:签订单位;default:(-)" json:"contractUnitUID"`
	EstimatedDeliveryDate string `gorm:"type:varchar(20);comment:合同交货日期" json:"estimatedDeliveryDate"`
	EndDeliveryDate       string `gorm:"type:varchar(20);comment:实际交货日期" json:"endDeliveryDate"`
	InvoiceType           int    `gorm:"type:int;comment:开票类型" json:"invoiceType"`
	InvoiceContent        string `gorm:"type:varchar(20);comment:开票内容" json:"invoiceContent"`
	IsSpecial             bool   `gorm:"type:boolean;comment:特殊合同?" json:"isSpecial"`
	TotalAmount           int    `gorm:"type:int;comment:总金额(元)" json:"totalAmount"`
	Remarks               string `gorm:"type:varchar(200);comment:备注" json:"remarks"`
	Status                int    `gorm:"type:int;comment:状态;not null" json:"status"`

	Area         Area       `gorm:"foreignKey:AreaUID;references:UID" json:"area"`
	Employee     Employee   `gorm:"foreignKey:EmployeeUID;references:UID" json:"employee"`
	Customer     Customer   `gorm:"foreignKey:CustomerUID;references:UID" json:"customer"`
	ContractUnit Dictionary `gorm:"foreignKey:ContractUnitUID;references:UID" json:"contractUnit"`
	Tasks        []Task     `gorm:"foreignKey:ContractUID;references:UID" json:"tasks"`
}

//回款记录Model
type Collection struct {
	gorm.Model
	UID         string `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	ContractUID string `gorm:"type:varchar(32);comment:合同ID" json:"contractUID"`
	Amount      int    `gorm:"type:int;comment:金额(元)" json:"totalAmount"`
	Remarks     string `gorm:"type:varchar(200);comment:回款详情记录" json:"remarks"`
}

type ContractQuery struct {
	AreaUID     string `json:"areaUID"`
	No          string `json:"no"`
	CompanyName string `json:"companyName"`
}

func InsertContract(contract *Contract) (code int) {
	contract.UID = uid.Generate()
	if contract.IsEntryCustomer {
		contract.Customer = Customer{}
	}
	err = db.Create(&contract).Error
	if err != nil {
		return msg.ERROR_CONTRACT_INSERT
	}
	return msg.SUCCESS
}

func DeleteContract(uid string) (code int) {
	err = db.Delete(&Contract{}, "uid = ?", uid).Error
	if err != nil {
		return msg.ERROR_CONTRACT_DELETE
	}
	return msg.SUCCESS
}

func UpdateContract(contract *Contract) (code int) {
	var maps = make(map[string]interface{})
	maps["Remarks"] = contract.Remarks
	err = db.Model(&Contract{}).Where("uid = ?", contract.UID).Updates(maps).Error
	if err != nil {
		return msg.ERROR_CONTRACT_UPDATE
	}
	return msg.SUCCESS
}

func SelectContract(uid string) (contract Contract, code int) {
	err = db.Preload("Area").Preload("ContractUnit").
		Preload("Employee").Preload("Employee.Office").
		Preload("Customer").Preload("Customer.Company").
		Preload("Tasks").Preload("Tasks.Product").
		First(&contract, "uid = ?", uid).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return contract, msg.ERROR_CONTRACT_NOT_EXIST
		} else {
			return contract, msg.ERROR_CONTRACT_SELECT
		}
	}
	return contract, msg.SUCCESS
}

func SelectContracts(pageSize int, pageNo int, contractQuery *ContractQuery) (contracts []Contract, code int, total int64) {
	var maps = make(map[string]interface{})
	if contractQuery.AreaUID != "" {
		maps["area_uid"] = contractQuery.AreaUID
	}

	err = db.Where("no LIKE ?", "%"+contractQuery.No+"%").Where(maps).
		Preload("Customer").Find(&contracts).Count(&total).
		Preload("Area").Preload("Employee").
		Limit(pageSize).Offset((pageNo - 1) * pageSize).
		Find(&contracts).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return contracts, msg.ERROR, total
	}
	return contracts, msg.SUCCESS, total
}
