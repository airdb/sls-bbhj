package po

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Babyinfo struct {
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

/*
//func GetAllBabyinfo() (data []Babyinfo, err error) {
//	pagesize := 10
//	conn.Limit(pagesize).Find(&data)
//	return
//}

func GetAllBabyinfo(category, page, pageSize int) (data []Babyinfo, err error) {
	conn.Limit(pageSize).Offset(page * pageSize).Find(&data)
	return
}

func GetBabyinfoById(id int) (data Babyinfo, err error) {
	conn.Find(&data, "id = ?", id)
	return
}

func GetAllBabyinfoByCondition(nickname string) (data []Babyinfo, err error) {
	pageSize := 10
	conn.Where("nickname  LIKE ?", "%"+nickname+"%").Limit(pageSize).Find(&data)
	return
}
*/
