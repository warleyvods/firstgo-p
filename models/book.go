package models

import "time"

type Book struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	MediumPrice float32   `json:"medium_price"`
	Author      string    `json:"author"`
	ImageURL    string    `json:"image_url"`
	CreatedAt   time.Time `json:"created"`
	UpdatedAt   time.Time `json:"updated"`
	DeletedAt   time.Time `gorm:"index" json:"deleted"`
}
