package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func CheckPermission(c *gin.Context) {
	fmt.Println(c.Request.URL)
	println(c.Request.Method)
	// permission, _ := model.SelectPermission(module, name)
	// if permission.ID > 0 {
	// 	Permission_uid_set := c.MustGet("Permission_uid_set").([]string)
	// 	for i := range Permission_uid_set {
	// 		if permission.UID == Permission_uid_set[i] {
	// 			return
	// 		}
	// 	}
	// }
	// msg.Message(c, msg.ERROR_PREMISSION_WORING, nil)
	// c.Abort()
}
