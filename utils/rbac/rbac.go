package rbac

import (
	"business-system_golang/model"
	"business-system_golang/utils/msg"

	"github.com/gin-gonic/gin"
)

//权限检查
func Check(c *gin.Context, module string, name string) (code int) {
	permission, _ := model.SelectPermission(module, name)
	if permission.ID == 0 {
		code = msg.SUCCESS
		return
	}
	Permission_uid_set := c.MustGet("Permission_uid_set").([]string)
	for i := range Permission_uid_set {
		if permission.UID == Permission_uid_set[i] {
			code = msg.SUCCESS
			return
		}
	}
	code = msg.ERROR
	return
}
