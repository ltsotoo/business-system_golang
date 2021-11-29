package v1

import (
	"business-system_golang/model"
	"business-system_golang/utils/msg"
	"business-system_golang/utils/pushMoney"
	"business-system_golang/utils/uid"
	"strconv"

	"github.com/gin-gonic/gin"
)

//录入合同
func EntryContract(c *gin.Context) {
	var contract model.Contract
	_ = c.ShouldBindJSON(&contract)
	contract.TotalAmount = 0
	for i := range contract.Tasks {
		contract.Tasks[i].UID = uid.Generate()
		contract.TotalAmount += contract.Tasks[i].TotalPrice
	}
	employee, _ := model.SelectEmployee(c.MustGet("employeeUID").(string))
	contract.EmployeeUID = employee.UID
	contract.OfficeUID = employee.OfficeUID

	code = model.InsertContract(&contract)
	msg.Message(c, code, contract)
}

//删除合同
func DelContract(c *gin.Context) {
	uid := c.Param("uid")
	code = model.DeleteContract(uid)
	msg.Message(c, code, nil)
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

	pageSize, pageSizeErr := strconv.Atoi(c.Query("pageSize"))
	pageNo, pageNoErr := strconv.Atoi(c.Query("pageNo"))
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
	var contractFlowQuery model.ContractFlowQuery
	_ = c.ShouldBindJSON(&contractFlowQuery)
	code = model.ApproveContract(contractFlowQuery.UID, contractFlowQuery.Status, c.MustGet("employeeUID").(string))
	msg.Message(c, code, nil)
}

func FinalApproveContract(c *gin.Context) {
	var contractPushMoney model.ContractPushMoney
	var contract model.Contract
	_ = c.ShouldBindJSON(&contractPushMoney)
	contractPushMoney.EmployeeUID = c.MustGet("employeeUID").(string)
	if contractPushMoney.Type == 2 {
		contractPushMoney.TotalMoney = contractPushMoney.TaskTotalMoney - contractPushMoney.PaymentMoneys
	}
	contract, code = model.SelectContract(contractPushMoney.ContractUID)
	if contract.UID != "" {
		code = model.FinalApproveContract(&contractPushMoney, &contract)
	}
	msg.Message(c, code, nil)
}

func RejectContract(c *gin.Context) {
	var contract model.Contract
	_ = c.ShouldBindJSON(&contract)
	code = model.Reject(contract.UID)
	msg.Message(c, code, nil)
}

func CalculatePushMoney(c *gin.Context) {
	var contract model.Contract
	var contractPushMoney model.ContractPushMoney
	_ = c.ShouldBindJSON(&contract)
	contract, code = model.SelectContract(contract.UID)
	if contract.UID != "" {
		contractPushMoney = pushMoney.Calculate(&contract)
	}
	msg.Message(c, msg.SUCCESS, contractPushMoney)
}
