package main

import (
	"LingDei_AuthServer/config"

	_ "LingDei_AuthServer/docs"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/gofiber/swagger"

	handler "LingDei_AuthServer/handler"
)

var jwtConfig = jwtware.New(jwtware.Config{
	SigningMethod: "RS256",
	SigningKey:    handler.PrivateKey.Public(),
	ErrorHandler:  config.CustomJwtError(),
})

func regiserService(app *fiber.App) {
	//Swagger
	app.Get("/swagger/*", swagger.HandlerDefault)

	//User
	user := app.Group("user")
	user.Post("/register", handler.RegisterHandler)
	user.Get("/exist", handler.CheckUserExistHandler)

	//Token
	token := app.Group("token")
	token.Post("/get", handler.GetToken)

	// 使用 JWT Middleware
	//以下为权限控制接口
	app.Use(jwtConfig)

	//User
	user.Post("/changepwd", handler.ChangePWDHandler)
	user.Post("/update", handler.UpdateUserHandler)
	user.Get("/getusers", handler.GetUsersHandler)
	user.Get("/getid", handler.GetUserIDHandler)
	user.Delete("/delete", handler.DeleteUserHandler)
	user.Post("/logout", handler.LogoutHandler)

	//Profile
	profile := app.Group("profile")
	profile.Get("/get", handler.GetProfileHandler)
	profile.Post("/update", handler.UpdateProfileHandler)
	profile.Delete("/delete", handler.DeleteProfileHandler)
	profile.Post("/avatar", handler.UploadAvatorHandler)
}
