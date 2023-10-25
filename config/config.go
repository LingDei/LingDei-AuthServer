package config

var (
	// 权限控制
	UserROLE_LEVEL        = 0
	Admin_ROLE_LEVEL      = 7
	SuperAdmin_ROLE_LEVEL = 9

	//私钥路径
	PrivateKeyPath = "./config/rsa_private_key.pem"

	// DSN
	DSN = "lingdei:lingdei666666@tcp(127.0.0.1:3306)/LingDei?charset=utf8mb4&parseTime=True&loc=Local"

	// QiNiu
	QINIU_ACCESS_KEY = "R1weO56kS2kZK6BWEelAGaEXirAXKPAOjZ_OoYqk"
	QINIU_SECRET_KEY = "6igpxhFmCVjjWALG8xfxi_MfSHQKkXcHlBS31wgA"
	QINIU_BUCKET     = "lingdei"
)
