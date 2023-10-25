package config

var (
	//私钥路径
	PrivateKeyPath = "./config/rsa_private_key.pem"

	// DSN
	DSN = "lingdei:lingdei666666@tcp(127.0.0.1:3306)/LingDei?charset=utf8mb4&parseTime=True&loc=Local"

	// QiNiu
	QINIU_ACCESS_KEY = "your access key"
	QINIU_SECRET_KEY = "your secret key"
	QINIU_BUCKET     = "your bucket"
)
