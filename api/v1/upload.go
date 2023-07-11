package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/ginblog/model"
	"github.com/ginblog/utils"
	"github.com/ginblog/utils/errmsg"
	"net/http"
)

func UpLoad(c *gin.Context) {
	Options := utils.Options
	file, fileHeader, _ := c.Request.FormFile("file")
	dst := "./" + "storage/" + fileHeader.Filename
	fileSize := fileHeader.Size

	if Options != "disk" {
		url, code := model.UpLoadFile(file, fileSize)

		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
			"url":     url,
		})
	} else {
		c.SaveUploadedFile(fileHeader, dst)
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
			"url":     dst,
		})
	}

}
