package models

import(
	"time"
)

type Product struct{
	ID uint `gorm:"primarykey" json:"id"`
	Name string `gorm:"not null" json:"name"`
	Description string `json:"Description"`
	Price float64	`jsonL:"not null" json:"price"`
	stock int 	`gorm:"not null" json:"stock"`
	ImageURL	string	`json:"image_url"`
	CreatedAt	time.Time `json:"created_at"`
	UpdatedAt	time.Time `json:"updated_at"`
}