package v1

import (
	"business-system_golang/model"
	"business-system_golang/utils/msg"
	"strconv"

	"github.com/gin-gonic/gin"
)

//录入产品
func EntryProduct(c *gin.Context) {
	var product model.Product
	_ = c.ShouldBindJSON(&product)
	code = model.CreateProduct(&product)
	msg.Message(c, code, product)
}

//删除产品
func DelProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteProduct(id)
	msg.Message(c, code, nil)
}

//编辑产品
func EditProduct(c *gin.Context) {
	var product model.Product
	_ = c.ShouldBindJSON(&product)
	code = model.UpdateProduct(&product)
	msg.Message(c, code, product)
}

//查询产品
func QueryProduct(c *gin.Context) {
	var product model.Product
	id, _ := strconv.Atoi(c.Param("id"))
	product, code = model.SelectProduct(id)
	msg.Message(c, code, product)
}

//查询产品列表
func QueryProducts(c *gin.Context) {
	var products []model.Product
	var total int64

	pageSize, pageSizeErr := strconv.Atoi(c.Query("pageSize"))
	pageNo, pageNoErr := strconv.Atoi(c.Query("pageNo"))
	if pageSizeErr != nil || pageSize <= 0 {
		pageSize = 10
	}
	if pageNoErr != nil || pageNo <= 0 {
		pageNo = 1
	}

	products, code, total = model.SelectProducts(pageSize, pageNo)
	msg.MessageForList(c, msg.SUCCESS, products, pageSize, pageNo, total)
}
