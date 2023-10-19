package controllers

import (
	"fmt"
	"gin-minio/common"
	"gin-minio/services"
	"gin-minio/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UploadMultipleFiles(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["files"]
	bucketName := c.PostForm("bucket_name")
	folderName := c.PostForm("folder_name")
	recordFn := folderName
	if folderName != "" {
		folderName = folderName + "/"
	}
	filePathMinio := make([]string, 0)
	for _, file := range files {
		objectName := utils.MD5(folderName + file.Filename)
		fmt.Println(objectName)
		err := services.ImageService.Upload(objectName, bucketName, file, "")
		if err != nil {
			c.JSON(http.StatusInternalServerError, common.BaseResponse{
				StatusCode: 1,
				StatusMsg:  "UploadMultipleFiles failed:" + err.Error(),
			})
			return
		}
		filePathMinio = append(filePathMinio, "/"+bucketName+"/"+objectName)
		err = services.ImageService.Create(bucketName, recordFn, "/"+bucketName+"/"+objectName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, common.BaseResponse{
				StatusCode: 1,
				StatusMsg:  "UploadMultipleFiles failed:" + err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, common.BaseResponse{
		StatusCode: 0,
		StatusMsg:  "UploadMultipleFiles success",
		Data:       filePathMinio,
	})
	return
}
