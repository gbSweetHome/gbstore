package controller

import (
	"fmt"
	"gbstore/model"
	"gbstore/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetUserCart(c *gin.Context) {
	appG := utils.Gin{C: c}
	userID := c.Param("id")
	intUserID, err := strconv.Atoi(userID)
	carts, err := model.GetUserCart(uint(intUserID))
	if err != nil {
		appG.Response(http.StatusOK, utils.ERROR, nil)
		return
	}
	fmt.Println(carts)
	appG.Response(http.StatusOK, utils.SUCCESS, carts)
}

func AddUserCart(c *gin.Context) {
	appG := utils.Gin{C: c}
	productID := c.PostForm("productID")
	userID := c.Param("id")
	intUserID, err := strconv.Atoi(userID)
	intProductID, err := strconv.Atoi(productID)
	err = model.AddUserCart(uint(intUserID), uint(intProductID))
	if err != nil {
		appG.Response(http.StatusOK, utils.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, utils.SUCCESS, nil)
}
func DELETEUserCart(c *gin.Context) {
	appG := utils.Gin{C: c}
	userID := c.Param("id")
	cartID := c.Param("cartID")
	intUserID, err := strconv.Atoi(userID)
	intCartID, err := strconv.Atoi(cartID)
	err = model.DeleteUserCart(uint(intUserID), []uint{uint(intCartID)})
	if err != nil {
		appG.Response(http.StatusOK, utils.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, utils.SUCCESS, nil)
}
