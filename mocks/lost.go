package mocks

import (
	"time"

	"github.com/airdb/mina-api/model/po"
	"github.com/jinzhu/gorm"
)

var Lost1 = &po.Lost{
	Model: gorm.Model{
		ID: 1, // nolint:gomnd
	},

	UUID:            "300ad9b5-afc8-41ff-972d-163bfed47b91",
	AvatarURL:       "http://www.baobeihuijia.com/photo/water/water_15334.jpg",
	Nickname:        "张作岭",
	Gender:          0,
	Title:           "[男孩] 1983年腊月初八出生于山西省怀仁县，被送养至河南新乡的张作岭寻亲 15334",
	Subject:         "",
	Characters:      "",
	Details:         "",
	DataFrom:        "",
	BirthedProvince: "",
	BirthedCity:     "",
	BirthedCountry:  "",
	BirthedAddress:  "",
	BirthedAt:       time.Time{},
	MissedCountry:   "",
	MissedProvince:  "",
	MissedCity:      "",
	MissedAddress:   "",
	MissedAt:        time.Time{},
	Handler:         "",
	Babyid:          "",
	Category:        "",
	Height:          "",
	SyncStatus:      0,
}
