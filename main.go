package main

import (
	"gbstore/config"
	"gbstore/controller"
	_ "gbstore/docs"
	"gbstore/model"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func main() {
	config.Setup()
	model.SetProduct()
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/login", controller.Login)

	r.GET("/products/:id", controller.GetProduct)
	r.GET("/products", controller.GetProducts)
	r.POST("/products", controller.AddProducts)
	r.PUT("/products/:id", controller.EditProduct)
	r.DELETE("/products/:id", controller.DelProducts)

	r.GET("/orders", controller.GetOrders)
	r.POST("/orders", controller.AddOrders)

	r.GET("/users/:id/carts", controller.GetUserCart)

	r.POST("/users/:id/carts", controller.AddUserCart)
	r.DELETE("/users/:id/carts/:cartID", controller.DELETEUserCart)

	r.POST("/users/:id/book_order", controller.BookOrders)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
