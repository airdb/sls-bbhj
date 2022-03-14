package aggregate

import (
	"context"
	"time"

	"github.com/airdb/sls-bbhj/internal/repository"
	"github.com/airdb/sls-bbhj/pkg/util"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

// PassportAggr defines functions used to handle passport request.
type PassportAggr interface {
	Middleware(r chi.Router, f func(chi.Router))
	Login(ctx context.Context, code string) (string, error)
}

type passportAggr struct {
	repo repository.Factory
	aggr aggregate
}

var _ PassportAggr = (*passportAggr)(nil)

var tokenAuth *jwtauth.JWTAuth

func init() {
	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)
}

func newPassport(aggr *aggregate) *passportAggr {
	return &passportAggr{repo: aggr.repo}
}

func (u *passportAggr) Middleware(r chi.Router, f func(chi.Router)) {
	r.Group(func(r chi.Router) {
		// Seek, verify and validate JWT tokens
		r.Use(jwtauth.Verifier(tokenAuth))

		// Handle valid / invalid tokens. In this example, we use
		// the provided authenticator middleware, but you can write your
		// own very easily, look at the Authenticator method in jwtauth.go
		// and tweak it, its not scary.
		r.Use(jwtauth.Authenticator)

		f(r)
	})

}

// Login returns passport login result.
func (u *passportAggr) Login(ctx context.Context, code string) (string, error) {
	wx := util.NewWechatMiniProgram(util.NewWechat())
	oid, err := wx.Code2SessionContext(ctx, code)
	if err != nil {
		return "", err
	}

	claims := map[string]interface{}{
		"subject": oid,
	}
	jwtauth.SetIssuedNow(claims)
	jwtauth.SetExpiryIn(claims, time.Hour*3)

	_, tokenString, err := tokenAuth.Encode(claims)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
