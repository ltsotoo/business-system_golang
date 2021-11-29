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
	code = model.ApproveBidBond(uid)
	msg.Message(c, code, nil)
}

func QueryBidBonds(c *gin.Context) {
	var bidBonds []model.BidBond
	var bidBond model.BidBond
	var total int64

	_ = c.ShouldBindJSON(&bidBond)

	pageSize, pageSizeErr := strconv.Atoi(c.Query("pageSize"))
	pageNo, pageNoErr := strconv.Atoi(c.Query("pageNo"))
	if pageSizeErr != nil || pageSize < 0 {
		pageSize = 10
	}
	if pageNoErr != nil || pageNo < 0 {
		pageNo = 1
	}

	bidBonds, code, total = model.SelectBidBonds(pageSize, pageNo, &bidBond)
	msg.MessageForList(c, code, bidBonds, pageSize, pageNo, total)
}
