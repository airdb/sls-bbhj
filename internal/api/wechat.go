package api

import (
	"net/http"
)

// CheckSession
// @Summary Check session.
// @Description Check session.
// @Tags wechat
// @Accept json
// @Produce json
// @Success 200 {string} response "api response"
// @Router /lost/list [get]
func CheckSession(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome hello"))
	w.WriteHeader(http.StatusOK)
}
