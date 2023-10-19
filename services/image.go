package services

import (
	"context"
	"gin-minio/models"
	"gin-minio/utils"
	"github.com/minio/minio-go/v7"
	"log"
	"mime/multipart"
)

var ImageService = ImageServiceType{}

type ImageServiceType struct{}

func (ImageServiceType) Create(bucketName string, folderName string, url string) error {
	image := models.Image{
		BucketName: bucketName,
		FolderName: folderName,
		Url:        url,
	}
	res := utils.DB.Create(&image)
	if res.Error == nil && res.RowsAffected == 1 {
		return nil
	}
	return res.Error
}

func (ImageServiceType) Upload(objectName string, bucketName string, file *multipart.FileHeader, contentType string) error {
	// 创建存储桶,如果有就跳过
	exists, err := utils.MinioClient.BucketExists(context.TODO(), bucketName)
	if err == nil && exists {
		log.Printf("Bucket:%s is already exist\n", bucketName)
	} else {
		err := utils.MinioClient.MakeBucket(context.TODO(), bucketName, minio.MakeBucketOptions{})
		if err != nil {
			return err
		}
	}

	src, openFileErr := file.Open()
	if openFileErr != nil {
		return openFileErr
	}
	defer src.Close()

	_, PutObjectErr := utils.MinioClient.PutObject(context.TODO(), bucketName, objectName, src, -1, minio.PutObjectOptions{ContentType: contentType})
	if PutObjectErr != nil {
		return PutObjectErr
	}
	return nil
}
