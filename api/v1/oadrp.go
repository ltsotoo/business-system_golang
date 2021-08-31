package v1

import (
	"business-system_golang/model"
	"business-system_golang/utils/msg"
	"business-system_golang/utils/rbac"

	"github.com/gin-gonic/gin"
)

func EntryOffice(c *gin.Context) {
	code = rbac.Check(c, "office", "all")
	if code == msg.ERROR {
		msg.MessageForNotPermission(c)
	} else {
		var office model.Office
		_ = c.ShouldBindJSON(&office)
		code = model.InsertOffice(&office)
		msg.Message(c, code, office)
	}
}

func DelOffice(c *gin.Context) {
	code = rbac.Check(c, "office", "all")
	if code == msg.ERROR {
		msg.MessageForNotPermission(c)
	} else {
		uid := c.Param("uid")
		code = model.DeleteOffice(uid)
		msg.Message(c, code, nil)
	}
}

func QueryOffices(c *gin.Context) {
	name := c.Query("name")
	var offices []model.Office
	offices, code = model.SelectOffices(name)
	msg.Message(c, code, offices)
}

func EntryArea(c *gin.Context) {
	code = rbac.Check(c, "area", "all")
	if code == msg.ERROR {
		msg.MessageForNotPermission(c)
	} else {
		var area model.Area
		_ = c.ShouldBindJSON(&area)
		code = model.InsertArea(&area)
		msg.Message(c, code, area)
	}
}

func DelArea(c *gin.Context) {
	code = rbac.Check(c, "area", "all")
	if code == msg.ERROR {
		msg.MessageForNotPermission(c)
	} else {
		uid := c.Param("uid")
		code = model.DeleteArea(uid)
		msg.Message(c, code, nil)
	}
}

func QueryAreas(c *gin.Context) {
	var areas []model.Area
	var areaQuery model.AreaQuery
	_ = c.ShouldBindJSON(&areaQuery)
	areas, code = model.SelectAreas(&areaQuery)
	msg.Message(c, code, areas)
}

func EditArea(c *gin.Context) {
	code = rbac.Check(c, "area", "all")
	if code == msg.ERROR {
		msg.MessageForNotPermission(c)
	} else {
		var area model.Area
		_ = c.ShouldBindJSON(&area)
		code = model.UpdateArea(&area)
		msg.Message(c, code, area)
	}
}

func EntryDepartment(c *gin.Context) {
	code = rbac.Check(c, "department", "all")
	if code == msg.ERROR {
		msg.MessageForNotPermission(c)
	} else {
		var department model.Department
		_ = c.ShouldBindJSON(&department)
		if department.OfficeUID == "" {
			visitor, _ := model.SelectEmployee(c.MustGet("employeeUID").(string))
			department.OfficeUID = visitor.OfficeUID
		}
		code = model.InsertDepartment(&department)
		msg.Message(c, code, department)
	}
}

func DelDepartment(c *gin.Context) {
	code = rbac.Check(c, "department", "all")
	if code == msg.ERROR {
		msg.MessageForNotPermission(c)
	} else {
		uid := c.Param("uid")
		code = model.DeleteDepartment(uid)
		msg.Message(c, code, nil)
	}
}

func QueryDepartments(c *gin.Context) {
	var departments []model.Department
	var department model.Department
	_ = c.ShouldBindJSON(&department)
	if department.OfficeUID == "" {
		visitor, _ := model.SelectEmployee(c.MustGet("employeeUID").(string))
		department.OfficeUID = visitor.OfficeUID
	}
	departments, code = model.SelectDepartments(&department)
	msg.Message(c, code, departments)
}

func AddRole(c *gin.Context) {
	code = rbac.Check(c, "role", "all")
	if code == msg.ERROR {
		msg.MessageForNotPermission(c)
	} else {
		var role model.Role
		_ = c.ShouldBindJSON(&role)
		code = model.InsertRole(&role)
		msg.Message(c, code, role)
	}
}

func EditRole(c *gin.Context) {
	code = rbac.Check(c, "role", "all")
	if code == msg.ERROR {
		msg.MessageForNotPermission(c)
	} else {
		var role model.Role
		_ = c.ShouldBindJSON(&role)
		code = model.UpdateRole(&role)
		msg.Message(c, code, role)
	}
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
