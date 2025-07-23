package response

import (
	"challenge-interview/entity"
	"time"
)

type OrderResponse struct {
	ID              int       `json:"id"`
	CarID           int       `json:"car_id"`
	OrderDate       string `json:"order_date"`
	PickupDate      string `json:"pickup_date"`
	DropoffDate     string `json:"dropoff_date"`
	PickupLocation  string    `json:"pickup_location"`
	DropoffLocation string    `json:"dropoff_location"`
	CreatedAt       string    `json:"created_at"`
	UpdatedAt       string    `json:"updated_at"`
}

func MapOrderToResponse(order *entity.Order) *OrderResponse {
	return &OrderResponse{
		ID:              order.ID,
		CarID:           order.CarID,
		OrderDate:       order.OrderDate.Format(time.RFC3339),
		PickupDate:      order.PickupDate.Format(time.DateOnly),
		DropoffDate:     order.DropoffDate.Format(time.DateOnly),
		PickupLocation:  order.PickupLocation,
		DropoffLocation: order.DropoffLocation,
		CreatedAt:       order.CreatedAt.Format(time.RFC3339),
		UpdatedAt:       order.UpdatedAt.Format(time.RFC3339),
	}
}