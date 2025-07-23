package main

import (
	"challenge-interview/config"
	"challenge-interview/handler"
	"challenge-interview/repository"
	"challenge-interview/service"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	config.LoadEnv()

	db := config.InitDB()
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to get database instance: " + err.Error())
	}
	defer sqlDB.Close()

	carRepo := repository.NewCarRepository(db)
	carService := service.NewCarService(carRepo)
	carHandler := handler.NewCarHandler(carService)

	orderRepo := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepo, carRepo)
	orderHandler := handler.NewOrderHandler(orderService)

	r := gin.Default()

	api := r.Group("/api/v1")

	api.POST("/car", carHandler.Create)
	api.GET("/cars", carHandler.GetAll)
	api.GET("/car/:id", carHandler.GetByID)
	api.PATCH("/car/:id", carHandler.Update)
	api.DELETE("/car/:id", carHandler.Delete)

	api.POST("/order", orderHandler.Create)
	api.GET("/orders", orderHandler.GetAll)
	api.GET("/order/:id", orderHandler.GetByID)
	api.PATCH("/order/:id", orderHandler.Update)
	api.DELETE("/order/:id", orderHandler.Delete)

	r.Run(":8080")
}
