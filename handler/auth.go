package handler

import (
	"LingDei-AuthServer/method"
	"LingDei-AuthServer/model"
	"LingDei-AuthServer/utils"
	"crypto/rsa"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

var (
	PrivateKey *rsa.PrivateKey
)

func init() {
	var err error
	PrivateKey, err = utils.LoadPrivateKey()
	if err != nil {
		panic(err)
	}
}

// GetToken godoc
//
//	@Summary		获取 Token
//	@Description	通过用户名和密码获取token
//	@Tags			token
//	@Accept			json
//	@Produce		json
//	@Param			username	query		string	true	"UserName"
//	@Param			password	query		string	true	"Password"
//	@Success		200			{object}	model.TokenResp
//	@Failure		400			{object}	model.OperationResp
//	@Router			/token/get [post]
func GetToken(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// 查找用户并验证密码，如不正确则返回未授权错误
	if err := method.CheckPwdValid(username, password); err != nil {
		return c.JSON(model.OperationResp{Code: 400, Msg: err.Error()})
	}

	//获取用户基本信息
	user, err := method.GetUser(username)
	if err != nil {
		return c.JSON(model.OperationResp{Code: 400, Msg: err.Error()})
	}

	// 创建用户凭证
	claims := jwt.MapClaims{
		"id":       user.ID,
		"username": user.UserName,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}

	// 创建token
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	// 签发token
	t, err := token.SignedString(PrivateKey)
	if err != nil {
		log.Printf("token.SignedString: %v", err)
		return c.JSON(model.OperationResp{Code: 400, Msg: err.Error()})
	}

	return c.JSON(model.TokenResp{Code: 200, Token: t, Role: user.Role})
}

// LogoutHandler godoc
//
//	@Summary		Token注销
//	@Description	用户登陆注销，本接口暂未实现实际功能，请在前端清除token
//	@Tags			用户
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.OperationResp
//	@Failure		400	{object}	model.OperationResp
//	@Security		ApiKeyAuth
//	@Router			/user/logout [delete]
func LogoutHandler(c *fiber.Ctx) error {
	return c.JSON(model.OperationResp{Code: 200, Msg: "ok"})
}
