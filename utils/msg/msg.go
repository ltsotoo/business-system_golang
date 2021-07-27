package msg

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	SUCCESS = 666
	ERROR   = 999

	//code = 10000-20000 员工模块错误
	ERROR_EMPLOYEE_EXIST      = 10001
	ERROR_EMPLOYEE_NOT_EXIST  = 10002
	ERROR_EMPLOYEE_LOGIN_FAIL = 10003

	ERRPR_TOKEN_NOT_EXIST = 10011
	ERROR_TOKEN_EXPIRED   = 10012
	ERRPR_TOKEN_WRONG     = 10013
	//code = 20000-30000 合同模块错误
	ERROR_CONTRACT_NOT_EXIST = 20001
	//code = 30000-40000 任务模块错误
	//code = 40000-50000 客户模块错误
	ERROR_CUSTOMER_NOT_EXIST = 40001
	//code = 50000-60000 产品模块错误
	ERROR_PRODUCT_NOT_EXIST = 50001
	//code = 60000-70000 供应商模块错误
	ERROR_SUPPLIER_NOT_EXIST = 60001
)

var codeMsg = map[int]string{
	SUCCESS: "成功",
	ERROR:   "失败",

	//员工模块
	ERROR_EMPLOYEE_EXIST:      "员工已录入",
	ERROR_EMPLOYEE_NOT_EXIST:  "员工未录入或已删除",
	ERROR_EMPLOYEE_LOGIN_FAIL: "手机号或者密码错误",

	ERRPR_TOKEN_NOT_EXIST: "凭证不存在",
	ERROR_TOKEN_EXPIRED:   "凭证过期",
	ERRPR_TOKEN_WRONG:     "凭证格式异常",
	//合同模块
	ERROR_CONTRACT_NOT_EXIST: "合同未录入或已删除",
	//任务模块
	//客户模块
	ERROR_CUSTOMER_NOT_EXIST: "客户未录入或已删除",
	//产品模块
	ERROR_PRODUCT_NOT_EXIST: "产品未录入或已删除",
	//供应商模块
	ERROR_SUPPLIER_NOT_EXIST: "供应商未录入或已删除",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}

func Message(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": GetErrMsg(code),
		"data":    data,
	})
}

func MessageForList(c *gin.Context, code int, data interface{}, pageSize int, PageNo int, total int64) {
	c.JSON(http.StatusOK, gin.H{
		"status":   code,
		"message":  GetErrMsg(code),
		"data":     data,
		"pageSize": pageSize,
		"PageNo":   PageNo,
		"total":    total,
	})
}
