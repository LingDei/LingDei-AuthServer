package handler

import (
	"LingDei-AuthServer/config"
	"LingDei-AuthServer/method"
	"LingDei-AuthServer/model"

	"github.com/gofiber/fiber/v2"
)

// GetProfileHandler godoc
//
//	@Summary		获取用户Profile
//	@Description	获取用户Profile，用户id通过token取得，不需要单独传参
//	@Tags			用户个人资料
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.ProfileResp
//	@Failure		400	{object}	model.OperationResp
//	@Security		ApiKeyAuth
//	@Router			/profile/get [get]
func GetProfileHandler(c *fiber.Ctx) error {
	id := method.GetUserFromToken(c).ID
	profile, err := method.GetProfile(id)
	if err != nil {
		return c.JSON(model.OperationResp{Code: 400, Msg: err.Error()})
	}
	return c.JSON(model.ProfileResp{Code: 200, Profile: profile})
}

// UpdateProfileHandler godoc
//
//	@Summary		更新用户个人资料
//	@Description	更新用户个人资料，用户id通过token取得，不需要单独传参。若Profile不存在则会自动创建新Profile。若参数为空则代表不对该项进行修改。
//	@Tags			用户个人资料
//	@Accept			json
//	@Produce		json
//	@Param			nickname	query		string	false	"用户昵称"
//	@Param			avatarurl	query		string	false	"头像地址"
//	@Param			email		query		string	false	"邮箱"
//	@Success		200			{object}	model.OperationResp
//	@Failure		400			{object}	model.OperationResp
//	@Security		ApiKeyAuth
//	@Router			/profile/update [post]
func UpdateProfileHandler(c *fiber.Ctx) error {
	var profile model.Profile
	profile.ID = method.GetUserFromToken(c).ID
	profile.NickName = c.FormValue("nickname")
	profile.AvatarURL = c.FormValue("avatarurl")
	profile.Email = c.FormValue("email")

	//修改Profile
	if err := method.UpdateProfile(profile); err != nil {
		return c.JSON(model.OperationResp{Code: 400, Msg: err.Error()})
	}
	return c.JSON(model.OperationResp{Code: 200, Msg: "ok"})
}

// DeleteProfileHandler godoc
//
//	@Summary		删除用户Profile🦸‍♂️
//	@Description	删除用户Profile，本接口对管理员权限的要求为等级7。一般不需要用到本接口。
//	@Tags			用户个人资料
//	@Accept			json
//	@Produce		json
//	@Param			id	query		string	true	"被删除Profile的用户id"
//	@Success		200	{object}	model.OperationResp
//	@Failure		400	{object}	model.OperationResp
//	@Security		ApiKeyAuth
//	@Router			/profile/delete [delete]
func DeleteProfileHandler(c *fiber.Ctx) error {
	var profile model.Profile
	profile.ID = c.FormValue("id")

	if profile.ID == "" {
		return c.JSON(model.OperationResp{Code: 400, Msg: "id is required"})
	}

	//检查用户权限等级
	if err := method.CheckUserRole(method.GetUserFromToken(c).Role, config.SuperAdmin_ROLE_LEVEL); err != nil {
		return c.JSON(model.OperationResp{Code: 400, Msg: err.Error()})
	}

	err := method.DeleteProfile(profile.ID)
	if err != nil {
		return c.JSON(model.OperationResp{Code: 400, Msg: err.Error()})
	}
	return c.JSON(model.OperationResp{Code: 200, Msg: "ok"})
}

// UploadAvatorHandler godoc
//
//	@Summary		上传头像
//	@Description	上传头像
//	@Tags			用户个人资料
//	@Accept			json
//	@Produce		json
//	@Param			file	formData	file	true	"头像文件"
//	@Success		200		{object}	model.AvatorResp
//	@Failure		400		{object}	model.OperationResp
//	@Security		ApiKeyAuth
//	@Router			/profile/avatar [post]
func UploadAvatorHandler(c *fiber.Ctx) error {
	profile := model.Profile{}
	profile.ID = method.GetUserFromToken(c).ID

	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(model.OperationResp{Code: 400, Msg: err.Error()})
	}

	//上传文件
	fileUrl, err := method.UploadAvator(file)
	if err != nil {
		return c.JSON(model.OperationResp{Code: 400, Msg: err.Error()})
	}

	profile.AvatarURL = "https://bucket.lingdei.doyi.online/" + fileUrl

	// 修改Profile
	if err := method.UpdateProfileWithoutCheck(profile); err != nil {
		return c.JSON(model.OperationResp{Code: 400, Msg: err.Error()})
	}

	return c.JSON(model.AvatorResp{Code: 200, Url: fileUrl})
}
