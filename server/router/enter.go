package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/global"
)

func InitRouter(r *gin.Engine) {
	v1 := r.Group(fmt.Sprintf("/%s/v1", global.Gva.GvaConfigs.Server.ApiPrefix))
	{
		InitDataSourceRouter(v1)
	}
}
