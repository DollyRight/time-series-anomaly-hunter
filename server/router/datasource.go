package router

import (
	"github.com/gin-gonic/gin"
	"server/api/v1"
)

func InitDataSourceRouter(rg *gin.RouterGroup) {
	rg.Group("datasource")
	{
		rg.POST("/", v1.AddDataSource)
	}
}
