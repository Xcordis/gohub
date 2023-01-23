package main

import (
	"flag"
	"fmt"
	"gohub/bootstarp"
	btsConfig "gohub/config"
	"gohub/pkg/config"

	"github.com/gin-gonic/gin"
)

func init() {
	// 加载 config 目录下的配置信息
	btsConfig.Initialize()
}

func main() {

	// 配置初始化，依赖命令行 --env 参数
	var env string
	flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")
	flag.Parse()
	config.InitConfig(env)
	//init log library
	bootstarp.SetupLogger()

	// 故此设置为 release，有特殊情况手动改为 debug 即可
	gin.SetMode(gin.ReleaseMode)
	// 初始化数据库
	bootstarp.SetupDB()

	// new 一个 Gin Engine 实例
	router := gin.New()

	// 初始化路由绑定
	bootstarp.SetupRoute(router)

	// 运行服务
	err := router.Run(":" + config.Get("app.port"))
	if err != nil {
		// 错误处理，端口被占用了或者其他错误
		fmt.Println(err.Error())
	}
}
