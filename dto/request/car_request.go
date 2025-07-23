package request

type CarRequest struct {
	Name      string  `json:"name" binding:"required"`
	DayRate   float64 `json:"day_rate" binding:"required"`
	MonthRate float64 `json:"month_rate" binding:"required"`
	Image     string  `json:"image" binding:"required"`
}
