package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/ginblog/model"
	"github.com/ginblog/utils/errmsg"
	"net/http"
)

// 添加播客
func AddPodcast(c *gin.Context) {
	var data model.Podcast
	_ = c.ShouldBindJSON(&data)
	code := model.CheckPodcast(data.Title)
	if code == errmsg.SUCCESS {
		code = model.CreatePodcast(&data)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除播客
func DeletePodcast(c *gin.Context) {
	var data model.Podcast
	_ = c.ShouldBindJSON(&data)
	code := model.DeletePodcast(data.Title)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})

}
