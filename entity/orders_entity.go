package entity

import "time"

type Order struct {
	ID              int       `gorm:"primaryKey"`
	CarID           int       `gorm:"not null"`
	OrderDate       time.Time `gorm:"not null"`
	PickupDate      time.Time `gorm:"not null"`
	DropoffDate     time.Time `gorm:"not null"`
	PickupLocation  string    `gorm:"not null"`
	DropoffLocation string    `gorm:"not null"`
	CreatedAt       time.Time `gorm:"not null"`
	UpdatedAt       time.Time `gorm:"not null"`
}
