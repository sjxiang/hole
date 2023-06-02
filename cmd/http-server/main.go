package main

import (
	"log"
	"os"
)

func main() {
	// set trial key for self-host users
	os.Setenv("HOLE_SECRET_KEY", "8xEMrWkBARcDDYQ")
	// init
	server, err := Initialize()
	if err != nil {
		log.Panic(err)
	}
	// TODO: gracefulStop
	server.Start()
}

// func main() {
// 	r := NewRouter()
// 	r.Run(":8000")
// }

// func NewRouter() *gin.Engine {
// 	// 设置为 Release，为的是默认在启动中不输出调试信息
// 	gin.SetMode(gin.ReleaseMode)
	
// 	r := gin.New()
	
// 	// 加载页面
// 	r.LoadHTMLGlob("D:\\workspace\\src\\hole\\internal\\web\\template" + "/*")
	
// 	// 业务绑定路由操作
// 	Routes(r)
	
// 	return r
// }

// // Routes 绑定业务层路由
// func Routes(r *gin.Engine) {

// 	// 如果配置了swagger，则显示swagger的中间件
// 	// if configService.GetBool("app.swagger") == true {
// 	// 	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
// 	// }

// 	// 用户模块
// 	// user.RegisterRoutes(r)
// 	// 问答模块
// 	// qa.RegisterRoutes(r)
// }
