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

func FinishPayments(c *gin.Context) {
	var contract model.Contract
	_ = c.ShouldBindJSON(&contract)

	code = model.UpdateContractCollectionStatusToFinish(&contract)

	msg.Message(c, code, nil)
}

func RejectPayments(c *gin.Context) {
	var contract model.Contract
	_ = c.ShouldBindJSON(&contract)

	code = model.UpdateContractCollectionStatusToNotFinish(&contract)

	msg.Message(c, code, nil)
}
