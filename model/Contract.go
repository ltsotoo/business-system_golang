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
	No                    string  `gorm:"type:varchar(100);comment:合同编号" json:"no"`
	UID                   string  `gorm:"type:varchar(32);comment:唯一标识;unique" json:"UID"`
	RegionUID             string  `gorm:"type:varchar(32);comment:省份UID;default:(-)" json:"regionUID"`
	OfficeUID             string  `gorm:"type:varchar(32);comment:办事处UID;default:(-)" json:"officeUID"`
	EmployeeUID           string  `gorm:"type:varchar(32);comment:业务员UID;default:(-)" json:"employeeUID"`
	IsEntryCustomer       bool    `gorm:"type:boolean;comment:客户是否录入" json:"isEntryCustomer"`
	CustomerUID           string  `gorm:"type:varchar(32);comment:客户ID;default:(-)" json:"customerUID"`
	ContractDate          XDate   `gorm:"type:date;comment:签订日期" json:"contractDate"`
	ContractUnitUID       string  `gorm:"type:varchar(32);comment:签订单位;default:(-)" json:"contractUnitUID"`
	EstimatedDeliveryDate XDate   `gorm:"type:date;comment:合同交货日期" json:"estimatedDeliveryDate"`
	EndDeliveryDate       XTime   `gorm:"type:datetime;comment:实际交货日期;default:(-)" json:"endDeliveryDate"`
	EndPaymentDate        XTime   `gorm:"type:datetime;comment:最后回款日期;default:(-)" json:"endPaymentDate"`
	InvoiceType           int     `gorm:"type:int;comment:开票类型" json:"invoiceType"`
	InvoiceContent        string  `gorm:"type:varchar(600);comment:开票内容" json:"invoiceContent"`
	PaymentContent        string  `gorm:"type:varchar(600);comment:付款方式" json:"paymentContent"`
	IsSpecial             bool    `gorm:"type:boolean;comment:是否是特殊合同" json:"isSpecial"`
	IsPreDeposit          bool    `gorm:"type:boolean;comment:是否是预存款合同" json:"isPreDeposit"`
	PreDeposit            float64 `gorm:"type:decimal(20,6);comment:可用预存款金额" json:"preDeposit"`
	PreDepositRecord      float64 `gorm:"type:decimal(20,6);comment:合同预存款金额" json:"preDepositRecord"`
	PayType               int     `gorm:"type:int;comment:付款类型(1:人民币 2:美元)" json:"payType"`
	TotalAmount           float64 `gorm:"type:decimal(20,6);comment:总金额" json:"totalAmount"`
	PaymentTotalAmount    float64 `gorm:"type:decimal(20,6);comment:回款总金额(人民币)" json:"paymentTotalAmount"`
	Remarks               string  `gorm:"type:varchar(600);comment:备注" json:"remarks"`
	Status                int     `gorm:"type:int;comment:状态(-1:审批驳回 1:待审批 2:未完成 3:已完成);not null" json:"status"`
	ProductionStatus      int     `gorm:"type:int;comment:生产状态(1:生产中 2:生产完成)" json:"productionStatus"`
	CollectionStatus      int     `gorm:"type:int;comment:回款状态(1:回款中 2:回款完成)" json:"collectionStatus"`
	AuditorUID            string  `gorm:"type:varchar(32);comment:审核员ID;default:(-)" json:"auditorUID"`
	IsDelete              bool    `gorm:"type:boolean;comment:是否删除" json:"isDelete"`

	Region       Dictionary `gorm:"foreignKey:RegionUID;references:UID" json:"region"`
	Office       Office     `gorm:"foreignKey:OfficeUID;references:UID" json:"office"`
	Employee     Employee   `gorm:"foreignKey:EmployeeUID;references:UID" json:"employee"`
	Customer     Customer   `gorm:"foreignKey:CustomerUID;references:UID" json:"customer"`
	ContractUnit Dictionary `gorm:"foreignKey:ContractUnitUID;references:UID" json:"contractUnit"`
	Tasks        []Task     `gorm:"foreignKey:ContractUID;references:UID" json:"tasks"`
	Invoices     []Invoice  `gorm:"foreignKey:ContractUID;references:UID" json:"invoices"`
	Payments     []Payment  `gorm:"foreignKey:ContractUID;references:UID" json:"payments"`

	IsFinalCollectionStatus bool `gorm:"-" json:"isFinalCollectionStatus"`
	IsOld                   bool `gorm:"-" json:"isOld"`
}

type ContractQuery struct {
	RegionUID        string `json:"RegionUID"`
	No               string `json:"no"`
	CompanyName      string `json:"companyName"`
	CustomerName     string `json:"customerName"`
	PayType          int    `json:"payType"`
	IsSpecial        int    `json:"isSpecial"`
	IsPreDeposit     int    `json:"isPreDeposit"`
	Status           int    `json:"status"`
	ProductionStatus int    `json:"productionStatus"`
	CollectionStatus int    `json:"collectionStatus"`
	EmployeeUID      string `json:"employeeUID"`
	StartDate        string `json:"startDate"`
	EndDate          string `json:"endDate"`
	EmployeeName     string `json:"employeeName"`
	InvoiceType      int    `json:"invoiceType"`
	OfficeUID        string `json:"officeUID"`
}

type ContractFlowQuery struct {
	UID    string `json:"UID"`
	Status int    `json:"status"`
}

func SaveContract(contract *Contract) (code int) {
	if contract.UID == "" {
		contract.UID = uid.Generate()
		err = db.Create(&contract).Error
	} else {
		err = db.Updates(&contract).Error
	}
	if err != nil {
		return msg.ERROR_CONTRACT_INSERT
	}
	return msg.SUCCESS
}

func InsertContract(contract *Contract) (code int) {
	contract.UID = uid.Generate()

	if contract.ID == 0 {
		err = db.Create(&contract).Error
	} else {
		err = db.Updates(&contract).Error
	}
	if err != nil {
		return msg.ERROR_CONTRACT_INSERT
	}
	return msg.SUCCESS
}

func CreateNo(contract *Contract) (no string) {
	office, _ := SelectOffice(contract.OfficeUID)
	employee, _ := SelectEmployee(contract.EmployeeUID)
	employeeCount := strconv.Itoa(employee.ContractCount + 1)
	if len(employeeCount) == 1 {
		employeeCount = "00" + employeeCount
	} else if len(employeeCount) == 2 {
		employeeCount = "0" + employeeCount
	}
	no = "Bjscistar" + strings.ReplaceAll(contract.CreatedAt.Format("2006-01-02"), "-", "") + "-" + office.Number + employee.Number + employeeCount
	return
}

func DeleteContract(uid string) (code int) {
	err = db.Transaction(func(tdb *gorm.DB) error {
		if tErr := tdb.Model(&Contract{}).Where("uid = ? AND status = ?", uid, -1).Update("is_delete", true).Error; tErr != nil {
			return tErr
		}
		if tErr := tdb.Delete(&Task{}, "contract_uid = ?", uid).Error; tErr != nil {
			return tErr
		}
		return nil
	})
	if err != nil {
		return msg.ERROR_CONTRACT_DELETE
	}
	return msg.SUCCESS
}

func SelectSimpleContract(uid string) (contract Contract, code int) {
	err = db.Preload("Customer.Company").First(&contract, "uid = ?", uid).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return contract, msg.ERROR_CONTRACT_NOT_EXIST
		} else {
			return contract, msg.ERROR_CONTRACT_SELECT
		}
	}
	return contract, msg.SUCCESS
}

func SelectContract(uid string) (contract Contract, code int) {
	err = db.Preload("Region").Preload("Office").Preload("Employee").
		Preload("Customer.Company").Preload("ContractUnit").
		Preload("Tasks.Product.Type").
		Preload("Tasks.TechnicianMan").Preload("Tasks.PurchaseMan").
		Preload("Tasks.InventoryMan").Preload("Tasks.ShipmentMan").
		Where("is_delete = ?", false).
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
	if contractQuery.RegionUID != "" {
		maps["contract.region_uid"] = contractQuery.RegionUID
	}
	if contractQuery.EmployeeUID != "" {
		maps["contract.employee_uid"] = contractQuery.EmployeeUID
	}
	if contractQuery.OfficeUID != "" {
		maps["contract.office_uid"] = contractQuery.OfficeUID
	}
	if contractQuery.PayType != 0 {
		maps["contract.pay_type"] = contractQuery.PayType
	}
	if contractQuery.InvoiceType != 0 {
		maps["contract.invoice_type"] = contractQuery.InvoiceType
	}
	if contractQuery.IsSpecial == 1 {
		maps["contract.is_special"] = true
	} else if contractQuery.IsSpecial == 2 {
		maps["contract.is_special"] = false
	}
	if contractQuery.IsPreDeposit == 1 {
		maps["contract.is_pre_deposit"] = true
	} else if contractQuery.IsPreDeposit == 2 {
		maps["contract.is_pre_deposit"] = false
	}
	if contractQuery.Status != 0 {
		maps["contract.status"] = contractQuery.Status
	}
	if contractQuery.ProductionStatus != 0 {
		maps["contract.production_status"] = contractQuery.ProductionStatus
	}
	if contractQuery.CollectionStatus != 0 {
		maps["contract.collection_status"] = contractQuery.CollectionStatus
	}

	tDb := db.Where(maps).Where("contract.is_delete = ?", false)

	if contractQuery.EmployeeUID == "" {
		tDb = tDb.Where("contract.status > ?", 0)
	}

	if contractQuery.StartDate != "" && contractQuery.EndDate != "" {
		tDb = tDb.Where("contract.contract_date BETWEEN ? AND ?", contractQuery.StartDate, contractQuery.EndDate)
	} else {
		if contractQuery.StartDate != "" {
			tDb = tDb.Where("contract.contract_date >= ?", contractQuery.StartDate)
		}
		if contractQuery.EndDate != "" {
			tDb = tDb.Where("contract.contract_date <= ?", contractQuery.EndDate)
		}
	}

	if contractQuery.CompanyName != "" {
		tDb = tDb.Joins("Customer").
			Joins("left join customer_company on Customer.company_uid = customer_company.uid").
			Where("Customer.name LIKE ? AND customer_company.name LIKE ?", "%"+contractQuery.CustomerName+"%", "%"+contractQuery.CompanyName+"%")
	} else {
		if contractQuery.CustomerName != "" {
			tDb = tDb.Joins("Customer").Where("Customer.name LIKE ?", "%"+contractQuery.CustomerName+"%")
		}
	}

	if contractQuery.EmployeeName != "" {
		tDb = tDb.Joins("Employee").Where("Employee.name LIKE ?", "%"+contractQuery.EmployeeName+"%")
	}

	if contractQuery.No != "" {
		tDb = tDb.Where("contract.no LIKE ?", "%"+contractQuery.No+"%")
	}

	err = tDb.Find(&contracts).Count(&total).Preload("Region").Preload("Office").
		Preload("Customer.Company").Preload("Employee").
		Limit(pageSize).Offset((pageNo - 1) * pageSize).
		Order("created_at desc").Find(&contracts).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return contracts, msg.ERROR, total
	}
	return contracts, msg.SUCCESS, total
}

//审批合同
func ApproveContract(uid string, status int, employeeUID string) (code int) {
	var maps = make(map[string]interface{})

	if status == magic.CONTRACT_STATUS_UNFINISHED {
		//通过
		maps["status"] = status
		maps["auditor_uid"] = employeeUID
		maps["production_status"] = magic.CONTATCT_PRODUCTION_STATUS_ING
		maps["collection_status"] = magic.CONTATCT_COLLECTION_STATUS_ING

		err = db.Transaction(func(tdb *gorm.DB) error {

			var contract Contract
			if tErr := tdb.Preload("Tasks.Product.Type").First(&contract, "uid = ?", uid).Error; tErr != nil {
				return tErr
			}

			//修改合同基础属性(编号、状态、生产状态、回款状态)
			if contract.No == "" {
				//业务员累计合同数目+1
				if tErr := tdb.Exec("UPDATE employee SET contract_count = contract_count + 1 WHERE uid = ?", contract.EmployeeUID).Error; tErr != nil {
					return tErr
				}

				maps["No"] = CreateNo(&contract)
			}
			if tErr := tdb.Model(&Contract{}).Where("uid = ?", uid).Updates(maps).Error; tErr != nil {
				return tErr
			}

			for i := range contract.Tasks {
				//产品可售库存减一
				if tErr := tdb.Exec("UPDATE product SET number = number - ? WHERE uid = ?", contract.Tasks[i].Number, contract.Tasks[i].ProductUID).Error; tErr != nil {
					return tErr
				}
			}

			//审批录入时的新客户
			if !contract.IsEntryCustomer {
				if tErr := tdb.Exec("UPDATE customer SET status = 1 WHERE uid = ?", contract.CustomerUID).Error; tErr != nil {
					return tErr
				}
			}
			t := time.Now().Format("2006-01-02 15:04:05")

			//修改合同产品任务的开始时间
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
	} else if status == magic.CONTRACT_STATUS_REJECT {
		//驳回
		maps["status"] = status
		maps["auditor_uid"] = employeeUID
		err = db.Transaction(func(tdb *gorm.DB) error {
			if tErr := tdb.Model(&Contract{}).Where("uid = ?", uid).Updates(maps).Error; tErr != nil {
				return tErr
			}
			if tErr := tdb.Delete(&Task{}, "contract_uid = ?", uid).Error; tErr != nil {
				return tErr
			}
			return nil
		})

	}

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
		code = checkContractFinish(uid)
	}
	return
}

//变更合同回款状态为已完成
func UpdateContractCollectionStatus(contract *Contract) (code int) {
	var maps = make(map[string]interface{})
	if contract.IsFinalCollectionStatus {
		maps["collection_status"] = magic.CONTATCT_COLLECTION_STATUS_FINISH
	} else {
		maps["collection_status"] = magic.CONTATCT_COLLECTION_STATUS_ING
		maps["status"] = magic.CONTRACT_STATUS_UNFINISHED
	}
	err = db.Model(&Contract{}).Where("uid = ?", contract.UID).Updates(maps).Error

	if err != nil {
		code = msg.ERROR_CONTRACT_UPDATE_P_STATUS
	} else {
		if contract.IsFinalCollectionStatus {
			code = checkContractFinish(contract.UID)
		} else {
			code = msg.SUCCESS
		}
	}
	return
}

//检查合同是否完成
func checkContractFinish(uid string) int {
	var contract Contract
	db.First(&contract, "uid = ?", uid)
	if contract.UID != "" {
		if contract.ProductionStatus == magic.CONTATCT_PRODUCTION_STATUS_FINISH &&
			contract.CollectionStatus == magic.CONTATCT_COLLECTION_STATUS_FINISH {
			err = db.Model(&Contract{}).Where("uid = ?", uid).Update("status", magic.CONTARCT_STATUS_FINISH).Error
			if err != nil {
				return msg.ERROR_CONTRACT_UPDATE_P_STATUS
			}
		}
		return msg.SUCCESS
	}
	return msg.ERROR
}

//合同执行中途驳回
func Reject(uid string, employeeUID string) (code int) {
	err = db.Transaction(func(tdb *gorm.DB) error {
		var payments []Payment
		var contract Contract
		if tErr := tdb.First(&contract, "uid = ?", uid).Error; tErr != nil {
			return tErr
		}
		if tErr := tdb.Preload("Task.Product.Type").Where("contract_uid = ?", uid).Find(&payments).Error; tErr != nil {
			return tErr
		}

		//回款记录统计
		var tempTargetLoad, tempMoney, tempMoneyCold, tempBusinessMoney float64
		if contract.IsPreDeposit {
			for k := range payments {
				if payments[k].TaskUID != "" {
					tempMoney += payments[k].PushMoney * 0.5
					tempMoneyCold += payments[k].PushMoney - payments[k].PushMoney*0.5
					tempBusinessMoney += payments[k].BusinessMoney
					if !payments[k].Task.Product.Type.IsTaskLoad {
						tempTargetLoad -= payments[k].Money
					}
				} else {
					tempTargetLoad += payments[k].Money
				}
			}
		} else {
			for k := range payments {
				tempMoney += payments[k].PushMoney * 0.5
				tempMoneyCold += payments[k].PushMoney - payments[k].PushMoney*0.5
				tempBusinessMoney += payments[k].BusinessMoney
				if payments[k].Task.Product.Type.IsTaskLoad {
					tempTargetLoad += payments[k].Money
				}
			}
		}

		//办事处变更记录
		var historyOffice HistoryOffice
		historyOffice.OfficeUID = contract.OfficeUID
		historyOffice.ChangeBusinessMoney = -tempBusinessMoney
		historyOffice.ChangeMoney = -tempMoney
		historyOffice.ChangeMoneyCold = -tempMoneyCold
		historyOffice.ChangeTargetLoad = -tempTargetLoad
		historyOffice.Remarks = "合同(" + contract.No + ")驳回"
		historyOffice.EmployeeUID = employeeUID
		if tErr := InsertHistoryOffice(&historyOffice, tdb); tErr != nil {
			return tErr
		}

		//办事处记录变更
		if tErr := tdb.Exec("UPDATE office SET target_load = target_load - ?,money = money - ?, money_cold = money_cold - ?, business_money = business_money - ? WHERE uid = ?", tempTargetLoad, tempMoney, tempMoneyCold, tempBusinessMoney, contract.OfficeUID).Error; tErr != nil {
			return tErr
		}
		//合同状态变更为驳回
		if tErr := tdb.Model(&Contract{}).Where("uid = ?", uid).Update("status", -1).Error; tErr != nil {
			return tErr
		}
		//删除任务，回款记录，支票
		if tErr := tdb.Delete(&Task{}, "contract_uid = ?", uid).Error; tErr != nil {
			return tErr
		}
		if tErr := tdb.Delete(&Payment{}, "contract_uid = ?", uid).Error; tErr != nil {
			return tErr
		}
		if tErr := tdb.Model(&Invoice{}).Where("contract_uid = ?", uid).Update("is_delete", true).Error; tErr != nil {
			return tErr
		}
		return nil
	})
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}

func UpdatePre(contract *Contract) (code int) {
	var maps = make(map[string]interface{})
	maps["pre_deposit"] = contract.PreDeposit
	maps["pre_deposit_record"] = contract.PreDepositRecord
	maps["payment_total_amount"] = contract.PaymentTotalAmount

	err = db.Model(&Contract{}).Where("uid = ?", contract.UID).Updates(maps).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUCCESS
}
