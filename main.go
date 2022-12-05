package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gohub/bootstarp"
)

func main() {
	// 初始化 Gin 实例
	r := gin.Default()
	//初始化路由
	bootstarp.SetupRoute(r)
	// 运行服务
	err := r.Run(":8080")
	if err != nil {
		fmt.Println(err.Error())
	}

}
