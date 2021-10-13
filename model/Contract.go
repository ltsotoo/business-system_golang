package model

import (
	"business-system_golang/utils/magic"
	"business-system_golang/utils/msg"
	"business-system_golang/utils/uid"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

// 合同 Model
type Contract struct {
	BaseModel
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
	ProductionStatus      int    `gorm:"type:int;comment:生产状态" json:"productionStatus"`
	CollectionStatus      int    `gorm:"type:int;comment:回款状态" json:"collectionStatus"`
	AuditorUID            string `gorm:"type:varchar(32);comment:审核员ID;default:(-)" json:"auditor"`

	Area         Area       `gorm:"foreignKey:AreaUID;references:UID" json:"area"`
	Employee     Employee   `gorm:"foreignKey:EmployeeUID;references:UID" json:"employee"`
	Customer     Customer   `gorm:"foreignKey:CustomerUID;references:UID" json:"customer"`
	ContractUnit Dictionary `gorm:"foreignKey:ContractUnitUID;references:UID" json:"contractUnit"`
	Tasks        []Task     `gorm:"foreignKey:ContractUID;references:UID" json:"tasks"`
}

//回款记录Model
type Collection struct {
	BaseModel
	UID         string `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	ContractUID string `gorm:"type:varchar(32);comment:合同ID" json:"contractUID"`
	Amount      int    `gorm:"type:int;comment:金额(元)" json:"totalAmount"`
	Remarks     string `gorm:"type:varchar(200);comment:回款详情记录" json:"remarks"`
}

type ContractQuery struct {
	AreaUID          string `json:"areaUID"`
	No               string `json:"no"`
	CompanyName      string `json:"companyName"`
	CustomerName     string `json:"customerName"`
	IsSpecial        int    `json:"isSpecial"`
	Status           int    `json:"status"`
	ProductionStatus int    `json:"productionStatus"`
	CollectionStatus int    `json:"collectionStatus"`
}

type ContractFlowQuery struct {
	UID    string `json:"UID"`
	Status int    `json:"status"`
}

func InsertContract(contract *Contract) (code int) {
	contract.UID = uid.Generate()
	if contract.IsEntryCustomer {
		contract.Customer = Customer{}
	}
	// err = db.Create(&contract).Error
	err = db.Transaction(func(tdb *gorm.DB) error {
		if tErr := tdb.Create(&contract).Error; tErr != nil {
			return tErr
		}
		if tErr := tdb.Model(&Contract{}).Where("uid = ?", contract.UID).Update("no", CreateNo(contract)).Error; tErr != nil {
			return tErr
		}
		return nil
	})
	if err != nil {
		return msg.ERROR_CONTRACT_INSERT
	}
	return msg.SUCCESS
}

func CreateNo(contract *Contract) (no string) {
	area, _ := SelectArea(contract.AreaUID)
	employee, _ := SelectEmployee(contract.EmployeeUID)
	tString := strings.ReplaceAll(contract.ContractDate, "-", "")
	no = "bjscistar-" + tString + "-" + area.Number + employee.Number + "0" + strconv.Itoa(int(contract.ID))
	return
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
	if contractQuery.IsSpecial == 1 {
		maps["is_special"] = true
	} else if contractQuery.IsSpecial == 2 {
		maps["is_special"] = false
	}
	if contractQuery.Status != 0 {
		maps["status"] = contractQuery.Status
	}
	if contractQuery.ProductionStatus != 0 {
		maps["production_status"] = contractQuery.ProductionStatus
	}
	if contractQuery.CollectionStatus != 0 {
		maps["collection_status"] = contractQuery.CollectionStatus
	}
	tDb := db

	if contractQuery.CompanyName != "" {
		tDb = tDb.Joins("Customer").
			Joins("left join customer_company on Customer.company_uid = customer_company.uid").
			Where("Customer.name LIKE ? AND customer_company.name LIKE ?", "%"+contractQuery.CustomerName+"%", "%"+contractQuery.CompanyName+"%")
	} else {
		if contractQuery.CustomerName != "" {
			tDb = tDb.Joins("Customer").
				Where("Customer.name LIKE ?", "%"+contractQuery.CustomerName+"%")
		}
	}

	if contractQuery.No != "" {
		tDb = tDb.Where("contract.no LIKE ?", "%"+contractQuery.No+"%")
	}
	if len(maps) > 0 {
		tDb = tDb.Where(maps)
	}
	err = tDb.Find(&contracts).Count(&total).
		Preload("Customer").Preload("Customer.Company").Preload("Area").Preload("Employee").
		Limit(pageSize).Offset((pageNo - 1) * pageSize).
		Find(&contracts).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return contracts, msg.ERROR, total
	}
	return contracts, msg.SUCCESS, total
}

//审批合同
func ApproveContract(uid string, status int, employeeUID string) (code int) {
	var maps = make(map[string]interface{})
	maps["Status"] = status
	maps["AuditorUID"] = employeeUID

	if status == 2 {
		maps["CollectionStatus"] = 1
		maps["ProductionStatus"] = 1
	}

	err = db.Model(&Contract{}).Where("uid = ?", uid).
		Updates(maps).Error

	if err != nil {
		code = msg.ERROR_CONTRACT_UPDATE_STATUS
	} else {
		code = msg.SUCCESS
	}
	return
}

//变更合同生产状态为已完成
func UpdateContractProductionStatusToFinish(uid string) (code int) {
	var maps = make(map[string]interface{})
	maps["production_status"] = magic.CONTATCT_PRODUCTION_STATUS_FINISH
	maps["end_delivery_date"] = time.Now().Format("2006-01-02 15:04:05")
	err = db.Model(&Contract{}).Where("uid = ?", uid).Updates(maps).Error

	if err != nil {
		code = msg.ERROR_CONTRACT_UPDATE_P_STATUS
	} else {
		code = msg.SUCCESS
	}
	return
}

//变更合同回款状态为已完成
func UpdateContractCollectionStatusToFinish(uid string) (code int) {
	err = db.Model(&Contract{}).Where("uid = ?", uid).Update("collection_status", magic.CONTATCT_COLLECTION_STATUS_FINISH).Error

	if err != nil {
		code = msg.ERROR_CONTRACT_UPDATE_P_STATUS
	} else {
		code = msg.SUCCESS
	}
	return
}
