package handler

import (
	"LingDei-AuthServer/config"
	"LingDei-AuthServer/method"
	"LingDei-AuthServer/model"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Register godoc
//
//	@Summary		用户注册
//	@Description	新用户注册
//	@Tags			用户
//	@Accept			json
//	@Produce		json
//	@Param			username	query		string	true	"UserName"
//	@Param			password	query		string	true	"Password"
//	@Success		200			{object}	model.OperationResp
//	@Failure		400			{object}	model.OperationResp
//	@Router			/user/register [post]
func RegisterHandler(c *fiber.Ctx) error {
	var user model.User
	user.ID = uuid.NewString() //生成随机id
	user.UserName = c.FormValue("username")
	user.Password = c.FormValue("password")

	if err := method.AddUser(user); err != nil {
		return c.JSON(model.OperationResp{Code: 400, Msg: err.Error()})
	}
	return c.JSON(model.OperationResp{Code: 200, Msg: "ok"})
}

// GetUsers godoc
//
//	@Summary		获取用户列表🦸‍♂️
//	@Description	获取用户列表，可选参数username的作用是搜索所有包括该参数的用户名，本接口对管理员权限的要求为等级6
//	@Tags			用户
//	@Accept			json
//	@Produce		json
//	@Param			username	query		string	false	"UserName"
//	@Success		200			{object}	model.UsersResp
//	@Failure		400			{object}	model.OperationResp
//	@Security		ApiKeyAuth
//	@Router			/user/getusers [get]
func GetUsersHandler(c *fiber.Ctx) error {
	UserName := c.FormValue("username")

	//检查用户权限等级
	if err := method.CheckUserRole(method.GetUserFromToken(c).Role, config.Admin_ROLE_LEVEL); err != nil {
		return c.JSON(model.OperationResp{Code: 400, Msg: err.Error()})
	}

	users, err := method.GetUsers(UserName)
	if err != nil {
		return c.JSON(model.OperationResp{Code: 400, Msg: err.Error()})
	}
	return c.JSON(model.UsersResp{Code: 200, Users: users})
}

// GetUserID godoc
//
//	@Summary		获取用户id
//	@Description	通过用户名获取用户id
//	@Tags			用户
//	@Accept			json
//	@Produce		json
//	@Param			username	query		string	true	"UserName"
//	@Success		200			{object}	model.UIDResp
//	@Failure		400			{object}	model.OperationResp
//	@Security		ApiKeyAuth
//	@Router			/user/getid [get]
func GetUserIDHandler(c *fiber.Ctx) error {
	UserName := c.FormValue("username")

	if UserName == "" {
		return c.JSON(model.OperationResp{Code: 400, Msg: "username is required"})
	}

	userID, err := method.GetUserID(UserName)
	if err != nil {
		return c.JSON(model.OperationResp{Code: 400, Msg: err.Error()})
	}
	return c.JSON(model.UIDResp{Code: 200, ID: userID})
}

// ChangePWD godoc
//
//	@Summary		修改用户密码
//	@Description	修改用户密码，username通过token取得，只需传入新密码、旧密码两个参数即可
//	@Tags			用户
//	@Accept			json
//	@Produce		json
//	@Param			old_password	query		string	true	"旧密码"
//	@Param			new_password	query		string	true	"新密码"
//	@Success		200				{object}	model.OperationResp
//	@Failure		400				{object}	model.OperationResp
//	@Security		ApiKeyAuth
//	@Router			/user/changepwd [post]
func ChangePWDHandler(c *fiber.Ctx) error {
	var user model.User
	user.UserName = method.GetUserFromToken(c).UserName
	oldPassword := c.FormValue("old_password")
	user.Password = c.FormValue("new_password")

	if user.Password == "" {
		return c.JSON(model.OperationResp{Code: 400, Msg: "new_password is required"})
	}

	//验证旧密码是否正确
	if err := method.CheckPwdValid(user.UserName, oldPassword); err != nil {
		return c.JSON(model.OperationResp{Code: 400, Msg: err.Error()})
	}

	//修改密码
	if err := method.UpdateUser(user); err != nil {
		return c.JSON(model.OperationResp{Code: 400, Msg: err.Error()})
	}
	return c.JSON(model.OperationResp{Code: 200, Msg: "ok"})
}

// ChangePWD godoc
//
//	@Summary		修改任意用户🦸‍♂️
//	@Description	修改任意用户,管理员权限要求为9
//	@Tags			用户
//	@Accept			json
//	@Produce		json
//	@Param			username	query		string	true	"用户名"
//	@Param			password	query		string	true	"新密码"
//	@Param			role		query		int		true	"权限等级1~9"
//	@Success		200			{object}	model.OperationResp
//	@Failure		400			{object}	model.OperationResp
//	@Security		ApiKeyAuth
//	@Router			/user/update [post]
func UpdateUserHandler(c *fiber.Ctx) error {
	var user model.User
	var err error
	user.UserName = c.FormValue("username")
	user.Password = c.FormValue("password")
	//如果未修改密码则不修改
	if len(user.Password) > 0 && user.Password[0:4] == "$2a$" {
		fmt.Println("not change pwd")
		user.Password = ""
	}

	fmt.Println(user.Password)

	role := c.FormValue("role")
	if role == "" {
		role = "0"
	}
	user.Role, err = strconv.Atoi(role)
	if err != nil {
		return c.JSON(model.OperationResp{Code: 400, Msg: err.Error()})
	}

	//检查id是否传入
	if user.UserName == "" {
		return c.JSON(model.OperationResp{Code: 400, Msg: "username is required"})
	}

	//检查用户权限等级
	if err := method.CheckUserRole(method.GetUserFromToken(c).Role, config.SuperAdmin_ROLE_LEVEL); err != nil {
		return c.JSON(model.OperationResp{Code: 400, Msg: err.Error()})
	}

	//修改用户
	if err := method.UpdateUser(user); err != nil {
		return c.JSON(model.OperationResp{Code: 400, Msg: err.Error()})
	}
	return c.JSON(model.OperationResp{Code: 200, Msg: "ok"})
}

// DeleteUser godoc
//
//	@Summary		删除用户🦸‍♂️
//	@Description	删除用户，本接口对管理员权限的要求为等级7。仅能删除权限比自己低的管理员，普通用户可以被随意删除。
//	@Tags			用户
//	@Accept			json
//	@Produce		json
//	@Param			id	query		string	true	"用户id"
//	@Success		200	{object}	model.OperationResp
//	@Failure		400	{object}	model.OperationResp
//	@Security		ApiKeyAuth
//	@Router			/user/delete [delete]
func DeleteUserHandler(c *fiber.Ctx) error {
	var user model.User
	user.ID = c.FormValue("id")

	if user.ID == "" {
		return c.JSON(model.OperationResp{Code: 400, Msg: "id is required"})
	}

	userRole, err := method.GetUserRole(user.ID)
	if err != nil {
		return c.JSON(model.OperationResp{Code: 400, Msg: err.Error()})
	}

	fmt.Println(method.GetUserFromToken(c).Role, userRole)

	//检查用户权限等级
	if err := method.CheckUserRole(method.GetUserFromToken(c).Role, userRole); err != nil {
		return c.JSON(model.OperationResp{Code: 400, Msg: err.Error()})
	}

	if err := method.DeleteUser(user); err != nil {
		return c.JSON(model.OperationResp{Code: 400, Msg: err.Error()})
	}
	return c.JSON(model.OperationResp{Code: 200, Msg: "ok"})
}

// CheckUserExistHandler godoc
//
//	@Summary		检查用户名是否已经存在
//	@Description	通过用户名检查用户名是否已经存在
//	@Tags			用户
//	@Accept			json
//	@Produce		json
//	@Param			username	query		string	true	"UserName"
//	@Success		200			{object}	model.OperationResp
//	@Failure		400			{object}	model.OperationResp
//	@Security		ApiKeyAuth
//	@Router			/user/exist [get]
func CheckUserExistHandler(c *fiber.Ctx) error {
	UserName := c.FormValue("username")

	if UserName == "" {
		return c.JSON(model.OperationResp{Code: 400, Msg: "username is required"})
	}

	err := method.CheckUserExist(UserName)
	if err != nil {
		return c.JSON(model.OperationResp{Code: 400, Msg: err.Error()})
	}
	return c.JSON(model.OperationResp{Code: 200, Msg: "user exist"})
}
