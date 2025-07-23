package request

import (
	"time"
)


type OrderRequest struct {
	CarID           int       `json:"car_id" binding:"required"`
	PickupDate      time.Time `json:"pickup_date" binding:"required"`
	DropoffDate     time.Time `json:"dropoff_date" binding:"required"`
	PickupLocation  string    `json:"pickup_location" binding:"required"`
	DropoffLocation string    `json:"dropoff_location" binding:"required"`
}