package controller

import (
	"gbstore/model"
	"gbstore/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type productResponse struct {
	Id          uint
	Name        string
	ImageUrl    string
	Price       float64
	MonthlySale string
	Describe    string
	Rate        *int64
}

// @Summary 新增文章标签
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Router /api/v1/tags [get]
func GetProduct(c *gin.Context) {
	appG := utils.Gin{C: c}
	//fmt.Println(c.Param("id"))
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		appG.Response(http.StatusInternalServerError, utils.ERROR, err)
		return
	}
	product, err := model.GetProduct(uint(id))
	if err != nil {
		appG.Response(http.StatusInternalServerError, utils.ERROR, err)
		return
	}
	productResponse := &productResponse{Id: product.ID, Name: product.Name, ImageUrl: product.ImageUrl, Price: product.Price, MonthlySale: product.MonthlySale, Describe: product.Describe, Rate: product.Rate}
	appG.Response(http.StatusOK, utils.SUCCESS, productResponse)
}

// @Summary 新增文章标签
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Router /api/v1/tags [post]
func AddProducts(c *gin.Context) {
	appG := utils.Gin{C: c}
	name := c.PostForm("name")
	describe := c.PostForm("describe")
	price, err := strconv.ParseFloat(c.PostForm("price"), 64)
	if err != nil {
		appG.Response(http.StatusInternalServerError, utils.ERROR, err)
		return
	}
	imageUrl := c.PostForm("imageUrl")
	id, err := model.AddProduct(name, describe, price, imageUrl)
	if err != nil {
		appG.Response(http.StatusInternalServerError, utils.ERROR, err)
		return
	}
	productResponse := &productResponse{Id: id}
	appG.Response(http.StatusOK, utils.SUCCESS, productResponse)
}
func GetProducts(c *gin.Context) {
	type data struct {
		Products interface{}
		Total    int
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
	products, total, err := model.GetProducts(pageNum, pageSize)
	var productResponses []productResponse
	for _, product := range products {
		productResponse := productResponse{Id: product.ID, Name: product.Name, ImageUrl: product.ImageUrl, Price: product.Price, MonthlySale: product.MonthlySale, Describe: product.Describe, Rate: product.Rate}
		productResponses = append(productResponses, productResponse)
	}
	dataResponse := &data{Products: productResponses, Total: total}
	appG.Response(http.StatusOK, utils.SUCCESS, dataResponse)
}
func EditProduct(c *gin.Context) {

}
func DelProducts(c *gin.Context) {}
