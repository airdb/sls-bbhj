package api

import (
	"net/http"
	"time"

	"github.com/airdb/sailor"
	"github.com/airdb/sls-mina/pkg/schema"
	"github.com/go-chi/render"
)

// RescueList - 显示信息
// @Summary List rescue item.
// @Description List rescue limit 10
// @Tags resue
// @Accept json
// @Produce json
// @Success 200 {string} response "api response"
// @Router /rescue/list [get]
func RescueList(w http.ResponseWriter, r *http.Request) {
	var d = []*schema.Rescue{}

	d = append(d, &schema.Rescue{
		ID:        1,
		Name:      "xxx 救助站123",
		Is24Hour:  sailor.Bool(false),
		StartedAt: time.Time{},
		EndedAt:   time.Time{},
		Province:  "广东",
		City:      "深圳",
		District:  "南山区",
	})

	d = append(d, &schema.Rescue{
		ID:        2,
		Name:      "yyy 救助站123",
		Is24Hour:  sailor.Bool(false),
		StartedAt: time.Time{},
		EndedAt:   time.Time{},
		Province:  "广东",
		City:      "深圳",
		District:  "福田区",
	})

	d = append(d, &schema.Rescue{
		ID:        3,
		Name:      "zzz 救助站123",
		Is24Hour:  sailor.Bool(false),
		StartedAt: time.Time{},
		EndedAt:   time.Time{},
		Province:  "广东",
		City:      "深圳",
		District:  "宝安区",
	})

	resp := schema.RescueListResp{
		Data:    d,
		Success: true,
	}

	// w.Write([]byte("welcome hello"))

	render.JSON(w, r, resp)
	w.WriteHeader(http.StatusOK)
}

// RescueSearch
// @Summary search rescue item.
// @Description Search rescue by id or name.
// @Tags resue
// @Accept json
// @Produce json
// @Success 200 {string} response "api response"
// @Router /rescue/RescueSearch [get]
func RescueSearch(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome hello"))
	w.WriteHeader(http.StatusOK)
}
