package models

import "time"

type Products struct {
	Id           uint `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
}
