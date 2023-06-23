package utils

import (
	"gin-minio/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

var MinioClient *minio.Client

func InitMinio() {
	// 初使化 minio client对象
	var err error
	MinioClient, err = minio.New(config.ENDPOINT, &minio.Options{Creds: credentials.NewStaticV4(config.ACCESSKEYID, config.SECRETACCESSKEY, "")})
	if err != nil {
		log.Println("Init Minio err:", err)
		return
	}
	log.Println("Init Minio...")
}
