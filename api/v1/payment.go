package v1

import (
	"business-system_golang/model"
	"business-system_golang/utils/msg"

	"github.com/gin-gonic/gin"
)

func DelPayment(c *gin.Context) {
	uid := c.Param("uid")
	code = model.DeletePayment(uid)
	msg.Message(c, code, nil)
}

func AddPayment(c *gin.Context) {
	var payment model.Payment
	_ = c.ShouldBindJSON(&payment)
	payment.EmployeeUID = c.MustGet("employeeUID").(string)
	code = model.InsertPayment(&payment)
	if code == msg.SUCCESS {
		code = checkPaymentsUpdateContract(payment.ContractUID)
	}
	msg.Message(c, code, payment)
}

func EditPayment(c *gin.Context) {
	var payment model.Payment
	_ = c.ShouldBindJSON(&payment)
	code = model.UpdatePayment(&payment)
	msg.Message(c, code, nil)
}

func QueryPaymentsForContract(c *gin.Context) {
	var payments []model.Payment
	contractUID := c.Param("contractUID")
	payments, code = model.SelectPaymentsByContractUID(contractUID)
	msg.Message(c, code, payments)
}

func checkPaymentsUpdateContract(contractUID string) int {
	var payments []model.Payment
	var contract model.Contract
	var amount float64
	payments, _ = model.SelectPaymentsByContractUID(contractUID)
	contract, _ = model.SelectContract(contractUID)
	for _, payment := range payments {
		amount += payment.Money
	}
	if contract.TotalAmount <= amount && amount != 0 {
		code = model.UpdateContractCollectionStatusToFinish(contractUID)
	}
	return code
}
