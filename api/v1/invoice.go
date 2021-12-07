package v1

import (
	"business-system_golang/model"
	"business-system_golang/utils/msg"
	"strconv"

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
	var total int64

	_ = c.ShouldBindJSON(&invoice)

	pageSize, pageSizeErr := strconv.Atoi(c.Query("pageSize"))
	pageNo, pageNoErr := strconv.Atoi(c.Query("pageNo"))
	if pageSizeErr != nil || pageSize < 0 {
		pageSize = 10
	}
	if pageNoErr != nil || pageNo < 0 {
		pageNo = 1
	}

	invoices, code, total = model.SelectInvoices(pageSize, pageNo, &invoice)
	msg.MessageForList(c, code, invoices, pageSize, pageNo, total)
}
