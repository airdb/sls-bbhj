package controller

import (
	"log"
	"net/http"

	"github.com/airdb/sls-bbhj/internal/aggregate"
	"github.com/airdb/sls-bbhj/internal/repository"
	"github.com/airdb/sls-bbhj/pkg/schema"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type PassportController struct {
	aggr aggregate.Aggregate
	repo repository.Factory
}

func NewPassportController(repo repository.Factory) *PassportController {
	return &PassportController{
		repo: repo,
		aggr: aggregate.New(repo),
	}
}

func (c PassportController) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", c.Login)

	return r
}

// PassportLogin
// @Summary 授权 登录。
// @Description 授权 登录。
// @Tags    passport
// @Accept  json
// @Produce json
// @Param   code   body string true "登录凭证code"
// @Success 200 {object} schema.PassportLoginResponse
// @Router  /v1/passport:login [post]
// @Example /mina/v1/passport
func (c PassportController) Login(w http.ResponseWriter, r *http.Request) {
	in := schema.PassportLoginRequest{}

	if err := render.Bind(r, &in); err != nil {
		log.Println(err)

		return
	}

	in.Valadate()

	token, err := c.aggr.Passport().Login(r.Context(), in.Code)
	if err != nil {
		log.Println(err)

		return
	}

	resp := schema.PassportLoginResponse{
		Token:   token,
		Success: true,
	}

	render.JSON(w, r, resp)
}
