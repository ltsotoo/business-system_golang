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
	ERROR_EMPLOYEE_INSERT     = 10011
	ERROR_EMPLOYEE_DELETE     = 10012
	ERROR_EMPLOYEE_UPDATE     = 10013
	ERROR_EMPLOYEE_SELECT     = 10014

	ERRPR_TOKEN_NOT_EXIST = 11001
	ERROR_TOKEN_EXPIRED   = 11002
	ERRPR_TOKEN_WRONG     = 11003

	ERROR_PREMISSION_WORING = 12001
	//code = 20000-30000 合同模块错误
	ERROR_CONTRACT_NOT_EXIST = 20001
	ERROR_CONTRACT_INSERT    = 20011
	ERROR_CONTRACT_DELETE    = 20012
	ERROR_CONTRACT_UPDATE    = 20013
	ERROR_CONTRACT_SELECT    = 20014
	//code = 30000-40000 任务模块错误
	ERROR_TASK_INSERT = 30011
	ERROR_TASK_DELETE = 30012
	ERROR_TASK_UPDATE = 30013
	ERROR_TASK_SELECT = 30014
	//code = 40000-50000 客户模块错误
	ERROR_CUSTOMER_NOT_EXIST         = 40001
	ERROR_CUSTOMER_COMPANY_NOT_EXIST = 40002
	ERROR_CUSTOMER_INSERT            = 40011
	ERROR_CUSTOMER_DELETE            = 40012
	ERROR_CUSTOMER_UPDATE            = 40013
	ERROR_CUSTOMER_SELECT            = 40014
	//code = 50000-60000 产品模块错误
	ERROR_PRODUCT_NOT_EXIST = 50001
	ERROR_PRODUCT_INSERT    = 50011
	ERROR_PRODUCT_DELETE    = 50012
	ERROR_PRODUCT_UPDATE    = 50013
	ERROR_PRODUCT_SELECT    = 50014
	//code = 60000-70000 供应商模块错误
	ERROR_SUPPLIER_NOT_EXIST = 60001
	ERROR_SUPPLIER_INSERT    = 60011
	ERROR_SUPPLIER_DELETE    = 60012
	ERROR_SUPPLIER_UPDATE    = 60013
	ERROR_SUPPLIER_SELECT    = 60014
	//code = 70000-80000 OADRP模块错误
	ERROR_OFFICE_NOT_EXIST = 70001
	ERROR_OFFICE_INSERT    = 70011
	ERROR_OFFICE_DELETE    = 70012
	ERROR_OFFICE_UPDATE    = 70013
	ERROR_OFFICE_SELECT    = 70014

	ERROR_AREA_INSERT = 71011
	ERROR_AREA_DELETE = 71012
	ERROR_AREA_UPDATE = 71013
	ERROR_AREA_SELECT = 71014

	ERROR_DEPARTMENT_INSERT = 72011
	ERROR_DEPARTMENT_DELETE = 72012
	ERROR_DEPARTMENT_UPDATE = 72013
	ERROR_DEPARTMENT_SELECT = 72014

	ERROR_ROLE_INSERT = 73011
	ERROR_ROLE_DELETE = 73012
	ERROR_ROLE_UPDATE = 73013
	ERROR_ROLE_SELECT = 73014

	ERROR_PERMISSION_INSERT = 74011
	ERROR_PERMISSION_DELETE = 74012
	ERROR_PERMISSION_UPDATE = 74013
	ERROR_PERMISSION_SELECT = 74014
	//code = 80000-90000 System模块错误
	ERROR_SYSTE_DIC_TYPE_INSERT = 80011
	ERROR_SYSTE_DIC_TYPE_DELETE = 80012
	ERROR_SYSTE_DIC_TYPE_SELECT = 80013
	ERROR_SYSTE_DIC_INSERT      = 80021
	ERROR_SYSTE_DIC_DELETE      = 80022
	ERROR_SYSTE_DIC_SELECT      = 80023
)

var codeMsg = map[int]string{
	SUCCESS: "请求成功",
	ERROR:   "请求失败",

	//员工模块
	ERROR_EMPLOYEE_EXIST:      "员工已录入",
	ERROR_EMPLOYEE_NOT_EXIST:  "员工未录入或已删除",
	ERROR_EMPLOYEE_LOGIN_FAIL: "手机号或者密码错误",
	ERROR_EMPLOYEE_INSERT:     "员工录入失败",
	ERROR_EMPLOYEE_DELETE:     "员工删除失败",
	ERROR_EMPLOYEE_UPDATE:     "员工信息编辑失败",
	ERROR_EMPLOYEE_SELECT:     "员工信息查找失败",

	ERRPR_TOKEN_NOT_EXIST: "凭证不存在",
	ERROR_TOKEN_EXPIRED:   "凭证过期",
	ERRPR_TOKEN_WRONG:     "凭证格式异常",

	ERROR_PREMISSION_WORING: "没有该操作的权限",
	//合同模块
	ERROR_CONTRACT_NOT_EXIST: "合同未录入或已删除",
	ERROR_CONTRACT_INSERT:    "合同录入失败",
	ERROR_CONTRACT_DELETE:    "合同删除失败",
	ERROR_CONTRACT_UPDATE:    "合同信息编辑失败",
	ERROR_CONTRACT_SELECT:    "合同信息查找失败",
	//任务模块
	ERROR_TASK_INSERT: "任务添加失败",
	ERROR_TASK_DELETE: "任务删除失败",
	ERROR_TASK_UPDATE: "任务信息编辑失败",
	ERROR_TASK_SELECT: "任务信息查找失败",
	//客户模块
	ERROR_CUSTOMER_NOT_EXIST:         "客户未录入或已删除",
	ERROR_CUSTOMER_COMPANY_NOT_EXIST: "客户公司未录入或已删除",
	ERROR_CUSTOMER_INSERT:            "客户录入失败",
	ERROR_CUSTOMER_DELETE:            "客户删除失败",
	ERROR_CUSTOMER_UPDATE:            "客户信息编辑失败",
	ERROR_CUSTOMER_SELECT:            "客户信息查找失败",
	//产品模块
	ERROR_PRODUCT_NOT_EXIST: "产品未录入或已删除",
	ERROR_PRODUCT_INSERT:    "产品录入失败",
	ERROR_PRODUCT_DELETE:    "产品删除失败",
	ERROR_PRODUCT_UPDATE:    "产品信息编辑失败",
	ERROR_PRODUCT_SELECT:    "产品信息查找失败",
	//供应商模块
	ERROR_SUPPLIER_NOT_EXIST: "供应商未录入或已删除",
	ERROR_SUPPLIER_INSERT:    "供应商录入失败",
	ERROR_SUPPLIER_DELETE:    "供应商删除失败",
	ERROR_SUPPLIER_UPDATE:    "供应商信息编辑失败",
	ERROR_SUPPLIER_SELECT:    "供应商信息查找失败",
	//OADRP模块
	ERROR_OFFICE_NOT_EXIST: "办事处未添加或已删除",
	ERROR_OFFICE_INSERT:    "办事处添加失败",
	ERROR_OFFICE_DELETE:    "办事处删除失败",
	ERROR_OFFICE_UPDATE:    "办事处信息编辑失败",
	ERROR_OFFICE_SELECT:    "办事处信息查找失败",

	ERROR_AREA_INSERT: "地区添加失败",
	ERROR_AREA_DELETE: "地区删除失败",
	ERROR_AREA_UPDATE: "地区信息编辑失败",
	ERROR_AREA_SELECT: "地区信息查找失败",

	ERROR_DEPARTMENT_INSERT: "部门添加失败",
	ERROR_DEPARTMENT_DELETE: "部门删除失败",
	ERROR_DEPARTMENT_UPDATE: "部门信息编辑失败",
	ERROR_DEPARTMENT_SELECT: "部门信息查找失败",

	ERROR_ROLE_INSERT: "角色添加失败",
	ERROR_ROLE_DELETE: "角色删除失败",
	ERROR_ROLE_UPDATE: "角色信息编辑失败",
	ERROR_ROLE_SELECT: "角色信息查找失败",

	ERROR_PERMISSION_INSERT: "权限添加失败",
	ERROR_PERMISSION_DELETE: "权限删除失败",
	ERROR_PERMISSION_UPDATE: "权限信息编辑失败",
	ERROR_PERMISSION_SELECT: "权限信息查找失败",
	//System模块错误
	ERROR_SYSTE_DIC_TYPE_INSERT: "系统字典类型添加失败",
	ERROR_SYSTE_DIC_TYPE_DELETE: "系统字典类型删除失败",
	ERROR_SYSTE_DIC_TYPE_SELECT: "系统字典类型查找失败",
	ERROR_SYSTE_DIC_INSERT:      "系统字典添加失败",
	ERROR_SYSTE_DIC_DELETE:      "系统字典删除失败",
	ERROR_SYSTE_DIC_SELECT:      "系统字典查找失败",
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

func MessageForNotPermission(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  ERROR_PREMISSION_WORING,
		"message": GetErrMsg(ERROR_PREMISSION_WORING),
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
