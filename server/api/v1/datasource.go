package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/utils/response"
)

type DataSourceInfo struct {
	Type     string `json:"type" binding:"required"`
	Host     string `json:"host" binding:"required""`
	Port     string `json:"port" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" bson:"password"`
}

func AddDataSource(c *gin.Context) {
	var dataSourceInfo DataSourceInfo
	respBody := map[string]interface{}{}
	if err := c.ShouldBind(&dataSourceInfo); err != nil {
		c.JSON(http.StatusBadRequest, response.FailRespBody("解析json请求体时发生错误", respBody))
	}
	//将数据源信息添加到数据库中
	c.JSON(http.StatusOK, response.SuccessRespBody("添加数据源成功", respBody))
}
