package v1

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/ginblog/model"
	"github.com/ginblog/utils"
	"github.com/ginblog/utils/errmsg"
	"net/http"
	"time"
)

var StandTime = []byte(time.Now().Format("20060102150405"))
var Base64Encode = base64.StdEncoding.EncodeToString(StandTime)

func UpLoad(c *gin.Context) {
	Options := utils.Options
	file, fileHeader, _ := c.Request.FormFile("file")
	// 拼接本地路径
	dst := "./" + "storage/" + time.Now().Format("20060102150405") + Base64Encode + ".jpg"
	fileSize := fileHeader.Size

	if Options != "disk" {
		url, code := model.UpLoadFile(file, fileSize)

		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
			"url":     url,
		})
	} else {
		_ = c.SaveUploadedFile(fileHeader, dst)

		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": errmsg.GetErrMsg(http.StatusOK),
			"url":     dst,
		})
	}

}
