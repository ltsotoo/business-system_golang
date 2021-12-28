package v1

import (
	"business-system_golang/model"
	"business-system_golang/utils/msg"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddBidBond(c *gin.Context) {
	var bidBond model.BidBond
	_ = c.ShouldBindJSON(&bidBond)

	bidBond.EmployeeUID = c.MustGet("employeeUID").(string)
	bidBond.Status = 1

	code = model.InsertBidBond(&bidBond)
	msg.Message(c, code, bidBond)
}

func DelBidBond(c *gin.Context) {
	uid := c.Param("uid")
	code = model.DeleteBidBond(uid)
	msg.Message(c, code, nil)
}

func EditBidBond(c *gin.Context) {
	var bidBond model.BidBond
	_ = c.ShouldBindJSON(&bidBond)

	code = model.UpdateBidBond(&bidBond)
	msg.Message(c, code, bidBond)
}

func ApproveBidBond(c *gin.Context) {
	uid := c.Param("uid")
	code = model.ApproveBidBond(uid, c.MustGet("employeeUID").(string))
	msg.Message(c, code, nil)
}

func QueryBidBonds(c *gin.Context) {
	var bidBonds []model.BidBond
	var bidBondQuery model.BidBondQuery
	var total int64

	_ = c.ShouldBindJSON(&bidBondQuery)

	pageSize, pageSizeErr := strconv.Atoi(c.DefaultQuery("pageSize", "0"))
	pageNo, pageNoErr := strconv.Atoi(c.DefaultQuery("pageNo", "0"))
	if pageSizeErr != nil || pageSize < 0 {
		pageSize = 10
	}
	if pageNoErr != nil || pageNo < 0 {
		pageNo = 1
	}

	bidBonds, code, total = model.SelectBidBonds(pageSize, pageNo, &bidBondQuery)
	msg.MessageForList(c, code, bidBonds, pageSize, pageNo, total)
}
