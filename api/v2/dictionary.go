package v2

import (
	"business-system_golang/model"
	"business-system_golang/utils/msg"

	"github.com/gin-gonic/gin"
)

func AddDictionaryType(c *gin.Context) {
	var dictionaryType model.DictionaryType
	_ = c.ShouldBindJSON(&dictionaryType)

	code = model.InsertDictionaryType(&dictionaryType)

	msg.Message(c, code, dictionaryType)
}
