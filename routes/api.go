package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 注册网络相关路由
func RegisterAPIRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		//注册一个路由
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"hello": "world"})
		})
	}
}
