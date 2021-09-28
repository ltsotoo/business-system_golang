package v1

import (
	"business-system_golang/model"
	"business-system_golang/utils/msg"

	"github.com/gin-gonic/gin"
)

func QueryMyTasks(c *gin.Context) {
	var tasks []model.Task
	tasks, code = model.SelectMyTasks(c.MustGet("employeeUID").(string))
	msg.Message(c, code, tasks)
}

func QueryMyExpense(c *gin.Context) {

}
