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
	No                    string  `gorm:"type:varchar(32);comment:合同编号" json:"no"`
	UID                   string  `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	AreaUID               string  `gorm:"type:varchar(32);comment:所属区域UID;default:(-)" json:"areaUID"`
	EmployeeUID           string  `gorm:"type:varchar(32);comment:业务员UID;default:(-)" json:"employeeUID"`
	IsEntryCustomer       bool    `gorm:"type:boolean;comment:客户是否录入" json:"isEntryCustomer"`
	CustomerUID           string  `gorm:"type:varchar(32);comment:客户ID;default:(-)" json:"customerUID"`
	ContractDate          string  `gorm:"type:varchar(20);comment:签订日期" json:"contractDate"`
	ContractUnitUID       string  `gorm:"type:varchar(32);comment:签订单位;default:(-)" json:"contractUnitUID"`
	EstimatedDeliveryDate string  `gorm:"type:varchar(20);comment:合同交货日期" json:"estimatedDeliveryDate"`
	EndDeliveryDate       XTime   `gorm:"type:datetime;comment:实际交货日期;default:(-)" json:"endDeliveryDate"`
	EndPaymentDate        string  `gorm:"type:varchar(20);comment:最终回款日期;default:(-)" json:"endPaymentDate"`
	InvoiceType           int     `gorm:"type:int;comment:开票类型" json:"invoiceType"`
	InvoiceContent        string  `gorm:"type:varchar(600);comment:开票内容" json:"invoiceContent"`
	IsSpecial             bool    `gorm:"type:boolean;comment:特殊合同?" json:"isSpecial"`
	PayType               int     `gorm:"type:int;comment:付款类型(1:人民币 2:美元)" json:"payType"`
	TotalAmount           float64 `gorm:"type:decimal(20,6);comment:总金额" json:"totalAmount"`
	PaymentTotalAmount    float64 `gorm:"type:decimal(20,6);comment:回款总金额(人民币)" json:"paymentTotalAmount"`
	Remarks               string  `gorm:"type:varchar(600);comment:备注" json:"remarks"`
	Status                int     `gorm:"type:int;comment:状态(-1:审批驳回 1:待审批 2:未完成 3:已完成);not null" json:"status"`
	ProductionStatus      int     `gorm:"type:int;comment:生产状态(1:生产中 2:生产完成)" json:"productionStatus"`
	CollectionStatus      int     `gorm:"type:int;comment:回款状态(1:回款中 2:回款完成)" json:"collectionStatus"`
	AuditorUID            string  `gorm:"type:varchar(32);comment:审核员ID;default:(-)" json:"auditorUID"`
	FinalAuditorUID       string  `gorm:"type:varchar(32);comment:最终审核员ID;default:(-)" json:"finalAuditorUID"`

	Area         Area       `gorm:"foreignKey:AreaUID;references:UID" json:"area"`
	Employee     Employee   `gorm:"foreignKey:EmployeeUID;references:UID" json:"employee"`
	Customer     Customer   `gorm:"foreignKey:CustomerUID;references:UID" json:"customer"`
	ContractUnit Dictionary `gorm:"foreignKey:ContractUnitUID;references:UID" json:"contractUnit"`
	Tasks        []Task     `gorm:"foreignKey:ContractUID;references:UID" json:"tasks"`
	Payments     []Payment  `gorm:"foreignKey:ContractUID;references:UID" json:"payments"`
}

type ContractPushMoney struct {
	BaseModel
	UID            string  `gorm:"type:varchar(32);comment:唯一标识;not null;unique" json:"UID"`
	ContractUID    string  `gorm:"type:varchar(32);comment:合同ID" json:"contractUID"`
	EmployeeUID    string  `gorm:"type:varchar(32);comment:提交用户UID" json:"employeeUID"`
	Type           int     `gorm:"type:int;comment:状态(1:机器计算 2:手动输入)" json:"type"`
	TaskTotalMoney float64 `gorm:"type:decimal(20,6);comment:任务提成总额" json:"taskTotalMoney"`
	PaymentDays    int     `gorm:"type:int;comment:回款延迟天数" json:"paymentDays"`
	PaymentMoneys  float64 `gorm:"type:decimal(20,6);comment:回款延迟扣除总额" json:"paymentMoneys"`
	TotalMoney     float64 `gorm:"type:decimal(20,6);comment:最终提成总额" json:"totalMoney"`
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
	EmployeeUID      string `json:"employeeUID"`
}

type ContractFlowQuery struct {
	UID    string `json:"UID"`
	Status int    `json:"status"`
}

type ContractPushMoneyQuery struct {
	ContractUID    string  `json:"contractUID"`
	Tasks          []Task  `json:"tasks"`
	TaskTotalMoney float64 `json:"taskTotalMoney"`
	PaymentDays    int     `json:"paymentDays"`
	PaymentMoneys  float64 `json:"paymentMoneys"`
	TotalMoney     float64 `json:"totalMoney"`
}

func InsertContract(contract *Contract) (code int) {
	contract.UID = uid.Generate()
	contract.Status = 1
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
	err = db.Preload("Area.Office").Preload("ContractUnit").Preload("Employee").
		Preload("Customer.Company").Preload("Tasks.Product").
		Preload("Tasks.TechnicianMan").Preload("Tasks.PurchaseMan").
		Preload("Tasks.InventoryMan").Preload("Tasks.ShipmentMan").
		Preload("Payments").
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
	if contractQuery.EmployeeUID != "" {
		maps["employee_uid"] = contractQuery.EmployeeUID
	}
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
		Preload("Customer").Preload("Customer.Company").Preload("Area.Office").Preload("Employee").
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

	if status == 2 {
		maps["Status"] = status
		maps["AuditorUID"] = employeeUID
		maps["CollectionStatus"] = 1
		maps["ProductionStatus"] = 1
	} else if status == -1 {
		maps["Status"] = status
		maps["AuditorUID"] = employeeUID
	}

	if status == 2 {
		err = db.Transaction(func(tdb *gorm.DB) error {
			if tErr := tdb.Model(&Contract{}).Where("uid = ?", uid).Updates(maps).Error; tErr != nil {
				return tErr
			}
			t := time.Now().Format("2006-01-02 15:04:05")
			if tErr := tdb.Model(&Task{}).Where("contract_uid = ? AND type = ?", uid, 1).Update("inventory_start_date", t).Error; tErr != nil {
				return tErr
			}
			if tErr := tdb.Model(&Task{}).Where("contract_uid = ? AND type = ?", uid, 2).Update("purchase_start_date", t).Error; tErr != nil {
				return tErr
			}
			if tErr := tdb.Model(&Task{}).Where("contract_uid = ? AND type = ?", uid, 3).Update("technician_start_date", t).Error; tErr != nil {
				return tErr
			}
			return nil
		})
	} else if status == -1 {
		err = db.Model(&Contract{}).Where("uid = ?", uid).Updates(maps).Error
	}

	if err != nil {
		code = msg.ERROR_CONTRACT_UPDATE_STATUS
	} else {
		code = msg.SUCCESS
	}
	return
}

func FinalApproveContract(contractPushMoney *ContractPushMoney, contract *Contract) (code int) {
	var maps = make(map[string]interface{})

	maps["Status"] = 3
	maps["FinalAuditorUID"] = contractPushMoney.EmployeeUID
	maps["CollectionStatus"] = 2
	maps["ProductionStatus"] = 2

	err = db.Transaction(func(tdb *gorm.DB) error {
		contractPushMoney.UID = uid.Generate()
		if tErr := tdb.Create(&contractPushMoney).Error; tErr != nil {
			return tErr
		}
		if tErr := tdb.Model(&Contract{}).Where("uid = ?", contractPushMoney.ContractUID).Updates(maps).Error; tErr != nil {
			return tErr
		}
		if tErr := tdb.Exec("UPDATE office SET money = money + ? WHERE uid = ?", contractPushMoney.TotalMoney, contract.Area.Office).Error; tErr != nil {
			return tErr
		}
		return nil
	})

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
func UpdateContractCollectionStatusToFinish(contract *Contract) (code int) {
	var maps = make(map[string]interface{})
	maps["collection_status"] = magic.CONTATCT_COLLECTION_STATUS_FINISH
	maps["end_payment_date"] = contract.EndPaymentDate
	err = db.Model(&Contract{}).Where("uid = ?", contract.UID).Updates(maps).Error

	if err != nil {
		code = msg.ERROR_CONTRACT_UPDATE_P_STATUS
	} else {
		code = msg.SUCCESS
	}
	return
}

func UpdateContractCollectionStatusToNotFinish(contract *Contract) (code int) {
	var maps = make(map[string]interface{})
	maps["collection_status"] = magic.CONTATCT_COLLECTION_STATUS_ING
	maps["end_payment_date"] = nil
	err = db.Model(&Contract{}).Where("uid = ?", contract.UID).Updates(maps).Error

	if err != nil {
		code = msg.ERROR_CONTRACT_UPDATE_P_STATUS
	} else {
		code = msg.SUCCESS
	}
	return
}

func Reject(uid string) (code int) {
	err = db.Transaction(func(tdb *gorm.DB) error {
		if tErr := tdb.Model(&Contract{}).Where("uid = ?", uid).Update("status", -1).Error; tErr != nil {
			return tErr
		}
		if tErr := tdb.Delete(&Payment{}, "contract_uid = ?", uid).Error; tErr != nil {
			return tErr
		}
		if tErr := tdb.Delete(&Task{}, "contract_uid = ?", uid).Error; tErr != nil {
			return tErr
		}
		return nil
	})
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}
