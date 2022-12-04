package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化 Gin 实例
	r := gin.Default()
	// 使用中间件
	r.Use(gin.Logger(), gin.Recovery())
	// 注册一个路由
	r.GET("/", func(c *gin.Context) {

		// 以 JSON 格式响应
		c.JSON(http.StatusOK, gin.H{
			"Hello": "World!",
		})
	})
	// deal 404
	r.NoRoute(func(c *gin.Context) {
		// get header
		acceptstr := c.Request.Header.Get("Accept")
		if strings.Contains(acceptstr, "text/html") {
			c.String(http.StatusNotFound, "页面返回404")
		} else {
			// default return JSON
			c.JSON(http.StatusNotFound, gin.H{
				"erros":         404,
				"error_message": "no route ,please identify url or request method ",
			})
		}
	})

	// 运行服务
	r.Run(":8000")

}
