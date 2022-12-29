package routes

import (
	"github.com/gin-gonic/gin"
	"gohub/app/http/controllers/auth"
	"net/http"
)

// 注册网络相关路由
func RegisterAPIRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		authGroup := v1.Group("auth")
		{
			succ := new(auth.SignupController)
			//check phone if exist
			authGroup.POST("/signup/phone/exist", succ.IsPhoneExist)
			authGroup.POST("/signup/email/exist", succ.IsEmailExist)
		}
		//注册一个路由
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"hello": "world"})
		})
	}
}
