package response

import (
	"challenge-interview/entity"
	"time"
)

type CarResponse struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	DayRate   float64 `json:"day_rate"`
	MonthRate float64 `json:"month_rate"`
	Image     string  `json:"image"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

func MapCarToResponse(car *entity.Car) *CarResponse {
	return &CarResponse{
		ID:        car.ID,
		Name:      car.Name,
		DayRate:   car.DayRate,
		MonthRate: car.MonthRate,
		Image:     car.Image,
		CreatedAt: car.CreatedAt.Format(time.RFC3339),
		UpdatedAt: car.UpdatedAt.Format(time.RFC3339),
	}
}
