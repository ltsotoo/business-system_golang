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
	code = model.InsertProduct(&product)
	msg.Message(c, code, product)
}

//删除产品
func DelProduct(c *gin.Context) {
	uid := c.Param("uid")
	code = model.DeleteProduct(uid)
	msg.Message(c, code, nil)
}

//编辑产品
func EditProductBase(c *gin.Context) {
	var product model.Product
	_ = c.ShouldBindJSON(&product)
	code = model.UpdateProductBase(&product)
	msg.Message(c, code, product)
}

func EditProductNumber(c *gin.Context) {
	var product model.Product
	_ = c.ShouldBindJSON(&product)
	code = model.UpdateProductNumber(&product)
	msg.Message(c, code, nil)
}

func EditProductPrice(c *gin.Context) {
	var product model.Product
	_ = c.ShouldBindJSON(&product)
	code = model.UpdateProductPrice(&product)
	msg.Message(c, code, nil)
}

//查询产品
func QueryProduct(c *gin.Context) {
	var product model.Product
	uid := c.Param("uid")
	product, code = model.SelectProduct(uid)
	msg.Message(c, code, product)
}

//查询产品列表
func QueryProducts(c *gin.Context) {
	var products []model.Product
	var total int64
	var productQuery model.ProductQuery

	_ = c.ShouldBindJSON(&productQuery)

	pageSize, pageSizeErr := strconv.Atoi(c.DefaultQuery("pageSize", "0"))
	pageNo, pageNoErr := strconv.Atoi(c.DefaultQuery("pageNo", "0"))
	if pageSizeErr != nil || pageSize <= 0 {
		pageSize = 10
	}
	if pageNoErr != nil || pageNo <= 0 {
		pageNo = 1
	}

	products, code, total = model.SelectProducts(pageSize, pageNo, &productQuery)
	msg.MessageForList(c, msg.SUCCESS, products, pageSize, pageNo, total)
}

func AddProductType(c *gin.Context) {
	var productType model.ProductType
	_ = c.ShouldBindJSON(&productType)
	code = model.InsertProductType(&productType)
	msg.Message(c, code, productType)
}

func DelProductType(c *gin.Context) {
	uid := c.Param("uid")
	code = model.DeleteProductType(uid)
	msg.Message(c, code, nil)
}

func EditProductType(c *gin.Context) {
	var productType model.ProductType
	_ = c.ShouldBindJSON(&productType)
	code = model.UpdateProductType(&productType)
	msg.Message(c, code, productType)
}

func QueryProductTypes(c *gin.Context) {
	var productTypes []model.ProductType
	var productType model.ProductType

	_ = c.ShouldBindJSON(&productType)

	productTypes, code = model.SelectProductTypes(&productType)
	msg.Message(c, code, productTypes)
}
