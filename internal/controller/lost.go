package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/airdb/sls-bbhj/internal/aggregate"
	"github.com/airdb/sls-bbhj/internal/repository"
	"github.com/airdb/sls-bbhj/pkg/schema"
	"github.com/airdb/sls-bbhj/pkg/util"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
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
	r.Get("/{lost_id}/"+aggregate.LOST_WXMP_CODE_FILENAME, c.GetMpCode)

	c.aggr.Passport().Middleware(r, func(r chi.Router) {
		r.Use(func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				token, _, _ := jwtauth.FromContext(r.Context())

				if token != nil && !c.allowCreate(token.Subject()) {
					http.Error(w, "no permission", http.StatusUnauthorized)
					return
				}

				next.ServeHTTP(w, r)
			})
		})
		r.Post("/", c.Create)
	})

	return r
}

// 检查权限
func (c LostController) allowCreate(subject string) bool {
	res, err := c.aggr.Redis().Get(fmt.Sprintf("wxmp_passport:%s", subject))
	if err != nil {
		return false
	}

	return strings.Contains(res, "lost::create")
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
	var (
		resp = schema.LostGetResponse{
			Success: true,
		}
		shareKey string
	)

	id, err := strconv.Atoi(chi.URLParam(r, "lost_id"))
	if err != nil {
		log.Println(err)

		resp.Success = false
		render.JSON(w, r, resp)

		return
	}

	// shareKey := chi.URLParam(r, "share_key")
	if shareKey = chi.URLParam(r, "share_key"); shareKey == "" {
		log.Println("shareKey is empty")

		resp.Success = false
		render.JSON(w, r, resp)

		return
	}

	shareKey = strings.Join([]string{shareKey, util.RemoteIp(r)}, ":")

	if _, err = c.repo.Losts().GetByID(r.Context(), uint(id)); err != nil {
		log.Println(err)

		resp.Success = false
		render.JSON(w, r, resp)

		return
	}

	shareKeyRedisValue, err := c.aggr.Redis().Get(shareKey)
	if err != nil {
		log.Println(err)

		resp.Success = false
		render.JSON(w, r, resp)

		return
	}

	var shareCount int
	if shareKeyRedisValue != "" {
		shareCount, err = strconv.Atoi(shareKeyRedisValue)
		if err != nil {
			log.Println(err)

			resp.Success = false
			render.JSON(w, r, resp)

			return
		}
	}

	if shareCount >= 3 {
		render.JSON(w, r, resp)

		return
	}

	shareCount++

	if err = c.aggr.Redis().Set(shareKey, strconv.Itoa(shareCount), time.Second*86400); err != nil {
		log.Println(err)

		resp.Success = false
		render.JSON(w, r, resp)

		return
	}

	if err = c.repo.Losts().IncreaseShare(r.Context(), uint(id)); err != nil {
		log.Println(err)

		resp.Success = false
		render.JSON(w, r, resp)

		return
	}

	render.JSON(w, r, resp)
}

// [TODO] GetMpCode
// @Summary 失踪信息 小程序码
// @Description 通过失踪信息ID获取小程序码.
// @Tags    lost
// @Accept  json
// @Produce json
// @Param   lost_id    path  int     true  "Lost ID"
// @Success 200 {object} schema.Response
// @Router  /v1/lost/{lost_id}/wxmp_code.jpg [get]
func (c LostController) GetMpCode(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "lost_id"))
	if err != nil {
		log.Println(err)

		render.Data(w, r, []byte("id not exist"))

		return
	}

	code := c.aggr.Losts().GetWxMpCode(r.Context(), uint(id))

	w.Header().Set("Content-Type", "image/jpeg")
	if status, ok := r.Context().Value(render.StatusCtxKey).(int); ok {
		w.WriteHeader(status)
	}

	w.Write(code)
}

// LostCreate
// @Summary 失踪信息 录入。
// @Description 失踪信息 录入；权限管理暂时放于redis中，以 wxmp_permission:<open_id> 为规则，内容为权限标识，以英文逗号分隔。权限标识：lost_create。
// @Tags    lost
// @Accept  json
// @Produce json
// @Param   lost_id  path  int  true  "Lost ID"
// @Success 200 {object} schema.LostCreateResponse
// @Router  /v1/lost [post]
func (c LostController) Create(w http.ResponseWriter, r *http.Request) {
	var in schema.LostCreateRequest

	if err := render.Bind(r, &in); err != nil {
		log.Println(err)

		return
	}

	resp := schema.LostCreateResponse{}

	if err := in.Valadate(); err != nil {
		resp.Message = err.Error()

		render.JSON(w, r, resp)
		return
	}

	if err := c.aggr.Losts().Create(r.Context(), in); err != nil {
		log.Println(err)
		resp.Message = "录入信息失败，请联系管理员。"

		render.JSON(w, r, resp)
		return
	}

	resp.Success = true

	render.JSON(w, r, resp)
}
