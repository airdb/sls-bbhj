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

type RescueController struct {
	repo repository.Factory
}

func NewRescueController(repo repository.Factory) *RescueController {
	return &RescueController{
		repo: repo,
	}
}

func (c RescueController) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", c.List)

	return r
}

// RescueList - 显示信息
// @Summary List rescue item.
// @Description List rescue limit 10
// @Tags resue
// @Accept json
// @Produce json
// @Success 200 {string} response "api response"
// @Router /rescue [get]
func (c RescueController) List(w http.ResponseWriter, r *http.Request) {
	req := schema.RescueListReq{}

	req.Keyword = r.URL.Query().Get("keyword")

	pageNoStr := r.URL.Query().Get("pageNo")
	req.PageNo, _ = strconv.Atoi(pageNoStr)

	pageSizeStr := r.URL.Query().Get("pageSize")
	req.PageSize, _ = strconv.Atoi(pageSizeStr)

	log.Println(req)

	items, err := c.repo.Rescues().List(r.Context(), req)
	if err != nil {
		log.Println(err)

		return
	}

	log.Println("item len: ", len(items))

	data := []*schema.RescueItem{}

	for _, item := range items {
		data = append(data, &schema.RescueItem{
			ID:        item.ID,
			Name:      item.Name,
			Is24Hour:  item.Is24Hour,
			StartedAt: item.StartedAt,
			EndedAt:   item.EndedAt,
			Province:  item.Province,
			City:      item.City,
			District:  item.District,
		})
	}

	resp := schema.RescueListResp{
		Data:    data,
		Success: true,
	}

	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, resp)
}
