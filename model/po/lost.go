package po

import (
	"strings"
	"time"

	"github.com/airdb/sailor/dbutils"
	"github.com/jinzhu/gorm"
)

const (
	keywordsLen1 = 1
	keywordsLen2 = 2
	keywordsLen3 = 3
)

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

func ListLost() []*Lost {
	var losts []*Lost

	pagesize := 10
	dbutils.DefaultDB().Debug().Limit(pagesize).Find(&losts)

	return losts
}

func QueryBBSByKeywords(keyword string) (articles []*Lost) {
	keys := strings.Split(keyword, " ")
	pagesize := 5

	switch len(keys) {
	case keywordsLen3:
		keys[0] = "%" + keys[0] + "%"
		keys[1] = "%" + keys[1] + "%"
		keys[2] = "%" + keys[2] + "%"
		dbutils.DefaultDB().Where(
			"subject like ? and subject like ? and subject like ? ", keys[0], keys[1], keys[2],
		).Select("subject, data_from").Order("missed_at desc").Limit(pagesize).Find(&articles)
	case keywordsLen2:
		keys[0] = "%" + keys[0] + "%"
		keys[1] = "%" + keys[1] + "%"
		dbutils.DefaultDB().Where(
			"subject like ? and subject like ? ", keys[0], keys[1],
		).Select("subject, data_from").Order("missed_at desc").Limit(pagesize).Find(&articles)
	case keywordsLen1:
		keys[0] = "%" + keys[0] + "%"
		dbutils.DefaultDB().Debug().Where(
			"subject like ?", keys[0],
		).Select("subject, data_from").Order("missed_at desc").Limit(pagesize).Find(&articles)
	}

	return
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
