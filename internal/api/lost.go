package api

import (
	"net/http"
)

// LostList
// @Summary List lost item.
// @Description List item limit 10
// @Tags lost
// @Accept json
// @Produce json
// @Success 200 {string} response "api response"
// @Router /lost/list [get]
func LostList(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome hello"))
	w.WriteHeader(http.StatusOK)
}

// LostSearch
// @Summary Search lost item.
// @Description List item limit 10
// @Tags lost
// @Accept json
// @Produce json
// @Success 200 {string} response "api response"
// @Router /lost/search [get]
func LostSearch(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome hello"))
	w.WriteHeader(http.StatusOK)
}
