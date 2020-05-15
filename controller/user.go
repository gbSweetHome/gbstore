package controller

import (
	"fmt"
	"gbstore/model"
	"gbstore/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	appG := utils.Gin{C: c}
	name := c.PostForm("name")
	avatar := c.PostForm("avatar")
	fmt.Println(name, avatar)
	isExist, err := model.Exist(name)
	if err != nil {
		appG.Response(http.StatusOK, utils.ERROR, nil)
		return
	}
	if !isExist {
		//创建新用户
		err := model.AddUser(name, avatar)
		if err != nil {
			appG.Response(http.StatusOK, utils.ERROR, nil)
			return
		}
	}
	//返回用户信息
	user, err := model.GetUser(name)
	if err != nil {
		appG.Response(http.StatusOK, utils.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, utils.SUCCESS, user)
}
