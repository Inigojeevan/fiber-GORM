package models

import "time"

type Order struct {
	Id           uint `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time
	ProductRefer int      `json:"product_id"`
	Product      Products `gorm:"foreignKey:ProductRefer"`
	UserRefer    int      `json:"user_id"`
	User         User     `gorm:"foreignKey:UserRefer"`
}
