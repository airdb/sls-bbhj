package controller

import (
	"net/http"

	"github.com/airdb/sls-bbhj/internal/aggregate"
	"github.com/airdb/sls-bbhj/internal/repository"
	"github.com/airdb/sls-bbhj/pkg/schema"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type IndexController struct {
	aggr aggregate.Aggregate
	repo repository.Factory
}

func NewIndexController(repo repository.Factory) *IndexController {
	return &IndexController{
		repo: repo,
		aggr: aggregate.New(repo),
	}
}

func (c IndexController) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/banner", c.Banner)

	return r
}

// Index Banner
// @Summary 首页轮播图.
// @Description 首页轮播图
// @Tags    index
// @Accept  json
// @Produce json
// @Success 200 {object} schema.Response
// @Router  /v1/index/banner [get]
func (c IndexController) Banner(w http.ResponseWriter, r *http.Request) {
	resp := schema.Response{
		Data:    []schema.CarouselItem{},
		Success: true,
	}

	render.JSON(w, r, resp)
}
