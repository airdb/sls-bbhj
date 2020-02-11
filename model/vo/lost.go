package vo

import (
	"fmt"
	"time"

	"github.com/airdb/mina-api/model/po"
)

type ListLostReq struct {
	Category uint `form:"category"`
	Page     uint `form:"page"`
	PageSize uint `form:"pageSize"`
}

type ListLostResp struct {
}

type QueryListReq struct {
}

type QueryListResp struct {
}

type SearchLostReq struct {
	Keywords string `json:"keywords"`
}

type SearchLostResp struct {
}

type QQRobotQueryReq struct {
	QQ      uint   `form:"qq"`
	Group   uint   `form:"group"`
	Command string `form:"command"`
	Message string `form:"message"`
}

type QQRobotQueryResp string

func QueryBBSByKeyword(keyword string) []*po.Lost {
	return po.QueryBBSByKeywords(keyword)
}

type Lost struct {
	ID        uint       `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

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
	BirthedAt       time.Time `json:"birthed_at"`

	MissedCountry  string    `json:"missed_country"`
	MissedProvince string    `json:"missed_province"`
	MissedCity     string    `json:"missed_city"`
	MissedAddress  string    `json:"missed_address"`
	MissedAt       time.Time `json:"missed_at"`
	Handler        string    `json:"handler"`
	Babyid         string    `json:"babyid"`
	Category       string    `json:"category"`
	Height         string    `json:"Height"`
	SyncStatus     int       `json:"sync_status"`
}

func FromPoLost(lost *po.Lost) *Lost {
	fmt.Println("xxx", lost)

	return &Lost{
		ID:              lost.ID,
		CreatedAt:       lost.CreatedAt,
		UpdatedAt:       lost.UpdatedAt,
		DeletedAt:       lost.DeletedAt,
		UUID:            lost.UUID,
		AvatarURL:       lost.AvatarURL,
		Nickname:        lost.Nickname,
		Gender:          lost.Gender,
		Title:           lost.Title,
		Subject:         lost.Subject,
		Characters:      lost.Characters,
		Details:         lost.Details,
		DataFrom:        lost.DataFrom,
		BirthedProvince: lost.BirthedProvince,
		BirthedCity:     lost.BirthedCity,
		BirthedCountry:  lost.BirthedCountry,
		BirthedAddress:  lost.BirthedAddress,
		BirthedAt:       lost.BirthedAt,
		MissedCountry:   lost.MissedCountry,
		MissedProvince:  lost.MissedProvince,
		MissedCity:      lost.MissedCity,
		MissedAddress:   lost.MissedAddress,
		MissedAt:        lost.MissedAt,
		Handler:         lost.Handler,
		Babyid:          lost.Babyid,
		Category:        lost.Category,
		Height:          lost.Height,
		SyncStatus:      lost.SyncStatus,
	}
}

func FromPoLosts(losts []*po.Lost) []*Lost {
	_losts := make([]*Lost, 0, len(losts))

	for _, lost := range losts {
		_losts = append(_losts, FromPoLost(lost))
	}

	return _losts
}

func ListLost(req ListLostReq) []*Lost {
	losts := []*Lost{}
	losts = append(losts, FromPoLosts(po.ListLost(req.Page, req.PageSize))...)

	return losts
}

func QueryLost(id *uint) *Lost {
	return FromPoLost(po.QueryLostByID(id))
}

func SearchLost(keywords string) []*Lost {
	var losts []*Lost

	losts = append(losts, FromPoLosts(po.SearchLost(keywords))...)

	return losts
}
