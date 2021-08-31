package v1

import (
	"business-system_golang/model"
	"business-system_golang/utils/msg"
	"business-system_golang/utils/rbac"
	"business-system_golang/utils/uid"
	"strconv"

	"github.com/gin-gonic/gin"
)

//录入合同
func EntryContract(c *gin.Context) {
	code = rbac.Check(c, "contract", "insert")
	if code == msg.ERROR {
		msg.MessageForNotPermission(c)
	} else {
		var contract model.Contract
		_ = c.ShouldBindJSON(&contract)
		for i := range contract.Tasks {
			contract.Tasks[i].UID = uid.Generate()
			contract.TotalAmount += contract.Tasks[i].TotalPrice
		}
		code = model.InsertContract(&contract)
		msg.Message(c, code, contract)
	}
}

//删除合同
func DelContract(c *gin.Context) {
	code = rbac.Check(c, "contract", "delete")
	if code == msg.ERROR {
		msg.MessageForNotPermission(c)
	} else {
		uid := c.Param("uid")
		code = model.DeleteContract(uid)
		msg.Message(c, code, nil)
	}
}

//编辑合同
func EditContract(c *gin.Context) {
	code = rbac.Check(c, "contract", "update")
	if code == msg.ERROR {
		msg.MessageForNotPermission(c)
	} else {
		var contract model.Contract
		_ = c.ShouldBindJSON(&contract)
		code = model.UpdateContract(&contract)
		msg.Message(c, code, contract)
	}
}

//查询合同
func QueryContract(c *gin.Context) {
	code = rbac.Check(c, "contract", "select")
	if code == msg.ERROR {
		msg.MessageForNotPermission(c)
	} else {
		var contract model.Contract
		uid := c.Param("uid")
		contract, code = model.SelectContract(uid)
		msg.Message(c, code, contract)
	}
}

//查询合同列表
func QueryContracts(c *gin.Context) {
	var contracts []model.Contract
	var total int64
	var contractQuery model.ContractQuery

	_ = c.ShouldBindJSON(&contractQuery)

	pageSize, pageSizeErr := strconv.Atoi(c.Query("pageSize"))
	pageNo, pageNoErr := strconv.Atoi(c.Query("pageNo"))
	if pageSizeErr != nil || pageSize < 0 {
		pageSize = 10
	}
	if pageNoErr != nil || pageNo < 0 {
		pageNo = 1
	}

	contracts, code, total = model.SelectContracts(pageSize, pageNo, &contractQuery)
	msg.MessageForList(c, msg.SUCCESS, contracts, pageSize, pageNo, total)
}
