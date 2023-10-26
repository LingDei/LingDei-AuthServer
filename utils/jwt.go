package utils

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"

	"LingDei-AuthServer/config"
)

func LoadPrivateKey() (*rsa.PrivateKey, error) {
	//打开文件
	file, err := os.Open(config.PrivateKeyPath)
	if err != nil {
		return &rsa.PrivateKey{}, err
	}
	defer file.Close()

	//获取文件内容
	info, err := file.Stat()
	if err != nil {
		return &rsa.PrivateKey{}, err
	}
	buf := make([]byte, info.Size())
	file.Read(buf)

	//pem解码
	block, _ := pem.Decode(buf)
	//X509解码
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	return privateKey, err
}
