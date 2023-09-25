package initialize

import (
	"github.com/gin-gonic/gin"
	"server/router"
)

func Router() *gin.Engine {
	r := gin.Default()

	// 路由组v1
	router.InitRouter(r)

	return r
}
