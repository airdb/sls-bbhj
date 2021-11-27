package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/airdb/sls-bbhj/internal/repository"
	"github.com/airdb/sls-bbhj/pkg/schema"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type CategoryController struct {
	repo repository.Factory
}

func NewCategoryController(repo repository.Factory) *CategoryController {
	return &CategoryController{
		repo: repo,
	}
}

func (c CategoryController) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", c.List)

	return r
}

// CategoryList - 查询分类
// @Summary List lost's category item.
// @Description List lost's category limit 10
// @Tags category
// @Accept json
// @Produce json
// @Success 200 {object} schema.CategoryListResponse
// @Router /v1/category [get]
func (c CategoryController) List(w http.ResponseWriter, r *http.Request) {
	msg := schema.CategoryListRequest{}

	msg.Keyword = r.URL.Query().Get("keyword")

	pageNoStr := r.URL.Query().Get("pageNo")
	msg.PageNo, _ = strconv.Atoi(pageNoStr)

	pageSizeStr := r.URL.Query().Get("pageSize")
	msg.PageSize, _ = strconv.Atoi(pageSizeStr)

	msg.Valadate()

	log.Println(msg)

	items, err := c.repo.Categories().List(r.Context(), msg)
	if err != nil {
		log.Println(err)

		return
	}

	log.Println("item len: ", len(items))

	data := []*schema.CategoryItem{}

	for _, item := range items {
		data = append(data, &schema.CategoryItem{
			Value: item.ID,
			Label: item.Name,
		})
	}

	resp := schema.CategoryListResponse{
		Data:    data,
		Success: true,
	}

	render.JSON(w, r, resp)
}
