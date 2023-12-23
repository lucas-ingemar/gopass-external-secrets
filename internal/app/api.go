package app

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"github.com/lucas-ingemar/gopass-external-secrets/internal/middleware"
)

func Router() http.Handler {
	api := NewApi()
	r := httprouter.New()

	// mw := alice.New(middleware.ContextLogger)
	mw := alice.New()

	r.Handler("GET", "/v1/parameter/:parameter", mw.Then(api.GetParameter()))

	preRouterMw := alice.New(middleware.AccessLog)
	return preRouterMw.Then(r)
}

type Api struct {
}

func NewApi() *Api {
	return &Api{}
}

func (a Api) GetParameter() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
	})
}
