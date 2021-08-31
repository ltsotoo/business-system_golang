package v1

import (
	"business-system_golang/model"
	"business-system_golang/utils/msg"
	"business-system_golang/utils/rbac"
	"strconv"

	"github.com/gin-gonic/gin"
)

//录入供应商
func EntrySupplier(c *gin.Context) {
	code = rbac.Check(c, "supplier", "insert")
	if code == msg.ERROR {
		msg.MessageForNotPermission(c)
	} else {
		var supplier model.Supplier
		_ = c.ShouldBindJSON(&supplier)
		code = model.InsertSupplier(&supplier)
		msg.Message(c, code, supplier)
	}
}

//删除供应商
func DelSupplier(c *gin.Context) {
	code = rbac.Check(c, "supplier", "delete")
	if code == msg.ERROR {
		msg.MessageForNotPermission(c)
	} else {
		uid := c.Param("uid")
		code = model.DeleteSupplier(uid)
		msg.Message(c, code, nil)
	}
}

//编辑供应商
func EditSupplier(c *gin.Context) {
	code = rbac.Check(c, "supplier", "update")
	if code == msg.ERROR {
		msg.MessageForNotPermission(c)
	} else {
		var supplier model.Supplier
		_ = c.ShouldBindJSON(&supplier)
		code = model.UpdateSupplier(&supplier)
		msg.Message(c, code, supplier)
	}
}

//查询供应商
func QuerySupplier(c *gin.Context) {
	code = rbac.Check(c, "supplier", "select")
	if code == msg.ERROR {
		msg.MessageForNotPermission(c)
	} else {
		var supplier model.Supplier
		uid := c.Param("uid")
		supplier, code = model.SelectSupplier(uid)
		msg.Message(c, code, supplier)
	}
}

//查询供应商列表
func QuerySuppliers(c *gin.Context) {
	var suppliers []model.Supplier
	var total int64
	var supplierQuery model.SupplierQuery

	_ = c.ShouldBindJSON(&supplierQuery)

	pageSize, pageSizeErr := strconv.Atoi(c.Query("pageSize"))
	pageNo, pageNoErr := strconv.Atoi(c.Query("pageNo"))
	if pageSizeErr != nil || pageSize < 0 {
		pageSize = 10
	}
	if pageNoErr != nil || pageNo < 0 {
		pageNo = 1
	}

	suppliers, code, total = model.SelectSuppliers(pageSize, pageNo, &supplierQuery)
	msg.MessageForList(c, msg.SUCCESS, suppliers, pageSize, pageNo, total)
}
