package schema

import (
	"time"

	"gorm.io/gorm"
)

type Lost struct {
	// gorm.Model
	// ID        string `gorm:"primary_key"`
	//  Timestamp int64
	// CreatedAt    time.Time `sql:"DEFAULT:current_timestamp"`
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	UUID      string `json:"uuid"`
	AvatarURL string `json:"avatar_url"`
	Nickname  string `json:"nickname"`
	// 0: unknown,  1: male   2: female
	Gender          uint      `json:"gender"`
	Title           string    `json:"title"`
	Subject         string    `json:"subject"`
	Characters      string    `json:"characters"`
	Details         string    `json:"details"`
	DataFrom        string    `json:"data_from"`
	BirthedProvince string    `json:"birthed_province"`
	BirthedCity     string    `json:"birthed_city"`
	BirthedCountry  string    `json:"birthed_country"`
	BirthedAddress  string    `json:"birthed_address"`
	BirthedAt       time.Time `gorm:"type:datetime" json:"birthed_at"`

	MissedCountry  string    `json:"missed_country"`
	MissedProvince string    `json:"missed_province"`
	MissedCity     string    `json:"missed_city"`
	MissedAddress  string    `json:"missed_address"`
	MissedAt       time.Time `gorm:"column:missed_at;type:datetime" json:"missed_at"`
	Handler        string    `json:"handler"`
	Babyid         string    `json:"babyid"`
	Category       string    `json:"category"`
	Height         string    `json:"height"`
	SyncStatus     int       `gorm:"column:syncstatus;default:0" json:"sync_status"`
}
