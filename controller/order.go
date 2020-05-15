package controller

import (
	"gbstore/model"
	"gbstore/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
	"strings"
)

type orderResponse struct {
	gorm.Model
	Id       uint
	UserId   uint
	Name     string
	ImageUrl string
	Price    float64
	State    string
	CreateAt string
}

func AddOrders(c *gin.Context) {
	appG := utils.Gin{C: c}
	name := c.PostForm("name")
	price, err := strconv.ParseFloat(c.PostForm("price"), 64)
	if err != nil {
		appG.Response(http.StatusInternalServerError, utils.ERROR, err)
		return
	}
	imageUrl := c.PostForm("imageUrl")
	userId, err := strconv.Atoi(c.PostForm("userId"))
	if err != nil {
		appG.Response(http.StatusInternalServerError, utils.ERROR, err)
		return
	}
	id, err := model.AddOrder(name, price, imageUrl, uint(userId), "0")
	if err != nil {
		appG.Response(http.StatusInternalServerError, utils.ERROR, err)
		return
	}
	orderResponse := &orderResponse{Id: id}
	appG.Response(http.StatusOK, utils.SUCCESS, orderResponse)
}

func GetOrders(c *gin.Context) {
	type data struct {
		Orders interface{}
		Total  int
	}
	appG := utils.Gin{C: c}
	pageNum, err := strconv.Atoi(c.Query("pageNum"))
	if err != nil {
		appG.Response(http.StatusInternalServerError, utils.ERROR, err)
		return
	}
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
	if err != nil {
		appG.Response(http.StatusInternalServerError, utils.ERROR, err)
		return
	}
	userId, err := strconv.Atoi(c.Query("userId"))
	if err != nil {
		appG.Response(http.StatusInternalServerError, utils.ERROR, err)
		return
	}
	orders, total, err := model.GetOrders(pageNum, pageSize, uint(userId))
	orderResponses := []orderResponse{}
	for _, order := range orders {
		formatTime := order.CreatedAt.Format("2006-01-02 15:04:05")
		orderResponse := orderResponse{Id: order.ID, Name: order.Name, ImageUrl: order.ImageUrl, Price: order.Price, State: order.State, CreateAt: formatTime}
		orderResponses = append(orderResponses, orderResponse)
	}
	orderResponse := &data{orderResponses, total}
	appG.Response(http.StatusOK, utils.SUCCESS, orderResponse)
}

func BookOrders(c *gin.Context) {
	appG := utils.Gin{C: c}
	userID := c.Param("id")
	productIDs := c.PostFormArray("productID[]")
	cartIDs := c.PostFormArray("cartsID[]")
	intUserID, err := strconv.Atoi(userID)
	var intProductIDs []uint
	for _, item := range productIDs {
		intItem, _ := strconv.Atoi(item)
		intProductIDs = append(intProductIDs, uint(intItem))
	}
	var intCardIDs []uint
	for _, item := range cartIDs {
		intItem, _ := strconv.Atoi(item)
		intCardIDs = append(intCardIDs, uint(intItem))
	}
	prods, err := model.GetProductsByIDS(intProductIDs)
	if err != nil {
		appG.Response(http.StatusOK, utils.ERROR, nil)
		return
	}
	var name []string
	var price float64
	var imageUrl string
	for idx, item := range prods {
		if imageUrl == "" {
			imageUrl = item.ImageUrl
		}
		if idx < 3 {
			name = append(name, item.Name)
		}
		price += item.Price
	}
	totalName := strings.Join(name, "、")
	if len(prods) > 3 {
		totalName = totalName + "等"
	}
	_, err = model.AddOrder(totalName, price, imageUrl, uint(intUserID), "0")
	if err != nil {
		appG.Response(http.StatusOK, utils.ERROR, nil)
	}
	err = model.DeleteUserCart(uint(intUserID), intCardIDs)
	if err != nil {
		appG.Response(http.StatusOK, utils.ERROR, nil)
	}
	appG.Response(http.StatusOK, utils.SUCCESS, nil)
}
