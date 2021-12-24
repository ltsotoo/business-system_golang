package v1

import (
	"business-system_golang/model"
	"business-system_golang/utils/msg"

	"github.com/gin-gonic/gin"
)

func AddInvoice(c *gin.Context) {
	var invoice model.Invoice
	_ = c.ShouldBindJSON(&invoice)

	invoice.EmployeeUID = c.MustGet("employeeUID").(string)
	invoice.Status = 1

	code = model.InsertInvoice(&invoice)
	msg.Message(c, code, invoice)
}

func DelInvoice(c *gin.Context) {
	uid := c.Param("uid")
	code = model.DeleteInvoice(uid)
	msg.Message(c, code, nil)
}

func EditInvoice(c *gin.Context) {
	var invoice model.Invoice
	_ = c.ShouldBindJSON(&invoice)

	code = model.UpdateInvoice(&invoice)
	msg.Message(c, code, invoice)
}

func ApproveInvoice(c *gin.Context) {
	uid := c.Param("uid")
	code = model.ApproveInvoice(uid)
	msg.Message(c, code, nil)
}

func QueryInvoices(c *gin.Context) {
	var invoices []model.Invoice
	var invoice model.Invoice

	_ = c.ShouldBindJSON(&invoice)

	invoices, code = model.SelectInvoices(&invoice)
	msg.Message(c, code, invoices)
}

func QueryInvoicesAndPayments(c *gin.Context) {
	var contract model.Contract
	_ = c.ShouldBindJSON(&contract)
	contract.Payments, code = model.SelectPayments(contract.UID)
	contract.Invoices, code = model.SelectInvoicesAndPayments(contract.UID)
	msg.Message(c, code, contract)
}
