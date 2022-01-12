package v1

import (
	"business-system_golang/middleware"
	"business-system_golang/model"
	"business-system_golang/utils/msg"
	"math"
	"sort"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var employee model.Employee
	var token string
	var urls []model.Url
	var nos []string
	_ = c.ShouldBindJSON(&employee)

	employee, code = model.CheckLogin(employee.Phone, employee.Password)
	if code == msg.SUCCESS {
		permission_uid_set := model.SelectAllPermission(employee.UID, employee.DepartmentUID)
		token, _ = middleware.SetToken(employee.UID, employee.OfficeUID, employee.DepartmentUID, permission_uid_set)
		token = "Bearer " + token
		//router导航
		urls = model.SelectUrls(permission_uid_set)
		//前端权限
		nos = model.SelectPermissionsNo(permission_uid_set)
	}
	msg.Message(c, code, gin.H{"employee": employee, "token": token, "urls": urls, "nos": nos})
}

func Regist(c *gin.Context) {
	var employee model.Employee
	_ = c.ShouldBindJSON(&employee)
	code = model.CheckEmployee(employee.Phone)
	if code == msg.ERROR_EMPLOYEE_NOT_EXIST {
		code = model.InsertEmployee(&employee)
	}
	msg.Message(c, code, employee)
}

func TopList(c *gin.Context) {
	var office model.Office
	var offices, offices1, offices2 []model.Office
	offices, code = model.SelectOffices(&office)
	offices1, offices2 = model.SelectTopList()
	permissionUIDs := c.MustGet("Permission_uid_set").([]string)
	officeUID := c.MustGet("officeUID").(string)
	seeAll := false

	for k := range permissionUIDs {
		if permissionUIDs[k] == "55c96c98d62d48e29f0326e6185f21eb" {
			seeAll = true
			break
		}
	}
	for i := range offices {
		for j := range offices1 {
			if offices[i].UID == offices1[j].UID {
				offices[i].NotPayment = offices[i].NotPayment + offices1[j].Money
				break
			}
		}
		for k := range offices2 {
			if offices[i].UID == offices2[k].UID {
				offices[i].NotPayment = offices[i].NotPayment + offices2[k].Money
				break
			}
		}
		if offices[i].TaskLoad == 0 {
			offices[i].FinalPercentages = offices[i].TargetLoad * 100
		} else {
			offices[i].FinalPercentages = round((offices[i].TargetLoad / offices[i].TaskLoad * 100), 2)
		}
		if !seeAll && officeUID != offices[i].UID {
			offices[i].TargetLoad = 0
			offices[i].TaskLoad = 0
			offices[i].NotPayment = 0
		}
	}
	sort.SliceStable(offices, func(i, j int) bool {
		return offices[i].FinalPercentages > offices[j].FinalPercentages
	})
	msg.Message(c, code, offices)
}

func round(f float64, n int) float64 {
	pow10_n := math.Pow10(n)
	return math.Trunc((f+0.5/pow10_n)*pow10_n) / pow10_n
}
