package v1

import (
	"business-system_golang/model"
	"business-system_golang/utils/msg"
	"strconv"

	"github.com/gin-gonic/gin"
)

//录入合同
func EntryContract(c *gin.Context) {
	var contract model.Contract
	_ = c.ShouldBindJSON(&contract)
	code = model.CreateContract(&contract)
	msg.Message(c, code, contract)
}

//删除合同
func DelContract(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteContract(id)
	msg.Message(c, code, nil)
}

//编辑合同
func EditContract(c *gin.Context) {
	var contract model.Contract
	_ = c.ShouldBindJSON(&contract)
	code = model.UpdateContract(&contract)
	msg.Message(c, code, contract)
}

//查询合同
func QueryContract(c *gin.Context) {
	var contract model.Contract
	id, _ := strconv.Atoi(c.Param("id"))
	contract, code = model.SelectContract(id)
	msg.Message(c, code, contract)
}

//查询合同列表
func QueryContracts(c *gin.Context) {
	var contracts []model.Contract
	var total int64

	pageSize, pageSizeErr := strconv.Atoi(c.Query("pageSize"))
	pageNo, pageNoErr := strconv.Atoi(c.Query("pageNo"))
	if pageSizeErr != nil || pageSize <= 0 {
		pageSize = 10
	}
	if pageNoErr != nil || pageNo <= 0 {
		pageNo = 1
	}

	contracts, code, total = model.SelectContracts(pageSize, pageNo)
	msg.MessageForList(c, msg.SUCCESS, contracts, pageSize, pageNo, total)
}
