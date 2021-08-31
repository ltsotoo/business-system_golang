package v1

import (
	"business-system_golang/model"
	"business-system_golang/utils/msg"
	"business-system_golang/utils/rbac"
	"strconv"

	"github.com/gin-gonic/gin"
)

//录入产品
func EntryProduct(c *gin.Context) {
	code = rbac.Check(c, "product", "insert")
	if code == msg.SUCCESS {
		var product model.Product
		_ = c.ShouldBindJSON(&product)
		code = model.InsertProduct(&product)
		msg.Message(c, code, product)
	} else {
		msg.MessageForNotPermission(c)
	}
}

//删除产品
func DelProduct(c *gin.Context) {
	code = rbac.Check(c, "product", "delete")
	if code == msg.SUCCESS {
		uid := c.Param("uid")
		code = model.DeleteProduct(uid)
		msg.Message(c, code, nil)
	} else {
		msg.MessageForNotPermission(c)
	}
}

//编辑产品
func EditProduct(c *gin.Context) {
	code = rbac.Check(c, "product", "update")
	if code == msg.ERROR {
		msg.MessageForNotPermission(c)
	} else {
		var product model.Product
		_ = c.ShouldBindJSON(&product)
		code = model.UpdateProduct(&product)
		msg.Message(c, code, product)
	}
}

//查询产品
func QueryProduct(c *gin.Context) {
	code = rbac.Check(c, "product", "select")
	if code == msg.ERROR {
		msg.MessageForNotPermission(c)
	} else {
		var product model.Product
		uid := c.Param("uid")
		product, code = model.SelectProduct(uid)
		msg.Message(c, code, product)
	}
}

//查询产品列表
func QueryProducts(c *gin.Context) {
	var products []model.Product
	var total int64
	var productQuery model.ProductQuery

	_ = c.ShouldBindJSON(&productQuery)

	pageSize, pageSizeErr := strconv.Atoi(c.Query("pageSize"))
	pageNo, pageNoErr := strconv.Atoi(c.Query("pageNo"))

	if pageSizeErr != nil || pageSize <= 0 {
		pageSize = 10
	}
	if pageNoErr != nil || pageNo <= 0 {
		pageNo = 1
	}

	products, code, total = model.SelectProducts(pageSize, pageNo, &productQuery)
	msg.MessageForList(c, msg.SUCCESS, products, pageSize, pageNo, total)
}
