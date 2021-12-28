package v1

import (
	"business-system_golang/model"
	"business-system_golang/utils/msg"

	"github.com/gin-gonic/gin"
)

func EntryOffice(c *gin.Context) {
	var office model.Office
	_ = c.ShouldBindJSON(&office)
	code = model.InsertOffice(&office)
	msg.Message(c, code, office)
}

func DelOffice(c *gin.Context) {
	uid := c.Param("uid")
	code = model.DeleteOffice(uid)
	msg.Message(c, code, nil)
}

func EditOffice(c *gin.Context) {
	var office model.Office
	_ = c.ShouldBindJSON(&office)
	code = model.UpdateOffice(&office)
	msg.Message(c, code, office)
}

func QueryOffice(c *gin.Context) {
	var office model.Office
	uid := c.Param("uid")
	office, code = model.SelectOffice(uid)
	msg.Message(c, code, office)
}

func QueryOffices(c *gin.Context) {
	var office model.Office
	var offices []model.Office
	_ = c.ShouldBindJSON(&office)
	offices, code = model.SelectOffices(&office)
	msg.Message(c, code, offices)
}

func EntryDepartment(c *gin.Context) {
	var department model.Department
	_ = c.ShouldBindJSON(&department)
	if department.OfficeUID != "" {
		code = model.InsertDepartment(&department)
	} else {
		code = msg.ERROR
	}
	msg.Message(c, code, department)

}

func DelDepartment(c *gin.Context) {
	uid := c.Param("uid")
	code = model.DeleteDepartment(uid)
	msg.Message(c, code, nil)
}

func EditDepartment(c *gin.Context) {
	var department model.Department
	_ = c.ShouldBindJSON(&department)
	code = model.UpdateDepartment(&department)
	msg.Message(c, code, department)
}

func QueryDepartments(c *gin.Context) {
	var departments []model.Department
	var department model.Department
	_ = c.ShouldBindJSON(&department)
	departments, code = model.SelectDepartments(&department)
	msg.Message(c, code, departments)
}

func AddRole(c *gin.Context) {
	var role model.Role
	_ = c.ShouldBindJSON(&role)
	code = model.InsertRole(&role)
	msg.Message(c, code, role)
}

func EditRole(c *gin.Context) {
	var role model.Role
	_ = c.ShouldBindJSON(&role)
	code = model.UpdateRole(&role)
	msg.Message(c, code, role)
}

func QueryRole(c *gin.Context) {
	var role model.Role
	uid := c.Param("uid")
	role, code = model.SelectRole(uid)
	msg.Message(c, code, role)
}

func QueryRoles(c *gin.Context) {
	var roles []model.Role
	roles, code = model.SelectRoles()
	msg.Message(c, code, roles)
}

func QueryAllRoles(c *gin.Context) {
	var roles []model.Role
	name := c.DefaultQuery("name", "")
	roles, code = model.SelectAllRoles(name)
	msg.Message(c, code, roles)
}

func QueryPermissions(c *gin.Context) {
	var permissions []model.Permission
	permissions, code = model.SelectPermissions()
	msg.Message(c, code, permissions)
}
