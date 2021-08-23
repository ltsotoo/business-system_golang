package v1

import (
	"business-system_golang/model"
	"business-system_golang/utils/msg"
	"strconv"

	"github.com/gin-gonic/gin"
)

func EntryOffice(c *gin.Context) {
	var office model.Office
	_ = c.ShouldBindJSON(&office)
	code = model.CreateOffice(&office)
	msg.Message(c, code, office)
}

func DelOffice(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteOffice(id)
	msg.Message(c, code, nil)
}

func QueryOffices(c *gin.Context) {
	name := c.Query("name")
	var offices []model.Office
	offices, code = model.SelectOffices(name)
	msg.Message(c, code, offices)
}

func EntryArea(c *gin.Context) {
	var area model.Area
	_ = c.ShouldBindJSON(&area)
	code = model.CreateArea(&area)
	msg.Message(c, code, area)
}

func DelArea(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteArea(id)
	msg.Message(c, code, nil)
}

func QueryAreas(c *gin.Context) {
	var areas []model.Area
	var area model.Area
	_ = c.ShouldBindJSON(&area)
	areas, code = model.SelectAreas(&area)
	msg.Message(c, code, areas)
}

func EditArea(c *gin.Context) {
	var area model.Area
	_ = c.ShouldBindJSON(&area)
	code = model.UpdateArea(&area)
	msg.Message(c, code, area)
}

func EntryDepartment(c *gin.Context) {
	var department model.Department
	_ = c.ShouldBindJSON(&department)
	code = model.CreateDepartment(&department)
	msg.Message(c, code, department)
}

func DelDepartment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteDepartment(id)
	msg.Message(c, code, nil)
}

func QueryDepartments(c *gin.Context) {
	var departments []model.Department
	var department model.Department
	_ = c.ShouldBindJSON(&department)
	if department.OfficeID == 0 {
		employee, _ := model.SelectEmployee(int(c.MustGet("employeeID").(uint)))
		department.OfficeID = employee.OfficeID
	}
	departments, code = model.SelectDepartments(&department)
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
