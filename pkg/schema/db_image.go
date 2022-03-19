package schema

import (
	"time"

	"gorm.io/gorm"
)

type Image struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	OpenID  string `json:"open_id" gorm:"column:openid"`
	UnionID string `json:"union_id" gorm:"column:unionid"`
	UUID    string `json:"uuid"`
	URL     string `json:"url"`
	Status  int    `json:"status"`
}
