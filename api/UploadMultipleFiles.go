package api

import (
	"context"
	"fmt"
	"gin-minio/config"
	"gin-minio/utils"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"log"
	"mime/multipart"
)

func UploadMultipleFiles(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["files"]
	bucketName := c.PostForm("bucket_name")
	fmt.Println("files:", files)
	fmt.Println("bucketName:", bucketName)
	filePathMinio := make([]string, 0)
	i := 0
	for _, file := range files {
		log.Printf("upload file path: %s\n", file.Filename)
		objectName := file.Filename
		err := UploadToMinio(objectName, bucketName, file, "")
		if err != nil {
			c.JSON(1, gin.H{
				"message": "failed",
			})
			return
		}
		filePathMinio = append(filePathMinio, "http://"+config.ENDPOINT+"/"+bucketName+"/"+objectName)
		i = i + 1
	}
	//打印文件下载地址
	fmt.Println(filePathMinio)
}

func UploadToMinio(objectName string, bucketName string, file *multipart.FileHeader, contentType string) error {
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

	// 使用PutObject上传p文件
	_, PutObjectErr := utils.MinioClient.PutObject(context.TODO(), bucketName, objectName, src, -1, minio.PutObjectOptions{ContentType: contentType})
	if PutObjectErr != nil {
		return PutObjectErr
	}
	return nil
}
