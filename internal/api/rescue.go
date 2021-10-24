package api

import (
	"net/http"
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
	w.Write([]byte("welcome hello"))
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
