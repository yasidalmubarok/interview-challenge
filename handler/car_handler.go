package handler

import (
	"challenge-interview/dto/request"
	"challenge-interview/helper"
	"challenge-interview/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CarHandler struct {
	carService service.CarService
}

func NewCarHandler(carService service.CarService) *CarHandler {
	return &CarHandler{
		carService: carService,
	}
}

func (h *CarHandler) Create(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		response := helper.APIResponse(err.Error(), "error", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	imagePath, err := helper.UploadAndCompressImage(file, 500)
	if err != nil {
		response := helper.APIResponse(err.Error(), "error", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	dayRate, _ := strconv.ParseFloat(c.PostForm("day_rate"), 64)
	monthRate, _ := strconv.ParseFloat(c.PostForm("month_rate"), 64)

	carRequest := &request.CarRequest{}
	carRequest.Name = c.PostForm("name")
	carRequest.DayRate = dayRate
	carRequest.MonthRate = monthRate
	carRequest.Image = imagePath

	newCar, err := h.carService.Create(carRequest)
	if err != nil {
		apiResponse := helper.APIResponse(err.Error(), "error", http.StatusInternalServerError, nil)
		c.JSON(http.StatusInternalServerError, apiResponse)
		return
	}

	apiResponse := helper.APIResponse("Car created successfully", "success", http.StatusCreated, newCar)
	c.JSON(http.StatusCreated, apiResponse)
}

func (h *CarHandler) GetAll(c *gin.Context) {
	cars, err := h.carService.GetAll()
	if err != nil {
		apiResponse := helper.APIResponse(err.Error(), "error", http.StatusInternalServerError, nil)
		c.JSON(http.StatusInternalServerError, apiResponse)
		return
	}

	apiResponse := helper.APIResponse("Cars retrieved successfully", "success", http.StatusOK, cars)
	c.JSON(http.StatusOK, apiResponse)
}

func (h *CarHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		apiResponse := helper.APIResponse("Invalid car ID", "error", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, apiResponse)
		return
	}

	car, err := h.carService.GetByID(id)
	if err != nil {
		apiResponse := helper.APIResponse(err.Error(), "error", http.StatusInternalServerError, nil)
		c.JSON(http.StatusInternalServerError, apiResponse)
		return
	}

	apiResponse := helper.APIResponse("Car retrieved successfully", "success", http.StatusOK, car)
	c.JSON(http.StatusOK, apiResponse)
}

func (h *CarHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		apiResponse := helper.APIResponse("Invalid car ID", "error", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, apiResponse)
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		response := helper.APIResponse(err.Error(), "error", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	imagePath, err := helper.UploadAndCompressImage(file, 500)
	if err != nil {
		response := helper.APIResponse(err.Error(), "error", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	dayRate, _ := strconv.ParseFloat(c.PostForm("day_rate"), 64)
	monthRate, _ := strconv.ParseFloat(c.PostForm("month_rate"), 64)

	carRequest := &request.CarRequest{}
	carRequest.Name = c.PostForm("name")
	carRequest.DayRate = dayRate
	carRequest.MonthRate = monthRate
	carRequest.Image = imagePath

	updatedCar, err := h.carService.Update(carRequest, id)
	if err != nil {
		apiResponse := helper.APIResponse(err.Error(), "error", http.StatusInternalServerError, nil)
		c.JSON(http.StatusInternalServerError, apiResponse)
		return
	}

	apiResponse := helper.APIResponse("Car updated successfully", "success", http.StatusOK, updatedCar)
	c.JSON(http.StatusOK, apiResponse)
}

func (h *CarHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		apiResponse := helper.APIResponse("Invalid car ID", "error", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, apiResponse)
		return
	}

	deleted, err := h.carService.Delete(id)
	if err != nil {
		apiResponse := helper.APIResponse(err.Error(), "error", http.StatusInternalServerError, nil)
		c.JSON(http.StatusInternalServerError, apiResponse)
		return
	}

	if !deleted {
		apiResponse := helper.APIResponse("Car not found", "error", http.StatusNotFound, nil)
		c.JSON(http.StatusNotFound, apiResponse)
		return
	}

	apiResponse := helper.APIResponse("Car deleted successfully", "success", http.StatusOK, nil)
	c.JSON(http.StatusOK, apiResponse)
}