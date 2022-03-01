package v1

import (
	"business-system_golang/model"
	"business-system_golang/utils/msg"
	"business-system_golang/utils/uid"
	"strconv"

	"github.com/gin-gonic/gin"
)

//暂存合同
func SaveContract(c *gin.Context) {
	var contract model.Contract
	_ = c.ShouldBindJSON(&contract)

	contract.TotalAmount = 0
	contract.Status = 0
	contract.Tasks = nil
	employee, _ := model.SelectEmployee(c.MustGet("employeeUID").(string))
	contract.EmployeeUID = employee.UID
	contract.OfficeUID = employee.OfficeUID
	contract.Customer = model.Customer{}

	if contract.IsOld && !contract.IsEntryCustomer {
		code = msg.ERROR
	} else {
		code = model.SaveContract(&contract)
	}
	msg.Message(c, code, contract)
}

//录入合同
func EntryContract(c *gin.Context) {

	if model.SystemSettlement {
		msg.MessageForSystemSettlement(c)
		return
	}

	var contract model.Contract
	_ = c.ShouldBindJSON(&contract)

	contract.TotalAmount = 0
	contract.Status = 1
	employee, _ := model.SelectEmployee(c.MustGet("employeeUID").(string))
	contract.EmployeeUID = employee.UID
	contract.OfficeUID = employee.OfficeUID

	if contract.IsPreDeposit {
		contract.Tasks = nil
	} else {
		for i := range contract.Tasks {
			contract.Tasks[i].UID = uid.Generate()
			contract.TotalAmount += contract.Tasks[i].TotalPrice
		}
	}

	if contract.IsEntryCustomer {
		contract.Customer = model.Customer{}
	} else {
		contract.CustomerUID = ""
		contract.Customer.UID = uid.Generate()
		contract.Customer.Status = 0
	}

	if contract.IsOld && !contract.IsEntryCustomer {
		code = msg.ERROR
	} else {
		code = model.InsertContract(&contract)
	}
	msg.Message(c, code, contract)
}

//删除合同
func DelContract(c *gin.Context) {
	uid := c.Param("uid")
	code = model.DeleteContract(uid)
	msg.Message(c, code, nil)
}

func QuerySimpleContract(c *gin.Context) {
	var contract model.Contract
	uid := c.Param("uid")
	contract, code = model.SelectSimpleContract(uid)
	msg.Message(c, code, contract)
}

//查询合同
func QueryContract(c *gin.Context) {
	var contract model.Contract
	uid := c.Param("uid")
	contract, code = model.SelectContract(uid)
	msg.Message(c, code, contract)
}

//查询合同列表
func QueryContracts(c *gin.Context) {
	var contracts []model.Contract
	var total int64
	var contractQuery model.ContractQuery

	_ = c.ShouldBindJSON(&contractQuery)

	pageSize, pageSizeErr := strconv.Atoi(c.DefaultQuery("pageSize", "0"))
	pageNo, pageNoErr := strconv.Atoi(c.DefaultQuery("pageNo", "0"))
	if pageSizeErr != nil || pageSize < 0 {
		pageSize = 10
	}
	if pageNoErr != nil || pageNo < 0 {
		pageNo = 1
	}

	contracts, code, total = model.SelectContracts(pageSize, pageNo, &contractQuery)
	msg.MessageForList(c, code, contracts, pageSize, pageNo, total)
}

func ApproveContract(c *gin.Context) {

	if model.SystemSettlement {
		msg.MessageForSystemSettlement(c)
		return
	}

	var contractFlowQuery model.ContractFlowQuery
	_ = c.ShouldBindJSON(&contractFlowQuery)
	code = model.ApproveContract(contractFlowQuery.UID, contractFlowQuery.Status, c.MustGet("employeeUID").(string))
	msg.Message(c, code, nil)
}

func RejectContract(c *gin.Context) {

	if model.SystemSettlement {
		msg.MessageForSystemSettlement(c)
		return
	}

	var contract model.Contract
	_ = c.ShouldBindJSON(&contract)
	if contract.UID != "" {
		code = model.Reject(contract.UID, c.MustGet("employeeUID").(string))
	}
	msg.Message(c, code, contract.UID)
}

//预存款合同完成接口
func ApproveContractProductionStatusToFinish(c *gin.Context) {

	if model.SystemSettlement {
		msg.MessageForSystemSettlement(c)
		return
	}

	var contract model.Contract
	_ = c.ShouldBindJSON(&contract)

	contract, code = model.SelectContract(contract.UID)

	if contract.ID != 0 {
		code = model.UpdateContractProductionStatusToFinish(contract.UID)
	}

	msg.Message(c, code, nil)
}

func EditContract(c *gin.Context) {

	if model.SystemSettlement {
		msg.MessageForSystemSettlement(c)
		return
	}

	var contract model.Contract
	_ = c.ShouldBindJSON(&contract)

	code = model.UpdatePre(&contract)
	msg.Message(c, code, nil)
}
