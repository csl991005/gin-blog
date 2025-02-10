package v1

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ginblog/model"
	"github.com/ginblog/utils"
	"github.com/ginblog/utils/errmsg"
)

func UpLoad(c *gin.Context) {
	Options := utils.Options
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  errmsg.ERROR,
			"message": "上传文件失败",
		})
		return
	}

	fileSize := fileHeader.Size

	if Options != "disk" {
		url, code := model.UpLoadFile(file, fileSize)
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
			"url":     url,
		})
		return
	}

	// 确保 storage 目录存在
	storagePath := "./storage"
	if _, err := os.Stat(storagePath); os.IsNotExist(err) {
		_ = os.MkdirAll(storagePath, os.ModePerm)
	}

	// 生成唯一文件名
	fileExt := filepath.Ext(fileHeader.Filename) // 获取原始文件扩展名
	fileName := time.Now().Format("20060102150405") + "_" + fmt.Sprintf("%d", time.Now().UnixNano())

	dst := filepath.Join(storagePath, fileName+fileExt)

	// 保存文件
	err = c.SaveUploadedFile(fileHeader, dst)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  errmsg.ERROR,
			"message": "保存文件失败",
		})
		return
	}

	// 拼接完整 URL，使用配置文件中的 BaseURL
	fullURL := utils.BaseURL + "/storage/" + fileName + fileExt

	// 返回 JSON
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": errmsg.GetErrMsg(http.StatusOK),
		"url":     fullURL,
	})
}
