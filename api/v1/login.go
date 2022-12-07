package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/ginblog/middleware"
	"github.com/ginblog/model"
	"github.com/ginblog/utils/errmsg"
	"net/http"
)

func Login(c *gin.Context) {
	var data model.User
	var code int
	var token string
	c.ShouldBindJSON(&data)
	code = model.CheckLogin(data.Username, data.Password)
	if code == errmsg.SUCCESS {
		token, code = middleware.GererateToken(data.Username)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"token":   token,
	})
}
