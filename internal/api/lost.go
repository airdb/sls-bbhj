package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/airdb/sailor/dbutil"
	"github.com/airdb/sls-mina/internal/repository"
	"github.com/airdb/sls-mina/internal/repository/store"
	"github.com/airdb/sls-mina/pkg/schema"
	"github.com/go-chi/render"
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
// /mina/v1/lost/list?pageNo=1&pageSize=10
func LostList(w http.ResponseWriter, r *http.Request) {
	req := schema.LostListReq{}

	pageNoStr := r.URL.Query().Get("pageNo")
	req.PageNo, _ = strconv.Atoi(pageNoStr)

	pageSizeStr := r.URL.Query().Get("pageSize")
	req.PageSize, _ = strconv.Atoi(pageSizeStr)

	log.Println(req)

	// w.Write([]byte("welcome hello"))
	var s Reply

	mysqlStore, err := store.GetFactoryOr(dbutil.WriteDefaultDB())
	if err != nil {
		log.Println(err)
		return
	}

	s.store = mysqlStore

	items, err := s.store.Losts().List()
	if err != nil {
		log.Println(err)

		return
	}

	log.Println("item", items)

	resp := schema.LostListResp{
		Data:    items,
		Success: true,
	}

	render.JSON(w, r, resp)
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
