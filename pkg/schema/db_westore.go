package schema

import (
	"time"

	"gorm.io/gorm"
)

type Westore struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	Name      string  `json:"name"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	Address   string  `json:"address"`
	Province  string  `json:"provice"`
	City      string  `json:"city"`
	District  string  `json:"district"`
	Phone1    string  `json:"phone1"`
	Phone2    string  `json:"phone2"`
	Phone3    string  `json:"phone3"`
}
