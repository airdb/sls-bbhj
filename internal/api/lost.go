package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/airdb/sls-mina/internal/repository"
	"github.com/airdb/sls-mina/pkg/schema"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type LostController struct {
	repo repository.Factory
}

func NewLostController(repo repository.Factory) *LostController {
	return &LostController{
		repo: repo,
	}
}

func (c LostController) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", c.List)
	r.Get("/{:uuids}", c.Show)

	return r
}

// LostList
// @Summary List lost item.
// @Description List item limit 10
// @Tags    lost
// @Accept  json
// @Produce json
// @Success 200 {string} response "api response"
// @Router  /v1/lost [get]
// @Example /mina/v1/lost?pageNo=1&pageSize=10
func (c LostController) List(w http.ResponseWriter, r *http.Request) {
	req := schema.LostListReq{}

	req.Keyword = r.URL.Query().Get("keyword")

	pageNoStr := r.URL.Query().Get("pageNo")
	req.PageNo, _ = strconv.Atoi(pageNoStr)

	pageSizeStr := r.URL.Query().Get("pageSize")
	req.PageSize, _ = strconv.Atoi(pageSizeStr)

	log.Println(req)

	items, err := c.repo.Losts().List(r.Context(), req)
	if err != nil {
		log.Println(err)

		return
	}

	log.Println("item", items)

	resp := schema.LostListResp{
		Data:    items,
		Success: true,
	}

	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, resp)
}

// LostShow
// @Summary Query lost item.
// @Description query item by id or name
// @Tags    lost
// @Accept  json
// @Produce json
// @Success 200 {string} response "api response"
// @Router  /v1/lost/{:uuid} [get]
func (c LostController) Show(w http.ResponseWriter, r *http.Request) {
	item, err := c.repo.Losts().GetByUUID(r.Context(), chi.URLParam(r, "uuid"))
	if err != nil {
		log.Println(err)

		return
	}

	resp := schema.LostGetResp{
		Data:    item,
		Success: true,
	}

	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, resp)
}
