package config

import (
	"fmt"

	"os"

	"github.com/spf13/viper"
)

var (
	// 权限控制
	UserROLE_LEVEL        = 0
	Admin_ROLE_LEVEL      = 7
	SuperAdmin_ROLE_LEVEL = 9

	//私钥路径
	PrivateKeyPath = "./config/rsa_private_key.pem"

	DSN = "lingdei:lingdei123456@tcp(127.0.0.1:3306)/LingDei?charset=utf8mb4&parseTime=True&loc=Local"

	// QiNiu
	QINIU_ACCESS_KEY = "your-access-key"
	QINIU_SECRET_KEY = "your-secret-key"
	QINIU_BUCKET     = "your-bucket"
)

func init() {
	// 检查配置文件是否存在
	if !checkConfigFileIsExist() {
		// 不存在则创建
		createConfigFile()
	}

	// 读取config.yaml
	// 设置配置文件的名字
	viper.SetConfigName("config")
	// 设置配置文件的类型
	viper.SetConfigType("yaml")
	// 添加配置文件的路径，指定 config 目录下寻找
	viper.AddConfigPath("./config")
	// 寻找配置文件并读取
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	// 读取配置文件
	DSN = viper.GetString("mysql.DSN")
	QINIU_ACCESS_KEY = viper.GetString("qiniu.access_key")
	QINIU_SECRET_KEY = viper.GetString("qiniu.secret_key")
	QINIU_BUCKET = viper.GetString("qiniu.bucket")
}

var default_config = `mysql:
  DSN: lingdei:lingdei123456@tcp(127.0.0.1:3306)/LingDei?charset=utf8mb4&parseTime=True&loc=Local

QINIU:
  access_key: your-access-key
  secret_key: your-secret-key
  bucket: your-bucket
`

// 检查配置文件是否存在
func checkConfigFileIsExist() bool {
	_, err := os.Stat("./config/config.yaml")
	return err == nil || os.IsExist(err)
}

// 创建配置文件
func createConfigFile() {
	// 创建文件
	f, err := os.Create("./config/config.yaml")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// 写入配置
	_, err = f.WriteString(default_config)
	if err != nil {
		panic(err)
	}

	// 保存文件
	err = f.Sync()
	if err != nil {
		panic(err)
	}
}
