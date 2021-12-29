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

func QueryInvoices(c *gin.Context) {
	var invoices []model.Invoice
	contractUID := c.Param("contractUID")

	invoices, code = model.SelectInvoices(contractUID)
	msg.Message(c, code, invoices)
}
