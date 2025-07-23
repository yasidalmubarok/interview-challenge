package entity

import "time"

type Car struct {
	ID        int       `gorm:"primaryKey"`
	Name	  string    `gorm:"not null"`
	DayRate   float64   `gorm:"not null"`
	MonthRate float64   `gorm:"not null"`
	Image     string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}
