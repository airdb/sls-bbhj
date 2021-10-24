package api

import (
	"log"
	"net/http"

	"github.com/airdb/sailor/dbutil"
	"github.com/airdb/sls-mina/internal/repository"
	"github.com/airdb/sls-mina/internal/repository/store"
)

type Reply struct {
	store repository.Factory
}

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
	var s Reply

	mysqlStore, err := store.GetFactoryOr(dbutil.WriteDefaultDB())

	s.store = mysqlStore

	err = s.store.Losts().List()
	if err != nil {
		log.Println(err)
	}
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
