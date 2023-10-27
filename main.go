package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	_ "LingDei-AuthServer/docs"
)

//	@title						灵嘚用户认证服务API
//	@version					v1.0.0
//	@description				灵嘚用户认证服务接口是一个基于 Fiber 的 RESTful API 服务，用于提供灵嘚（LingDei）的用户认证服务。
//	@description				注意，有 🦸 标识的接口需要管理员权限才能访问。
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
	regiserService(app)

	// 启动服务
	app.Listen(":9000")
}
