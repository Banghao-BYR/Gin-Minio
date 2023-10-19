package utils

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

var MinioClient *minio.Client

func InitMinio() {
	// 初使化 minio client对象
	var err error
	MinioClient, err = minio.New(Config.Minio.URL, &minio.Options{Creds: credentials.NewStaticV4(Config.Minio.Username, Config.Minio.Password, "")})
	if err != nil {
		log.Println("Init Minio err:", err.Error())
		return
	}
	log.Println("Init Minio...")
}
