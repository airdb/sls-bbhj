package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/airdb/sls-bbhj/internal/aggregate"
	"github.com/airdb/sls-bbhj/internal/repository"
	"github.com/airdb/sls-bbhj/pkg/schema"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type LostController struct {
	aggr aggregate.Aggregate
	repo repository.Factory
}

func NewLostController(repo repository.Factory) *LostController {
	return &LostController{
		repo: repo,
		aggr: aggregate.New(repo),
	}
}

func (c LostController) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", c.List)
	r.Get("/{lost_id}", c.Show)
	r.Get("/{lost_id}/share/{share_key}/callback", c.ShareCallback)

	return r
}

// LostList
// @Summary 失踪信息 列表。
// @Description 失踪信息 列表 默认单页大小为10。
// @Tags    lost
// @Accept  json
// @Produce json
// @Param   pageNo   query int false "page number"
// @Param   pageSize query int false "page size"
// @Param   keyword  query int false "search keyword"
// @Param   category query int false "lost category id"
// @Success 200 {object} schema.LostListResponse
// @Router  /v1/lost [get]
// @Example /mina/v1/lost?pageNo=1&pageSize=10
func (c LostController) List(w http.ResponseWriter, r *http.Request) {
	msg := schema.LostListRequest{}

	pageNoStr := r.URL.Query().Get("pageNo")
	msg.PageNo, _ = strconv.Atoi(pageNoStr)

	pageSizeStr := r.URL.Query().Get("pageSize")
	msg.PageSize, _ = strconv.Atoi(pageSizeStr)

	msg.Keyword = r.URL.Query().Get("keyword")
	msg.Category = r.URL.Query().Get("category")

	msg.Valadate()

	log.Println(msg)

	items, err := c.aggr.Losts().List(r.Context(), msg)
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
// @Summary 失踪信息 详情。
// @Description 失踪信息 详情。lost_id为对应列表页中的id.
// @Tags    lost
// @Accept  json
// @Produce json
// @Param   lost_id  path  int  true  "Lost ID"
// @Success 200 {object} schema.LostGetResponse
// @Router  /v1/lost/{lost_id} [get]
func (c LostController) Show(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "lost_id"))
	if err != nil {
		log.Println(err)

		return
	}

	item, err := c.aggr.Losts().GetByID(r.Context(), uint(id))
	if err != nil {
		log.Println(err)

		return
	}

	if err = c.repo.Losts().IncreaseShow(r.Context(), uint(id)); err != nil {
		log.Printf("increase lost show failed: %s", err.Error())
	}

	resp := schema.LostGetResponse{
		Data:    item,
		Success: true,
	}

	render.JSON(w, r, resp)
}

// [TODO] LostShareCallback
// @Summary 失踪信息 分享后回传
// @Description 分享后回传，通过缓存+IP+Key来去重。share_key为详情页中 wx_more 里的 share_key.
// @Tags    lost
// @Accept  json
// @Produce json
// @Param   lost_id    path  int     true  "Lost ID"
// @Param   share_key  path  string  true  "Share Key from WxMore"
// @Success 200 {object} schema.Response
// @Router  /v1/lost/{lost_id}/share/{share_key}/callback [get]
func (c LostController) ShareCallback(w http.ResponseWriter, r *http.Request) {
	resp := schema.LostGetResponse{
		Success: true,
	}

	render.JSON(w, r, resp)
}
