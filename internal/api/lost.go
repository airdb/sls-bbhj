package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/airdb/sls-bbhj/internal/repository"
	"github.com/airdb/sls-bbhj/pkg/schema"
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
	r.Get("/{id}", c.Show)

	return r
}

// LostList
// @Summary List lost item.
// @Description List item limit 10
// @Tags    lost
// @Accept  json
// @Produce json
// @param page query int false "page"
// @Success 200 {object} schema.LostListResponse
// @Router  /v1/lost [get]
// @Example /mina/v1/lost?pageNo=1&pageSize=10
func (c LostController) List(w http.ResponseWriter, r *http.Request) {
	msg := schema.LostListRequest{}

	msg.Keyword = r.URL.Query().Get("keyword")

	pageNoStr := r.URL.Query().Get("pageNo")
	msg.PageNo, _ = strconv.Atoi(pageNoStr)

	pageSizeStr := r.URL.Query().Get("pageSize")
	msg.PageSize, _ = strconv.Atoi(pageSizeStr)

	msg.Valadate()

	log.Println(msg)

	items, err := c.repo.Losts().List(r.Context(), msg)
	if err != nil {
		log.Println(err)

		return
	}

	log.Println("item len: ", len(items))

	resp := schema.LostListResponse{
		Data:    items,
		Success: true,
	}

	render.JSON(w, r, resp)
}

// LostShow
// @Summary Query lost item.
// @Description query item by id or name
// @Tags    lost
// @Accept  json
// @Produce json
// @Success 200 {string} response "api response"
// @Router  /v1/lost/{id} [get]
func (c LostController) Show(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Println(err)

		return
	}

	item, err := c.repo.Losts().GetByID(r.Context(), id)
	if err != nil {
		log.Println(err)

		return
	}

	resp := schema.LostGetResponse{
		Data:    item,
		Success: true,
	}

	render.JSON(w, r, resp)
}
