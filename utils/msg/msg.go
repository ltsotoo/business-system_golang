package msg

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	SUCCESS = 666
	ERROR   = 999

	//code = [1000-1100) 员工模块错误
	ERROR_EMPLOYEE_EXIST         = 1001
	ERROR_EMPLOYEE_NOT_EXIST     = 1002
	ERROR_EMPLOYEE_LOGIN_FAIL    = 1003
	ERROR_EMPLOYEE_INSERT        = 1011
	ERROR_EMPLOYEE_DELETE        = 1012
	ERROR_EMPLOYEE_UPDATE        = 1013
	ERROR_EMPLOYEE_SELECT        = 1014
	ERROR_EMPLOYEE_PASSWORD_FAIL = 1015

	//code = [1100-1200) 员工模块错误
	ERRPR_TOKEN_NOT_EXIST = 1101
	ERROR_TOKEN_EXPIRED   = 1102
	ERRPR_TOKEN_WRONG     = 1103
	ERROR_TOKEN_UID       = 1104

	//code = [1200-1300) 合同模块错误
	ERROR_CONTRACT_NOT_EXIST = 1201
	ERROR_CONTRACT_INSERT    = 1211
	ERROR_CONTRACT_DELETE    = 1212
	ERROR_CONTRACT_UPDATE    = 1213
	ERROR_CONTRACT_SELECT    = 1214

	ERROR_CONTRACT_UPDATE_STATUS   = 1221
	ERROR_CONTRACT_UPDATE_P_STATUS = 1222
	ERROR_CONTRACT_UPDATE_C_STATUS = 1223

	//code = [1300-1400) 任务模块错误
	ERROR_TASK_INSERT    = 1311
	ERROR_TASK_DELETE    = 1312
	ERROR_TASK_UPDATE    = 1313
	ERROR_TASK_SELECT    = 1314
	ERROR_TASK_NOT_MONEY = 1315

	//code = [1400-1500) 客户模块错误
	ERROR_CUSTOMER_NOT_EXIST         = 1401
	ERROR_CUSTOMER_COMPANY_NOT_EXIST = 1402
	ERROR_CUSTOMER_INSERT            = 1411
	ERROR_CUSTOMER_DELETE            = 1412
	ERROR_CUSTOMER_UPDATE            = 1413
	ERROR_CUSTOMER_SELECT            = 1414

	//code = [1500-1600) 产品模块错误
	ERROR_PRODUCT_NOT_EXIST = 1501
	ERROR_PRODUCT_INSERT    = 1511
	ERROR_PRODUCT_DELETE    = 1512
	ERROR_PRODUCT_UPDATE    = 1513
	ERROR_PRODUCT_SELECT    = 1514

	//code = [1600-1700) 供应商模块错误
	ERROR_SUPPLIER_NOT_EXIST = 1601
	ERROR_SUPPLIER_INSERT    = 1611
	ERROR_SUPPLIER_DELETE    = 1612
	ERROR_SUPPLIER_UPDATE    = 1613
	ERROR_SUPPLIER_SELECT    = 1614

	//code = [1700-2200) OADRP模块错误
	ERROR_OFFICE_NOT_EXIST = 1701
	ERROR_OFFICE_INSERT    = 1711
	ERROR_OFFICE_DELETE    = 1712
	ERROR_OFFICE_UPDATE    = 1713
	ERROR_OFFICE_SELECT    = 1714

	ERROR_DEPARTMENT_INSERT = 1911
	ERROR_DEPARTMENT_DELETE = 1912
	ERROR_DEPARTMENT_UPDATE = 1913
	ERROR_DEPARTMENT_SELECT = 1914

	ERROR_ROLE_INSERT = 2011
	ERROR_ROLE_DELETE = 2012
	ERROR_ROLE_UPDATE = 2013
	ERROR_ROLE_SELECT = 2014

	ERROR_PREMISSION_WORING = 2101
	ERROR_PERMISSION_INSERT = 2111
	ERROR_PERMISSION_DELETE = 2112
	ERROR_PERMISSION_UPDATE = 2113
	ERROR_PERMISSION_SELECT = 2114

	//code = [2200-2300) 财务模块错误
	ERROR_EXPENSE_INSERT     = 2211
	ERROR_EXPENSE_DELETE     = 2212
	ERROR_EXPENSE_UPDATE     = 2213
	ERROR_EXPENSE_SELECT     = 2214
	ERROR_EXPENSE_MONEY_LESS = 2215

	//code = [9900-10000) System模块错误
	ERROR_SYSTE_DIC_TYPE_INSERT = 9911
	ERROR_SYSTE_DIC_TYPE_DELETE = 9912
	ERROR_SYSTE_DIC_TYPE_SELECT = 9913
	ERROR_SYSTE_DIC_INSERT      = 9921
	ERROR_SYSTE_DIC_DELETE      = 9922
	ERROR_SYSTE_DIC_SELECT      = 9923

	ERROR_SYSTEM_SETTLEMENT = 9999
)

var codeMsg = map[int]string{
	SUCCESS: "请求成功",
	ERROR:   "请求失败",

	//员工模块
	ERROR_EMPLOYEE_EXIST:         "员工已录入",
	ERROR_EMPLOYEE_NOT_EXIST:     "员工未录入或已删除",
	ERROR_EMPLOYEE_LOGIN_FAIL:    "手机号或者密码错误",
	ERROR_EMPLOYEE_INSERT:        "员工录入失败",
	ERROR_EMPLOYEE_DELETE:        "员工删除失败",
	ERROR_EMPLOYEE_UPDATE:        "员工信息编辑失败",
	ERROR_EMPLOYEE_SELECT:        "员工信息查找失败",
	ERROR_EMPLOYEE_PASSWORD_FAIL: "密码错误",

	ERRPR_TOKEN_NOT_EXIST: "凭证不存在",
	ERROR_TOKEN_EXPIRED:   "凭证过期",
	ERRPR_TOKEN_WRONG:     "凭证格式异常",
	ERROR_TOKEN_UID:       "个人信息异常",

	//合同模块
	ERROR_CONTRACT_NOT_EXIST: "合同未录入或已删除",
	ERROR_CONTRACT_INSERT:    "合同录入失败",
	ERROR_CONTRACT_DELETE:    "合同删除失败",
	ERROR_CONTRACT_UPDATE:    "合同信息编辑失败",
	ERROR_CONTRACT_SELECT:    "合同信息查找失败",

	ERROR_CONTRACT_UPDATE_STATUS:   "合同状态修改失败",
	ERROR_CONTRACT_UPDATE_P_STATUS: "合同生产状态修改失败",
	ERROR_CONTRACT_UPDATE_C_STATUS: "合同回款状态修改失败",
	//任务模块
	ERROR_TASK_INSERT:    "任务添加失败",
	ERROR_TASK_DELETE:    "任务删除失败",
	ERROR_TASK_UPDATE:    "任务信息编辑失败",
	ERROR_TASK_SELECT:    "任务信息查找失败",
	ERROR_TASK_NOT_MONEY: "预存款金额不够",
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

	ERROR_DEPARTMENT_INSERT: "部门添加失败",
	ERROR_DEPARTMENT_DELETE: "部门删除失败",
	ERROR_DEPARTMENT_UPDATE: "部门信息编辑失败",
	ERROR_DEPARTMENT_SELECT: "部门信息查找失败",

	ERROR_ROLE_INSERT: "角色添加失败",
	ERROR_ROLE_DELETE: "角色删除失败",
	ERROR_ROLE_UPDATE: "角色信息编辑失败",
	ERROR_ROLE_SELECT: "角色信息查找失败",

	ERROR_PREMISSION_WORING: "没有该操作的权限",
	ERROR_PERMISSION_INSERT: "权限添加失败",
	ERROR_PERMISSION_DELETE: "权限删除失败",
	ERROR_PERMISSION_UPDATE: "权限信息编辑失败",
	ERROR_PERMISSION_SELECT: "权限信息查找失败",

	//财务模块错误
	ERROR_EXPENSE_INSERT:     "报销添加失败",
	ERROR_EXPENSE_DELETE:     "报销删除失败",
	ERROR_EXPENSE_UPDATE:     "报销信息编辑失败",
	ERROR_EXPENSE_SELECT:     "报销信息查找失败",
	ERROR_EXPENSE_MONEY_LESS: "可用报销余额不足",

	//System模块错误
	ERROR_SYSTE_DIC_TYPE_INSERT: "系统字典类型添加失败",
	ERROR_SYSTE_DIC_TYPE_DELETE: "系统字典类型删除失败",
	ERROR_SYSTE_DIC_TYPE_SELECT: "系统字典类型查找失败",
	ERROR_SYSTE_DIC_INSERT:      "系统字典添加失败",
	ERROR_SYSTE_DIC_DELETE:      "系统字典删除失败",
	ERROR_SYSTE_DIC_SELECT:      "系统字典查找失败",

	ERROR_SYSTEM_SETTLEMENT: "系统结算中",
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

func MessageForSystemSettlement(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  ERROR_SYSTEM_SETTLEMENT,
		"message": GetErrMsg(ERROR_SYSTEM_SETTLEMENT),
		"data":    nil,
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
