package store

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type lost struct {
	db *gorm.DB
}

type Lost struct {
	gorm.Model
	// ID        string `gorm:"primary_key"`
	//  Timestamp int64
	// CreatedAt    time.Time `sql:"DEFAULT:current_timestamp"`
	UUID      string
	AvatarURL string
	Nickname  string
	// 0: unknown,  1: male   2: female
	Gender          uint
	Title           string
	Subject         string
	Characters      string
	Details         string
	DataFrom        string
	BirthedProvince string
	BirthedCity     string
	BirthedCountry  string
	BirthedAddress  string
	BirthedAt       time.Time `gorm:"type:datetime"`

	MissedCountry  string
	MissedProvince string
	MissedCity     string
	MissedAddress  string
	MissedAt       time.Time `gorm:"column:missed_at;type:datetime"`
	Handler        string
	Babyid         string
	Category       string
	Height         string
	SyncStatus     int `gorm:"column:syncstatus;default:0"`
}

func newLost(ds *datastore) *lost {
	return &lost{db: ds.db}
}

// Create creates a new talk item.
func (d *lost) List() error {
	var info []*Lost

	r := d.db.Find(&info).Limit(10)
	log.Println("len: ", len(info))

	return r.Error
}
