package method

import (
	"LingDei_AuthServer/model"
	"LingDei_AuthServer/utils"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// AddUser 添加用户
func AddUser(user model.User) error {
	db, err := getDB()
	sqlDB, _ := db.DB() //结束后关闭 DB
	defer sqlDB.Close()
	if err != nil {
		return err
	}

	if err = validate.Struct(user); err != nil {
		return err
	}

	//密码Bcrypt加密
	if user.Password, err = utils.GetPwdBcryptHash([]byte(user.Password)); err != nil {
		return err
	}

	if result := db.Create(&user); result.Error != nil {
		return result.Error
	}

	//创建Profile
	var EmptyProfile = model.Profile{ID: user.ID, NickName: "用户名称", AvatarURL: "https://qingting-bucket.oss-cn-shanghai.aliyuncs.com/avatars/default.png"}
	if err := AddProfile(EmptyProfile); err != nil {
		return err
	}

	return nil
}

// UpdateUser 更新用户
func UpdateUser(user model.User) error {
	db, err := getDB()
	sqlDB, _ := db.DB() //结束后关闭 DB
	defer sqlDB.Close()
	if err != nil {
		return err
	}

	//密码Bcrypt加密
	if user.Password != "" {
		if user.Password, err = utils.GetPwdBcryptHash([]byte(user.Password)); err != nil {
			return err
		}
	}

	// 根据 `struct` 更新属性，只会更新非零值的字段
	if result := db.Model(&user).Where("user_name = ?", user.UserName).Updates(user); result.Error != nil {
		return result.Error
	}

	return nil
}

// GetUser 获取用户
func GetUser(username string) (model.User, error) {
	db, err := getDB()
	sqlDB, _ := db.DB() //结束后关闭 DB
	defer sqlDB.Close()
	if err != nil {
		return model.User{}, err
	}

	var users []model.User
	if result := db.Find(&users, "user_name = ?", username); result.Error != nil {
		return model.User{}, result.Error
	}

	return users[0], nil
}

// GetUserByID 通过id获取用户
func GetUserByID(id string) (model.User, error) {
	db, err := getDB()
	sqlDB, _ := db.DB() //结束后关闭 DB
	defer sqlDB.Close()
	if err != nil {
		return model.User{}, err
	}

	var users []model.User
	if result := db.Find(&users, "id = ?", id); result.Error != nil {
		return model.User{}, result.Error
	}

	fmt.Println(users)

	return users[0], nil
}

// GetUserRole 获取用户权限等级
func GetUserRole(id string) (int, error) {
	db, err := getDB()
	sqlDB, _ := db.DB() //结束后关闭 DB
	defer sqlDB.Close()
	if err != nil {
		return 0, err
	}

	var users []model.User
	if result := db.Find(&users, "id = ?", id); result.Error != nil {
		return 0, result.Error
	}

	return users[0].Role, nil
}

// GetUserID 获取用户id
func GetUserID(username string) (string, error) {
	db, err := getDB()
	sqlDB, _ := db.DB() //结束后关闭 DB
	defer sqlDB.Close()
	if err != nil {
		return "", err
	}

	var users []model.User
	if result := db.Find(&users, "user_name = ?", username); result.Error != nil {
		return "", result.Error
	}

	return users[0].ID, nil
}

// CheckUserExist 检查用户是否存在，存在不报错，不存在报错
func CheckUserExist(username string) error {
	var users []model.User
	db, err := getDB()
	sqlDB, _ := db.DB() //结束后关闭 DB
	defer sqlDB.Close()
	if err != nil {
		return err
	}

	if result := db.Find(&users, "user_name = ?", username); result.Error != nil {
		return result.Error
	}

	if len(users) == 0 {
		return errors.New("user not exist")
	}

	return nil
}

// CheckPwdValid 检查用户密码是否正确
func CheckPwdValid(username, password string) error {
	var users []model.User
	db, err := getDB()
	sqlDB, _ := db.DB() //结束后关闭 DB
	defer sqlDB.Close()
	if err != nil {
		return err
	}

	if result := db.Find(&users, "user_name = ?", username); result.Error != nil {
		return result.Error
	}

	if len(users) == 0 {
		return errors.New("user not exist")
	}

	if utils.ValidatePwdBcryptHash(users[0].Password, []byte(password)) != nil {
		return errors.New("password is incorrect")
	}

	return nil
}

// GetUsers 获取用户列表
func GetUsers(name string) ([]model.User, error) {
	var users []model.User
	db, err := getDB()
	sqlDB, _ := db.DB() //结束后关闭 DB
	defer sqlDB.Close()
	if err != nil {
		return users, err
	}

	if result := db.Find(&users, "user_name LIKE ?", "%"+name+"%"); result.Error != nil {
		return users, err
	}

	return users, nil
}

// DeleteUser 删除用户
func DeleteUser(user model.User) error {
	db, err := getDB()
	sqlDB, _ := db.DB() //结束后关闭 DB
	defer sqlDB.Close()
	if err != nil {
		return err
	}

	result := db.Where("id = ?", user.ID).Delete(&model.User{})

	if result.Error != nil {
		return result.Error
	}

	return nil
}
