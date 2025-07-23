package handler

import (
	"challenge-interview/dto/request"
	"challenge-interview/helper"
	"challenge-interview/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type orderHandler struct {
	orderService service.OrderService
}

func NewOrderHandler(orderService service.OrderService) *orderHandler {
	return &orderHandler{
		orderService: orderService,
	}
}

func (h *orderHandler) Create(c *gin.Context) {
	payload := &request.OrderRequest{}
	err := c.ShouldBindJSON(payload)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errBindJson := helper.NewStatusUnProcessableEntityError(errors[0])
		response := helper.APIResponse(errBindJson.Error(), "error", errBindJson.StatusCode(), nil)
		c.JSON(errBindJson.StatusCode(), response)
		return
	}

	createdOrder, err2 := h.orderService.Create(payload)
	if err2 != nil {
		response := helper.APIResponse(err2.Error(), "error", err2.StatusCode(), nil)
		c.JSON(err2.StatusCode(), response)
		return
	}

	response := helper.APIResponse("Order created successfully", "success", http.StatusCreated, createdOrder)
	c.JSON(http.StatusCreated, response)
}

func (h *orderHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := helper.APIResponse("Invalid order ID", "error", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	order, err2 := h.orderService.GetByID(id)
	if err2 != nil {
		response := helper.APIResponse(err2.Error(), "error", err2.StatusCode(), nil)
		c.JSON(err2.StatusCode(), response)
		return
	}

	response := helper.APIResponse("Order retrieved successfully", "success", http.StatusOK, order)
	c.JSON(http.StatusOK, response)
}

func (h *orderHandler) GetAll(c *gin.Context) {
	orders, err := h.orderService.GetAll()
	if err != nil {
		response := helper.APIResponse(err.Error(), "error", err.StatusCode(), nil)
		c.JSON(err.StatusCode(), response)
		return
	}

	response := helper.APIResponse("Orders retrieved successfully", "success", http.StatusOK, orders)
	c.JSON(http.StatusOK, response)
}

func (h *orderHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := helper.APIResponse("Invalid order ID", "error", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	payload := &request.OrderRequest{}
	err = c.ShouldBindJSON(payload)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errBindJson := helper.NewStatusUnProcessableEntityError(errors[0])
		response := helper.APIResponse(errBindJson.Error(), "error", errBindJson.StatusCode(), nil)
		c.JSON(errBindJson.StatusCode(), response)
		return
	}

	updatedOrder, err2 := h.orderService.Update(payload, id)
	if err2 != nil {
		response := helper.APIResponse(err2.Error(), "error", err2.StatusCode(), nil)
		c.JSON(err2.StatusCode(), response)
		return
	}

	response := helper.APIResponse("Order updated successfully", "success", http.StatusOK, updatedOrder)
	c.JSON(http.StatusOK, response)
}

func (h *orderHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := helper.APIResponse("Invalid order ID", "error", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	deleted, err2 := h.orderService.Delete(id)
	if err2 != nil {
		response := helper.APIResponse(err2.Error(), "error", err2.StatusCode(), nil)
		c.JSON(err2.StatusCode(), response)
		return
	}

	if !deleted {
		response := helper.APIResponse("Order not found", "error", http.StatusNotFound, nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	response := helper.APIResponse("Order deleted successfully", "success", http.StatusOK, nil)
	c.JSON(http.StatusOK, response)
}
