package schema

import (
	"time"

	"gorm.io/gorm"
)

// Category 分类
type Category struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	Name        string `json:"name"`
	Description string `json:"description"`
}
