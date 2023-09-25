package response

import (
	"github.com/gin-gonic/gin"
)

func SuccessRespBody(mes string, data map[string]interface{}) gin.H {
	return gin.H{"status": true, "mes": mes, "data": data}
}

func FailRespBody(mes string, data map[string]interface{}) gin.H {
	return gin.H{"status": false, "mes": mes, "data": data}
}
