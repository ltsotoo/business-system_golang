package v1

import (
	"business-system_golang/model"
	"business-system_golang/utils/msg"
	"strconv"

	"github.com/gin-gonic/gin"
)

//创建预研究
func CreatePreResearch(c *gin.Context) {
	var preResearch model.PreResearch
	_ = c.ShouldBindJSON(&preResearch)
	preResearch.EmployeeUID = c.MustGet("employeeUID").(string)
	code = model.InsertPreResearch(&preResearch)
	msg.Message(c, code, preResearch)
}

func DelPreResearch(c *gin.Context) {
	uid := c.Param("uid")
	code = model.DeletePreResearch(uid)
	msg.Message(c, code, nil)
}

func QueryPreResearch(c *gin.Context) {
	var preResearch model.PreResearch
	uid := c.Param("uid")
	preResearch, code = model.SelectPreReasearch(uid)
	msg.Message(c, code, preResearch)
}

func QueryPreResearchs(c *gin.Context) {
	var preResearchs []model.PreResearch
	var total int64
	var preResearchQuery model.PreResearchQuery

	_ = c.ShouldBindJSON(&preResearchQuery)

	pageSize, pageSizeErr := strconv.Atoi(c.Query("pageSize"))
	pageNo, pageNoErr := strconv.Atoi(c.Query("pageNo"))
	if pageSizeErr != nil || pageSize < 0 {
		pageSize = 10
	}
	if pageNoErr != nil || pageNo < 0 {
		pageNo = 1
	}

	preResearchs, code, total = model.SelectPreReasearchs(pageSize, pageNo, &preResearchQuery)
	msg.MessageForList(c, code, preResearchs, pageSize, pageNo, total)
}

func QueryPreResearchTask(c *gin.Context) {
	var preResearchTask model.PreResearchTask
	uid := c.Param("uid")
	preResearchTask, code = model.SelectPreReasearchTask(uid)
	msg.Message(c, code, preResearchTask)
}

func QueryPreResearchTasks(c *gin.Context) {
	var preResearchTasks []model.PreResearchTask
	var total int64
	var preResearchTask model.PreResearchTask

	_ = c.ShouldBindJSON(&preResearchTask)

	pageSize, pageSizeErr := strconv.Atoi(c.Query("pageSize"))
	pageNo, pageNoErr := strconv.Atoi(c.Query("pageNo"))
	if pageSizeErr != nil || pageSize < 0 {
		pageSize = 10
	}
	if pageNoErr != nil || pageNo < 0 {
		pageNo = 1
	}

	preResearchTasks, code, total = model.SelectPreReasearchTasks(pageSize, pageNo, &preResearchTask)
	msg.MessageForList(c, code, preResearchTasks, pageSize, pageNo, total)
}

func ApprovePreResearch(c *gin.Context) {
	var preResearchQuery model.PreResearchQuery
	_ = c.ShouldBindJSON(&preResearchQuery)
	preResearchQuery.AuditorUID = c.MustGet("employeeUID").(string)
	code = model.UpdatePreResearchStatus(&preResearchQuery)
	msg.Message(c, code, nil)
}

func ApprovePreResearchTask(c *gin.Context) {
	var preResearchTask model.PreResearchTask
	_ = c.ShouldBindJSON(&preResearchTask)
	preResearchTask.AuditorUID = c.MustGet("employeeUID").(string)
	code = model.UpdatePreResearchTaskStatus(&preResearchTask)
	msg.Message(c, code, nil)
}
