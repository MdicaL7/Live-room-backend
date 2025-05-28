package handler

import (
	"live_room/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
}

func (user *User) Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	u := &service.User{}
	res, err := u.Login(username, password)
	if err != nil {
		if err.Error() == "密码错误" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  err.Error(),
				"data": nil,
			})
			return
		} else if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"code": 404,
				"msg":  "用户不存在",
				"data": nil,
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": 500,
				"msg":  "服务端异常",
				"data": nil,
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "登录成功",
		"data": gin.H{
			"username": res.Username,
		},
	})
}
