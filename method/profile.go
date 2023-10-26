package method

import (
	"LingDei-AuthServer/model"
	"errors"
)

// AddProfile 添加Profile
func AddProfile(profile model.Profile) error {
	db, err := getDB()
	sqlDB, _ := db.DB() //结束后关闭 DB
	defer sqlDB.Close()
	if err != nil {
		return err
	}

	if err = validate.Struct(profile); err != nil {
		return err
	}

	if result := db.Create(&profile); result.Error != nil {
		return result.Error
	}

	return nil
}

// UpdateProfile 修改Profile
func UpdateProfile(profile model.Profile) error {
	db, err := getDB()
	sqlDB, _ := db.DB() //结束后关闭 DB
	defer sqlDB.Close()
	if err != nil {
		return err
	}

	if err := validate.Struct(profile); err != nil {
		return err
	}

	//Profile不存在则新创建
	if CheckProfileExist(profile.ID) != nil {
		AddProfile(profile)
	}

	// 根据 `struct` 更新属性，只会更新非零值的字段
	if result := db.Model(&profile).Where("id = ?", profile.ID).Updates(profile); result.Error != nil {
		return result.Error
	}

	return nil
}

// GetProfile 获取Profile
func GetProfile(id string) (model.Profile, error) {
	db, err := getDB()
	sqlDB, _ := db.DB() //结束后关闭 DB
	defer sqlDB.Close()
	if err != nil {
		return model.Profile{}, err
	}

	var profiles []model.Profile
	if result := db.Find(&profiles, "id = ?", id); result.Error != nil {
		return model.Profile{}, result.Error
	}

	//如果找不到profile记录
	if len(profiles) < 1 {
		return model.Profile{}, errors.New("profile not exist")
	}

	return profiles[0], nil
}

// CheckProfileExist 检查Profile是否存在
func CheckProfileExist(id string) error {
	var profiles []model.Profile
	db, err := getDB()
	sqlDB, _ := db.DB() //结束后关闭 DB
	defer sqlDB.Close()
	if err != nil {
		return err
	}

	if result := db.Find(&profiles, "id = ?", id); result.Error != nil {
		return result.Error
	}

	if len(profiles) == 0 {
		return errors.New("profile not exist")
	}

	return nil
}

// DeleteProfile 删除Profile
func DeleteProfile(id string) error {
	db, err := getDB()
	sqlDB, _ := db.DB() //结束后关闭 DB
	defer sqlDB.Close()
	if err != nil {
		return err
	}

	result := db.Where("id = ?", id).Delete(&model.Profile{})

	if result.Error != nil {
		return result.Error
	}

	return nil
}
