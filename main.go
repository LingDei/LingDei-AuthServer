package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

//	@title						灵嘚用户认证服务API
//	@version					v1.0.0
//	@description				灵嘚用户认证服务接口
//	@host						127.0.0.1:9000
//	@BasePath					/
//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization

func main() {
	// 初始化
	app := fiber.New()

	// CORS
	app.Use(cors.New(cors.Config{}))

	// 注册路由
	// regiserService(app)

	// 启动服务
	app.Listen(":9000")
}
