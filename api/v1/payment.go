package v1

import (
	"business-system_golang/model"
	"business-system_golang/utils/msg"

	"github.com/gin-gonic/gin"
)

func AddPayment(c *gin.Context) {
	if model.SystemSettlement {
		msg.MessageForSystemSettlement(c)
		return
	}

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

func ChangeContractCollectionStatus(c *gin.Context) {
	var contract model.Contract
	_ = c.ShouldBindJSON(&contract)

	code = model.UpdateContractCollectionStatus(&contract)

	msg.Message(c, code, nil)
}

func QueryPayments(c *gin.Context) {
	var payments []model.Payment
	contractUID := c.Param("contractUID")

	payments, code = model.SelectPayments(contractUID)
	msg.Message(c, code, payments)
}

func QueryPrePayments(c *gin.Context) {
	var payments []model.Payment
	contractUID := c.Param("contractUID")

	payments, code = model.SelectPrePayments(contractUID)
	msg.Message(c, code, payments)
}
