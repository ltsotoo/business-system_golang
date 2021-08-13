package v1

import (
	"business-system_golang/model"
	"business-system_golang/utils/msg"
	"strconv"

	"github.com/gin-gonic/gin"
)

func QueryOffices(c *gin.Context) {
	name := c.Query("name")
	var offices []model.Office
	offices, code = model.SelectOffices(name)
	msg.Message(c, code, offices)
}

func QueryAreas(c *gin.Context) {
	var areas []model.Area
	areas, code = model.SelectAreas()
	msg.Message(c, code, areas)
}

func QueryAreasByOfficeID(c *gin.Context) {
	var areas []model.Area
	officeID, _ := strconv.Atoi(c.Query("officeID"))
	areas, code = model.SelectAreasByOfficeID(officeID)
	msg.Message(c, code, areas)
}

func QueryDepartmentsByOfficeID(c *gin.Context) {
	var departments []model.Department
	officeID, _ := strconv.Atoi(c.Query("officeID"))
	departments, code = model.SelectDepartmentsByOfficeID(officeID)
	msg.Message(c, code, departments)
}

func QueryRoles(c *gin.Context) {
	var roles []model.Role
	roles, code = model.SelectRoles()
	msg.Message(c, code, roles)
}

func QueryPermissions(c *gin.Context) {
	var permissions []model.Permission
	permissions, code = model.SelectPermissions()
	msg.Message(c, code, permissions)
}
