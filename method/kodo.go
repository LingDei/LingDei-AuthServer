package method

import (
	"LingDei_AuthServer/config"
	"context"
	"fmt"
	"mime/multipart"
	"strings"

	"github.com/google/uuid"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

func UploadAvator(file *multipart.FileHeader) (string, error) {
	//获取文件内容
	fileBytes, err := file.Open()
	if err != nil {
		return "", err
	}

	//获取文件扩展名
	fileExtension := strings.Split(file.Filename, ".")[1]

	mac := qbox.NewMac(config.QINIU_ACCESS_KEY, config.QINIU_SECRET_KEY)

	putPolicy := storage.PutPolicy{
		Scope: config.QINIU_BUCKET,
	}
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Region = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "avatar",
		},
	}

	randomName := strings.Replace(uuid.NewString(), "-", "", -1) + "." + fileExtension
	if err := formUploader.Put(context.Background(), &ret, upToken, "avatars/"+randomName, fileBytes, file.Size, &putExtra); err != nil {
		fmt.Println(err)
		return "", err
	}

	fmt.Println(ret)

	return "avatars/" + randomName, nil
}
