package app

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"github.com/lucas-ingemar/gopass-external-secrets/internal/config"
	"github.com/lucas-ingemar/gopass-external-secrets/internal/middleware"
	"github.com/lucas-ingemar/gopass-external-secrets/internal/shared"
	"github.com/rs/zerolog/log"
)

func Router(api Api) http.Handler {
	r := httprouter.New()

	mw := alice.New()
	if config.AUTH_ACTIVE {
		mw.Append(middleware.AuthCheck)
	}

	r.Handler("GET", "/v1/parameter/:namespace/:secret/:parameter", mw.Then(api.GetParameter()))

	preRouterMw := alice.New(middleware.AccessLog)
	return preRouterMw.Then(r)
}

type Api struct {
	app AppFace
}

func NewApi(app AppFace) *Api {
	return &Api{
		app: app,
	}
}

func (a Api) GetParameter() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		params := httprouter.ParamsFromContext(ctx)
		namespace := params.ByName("namespace")
		secret := params.ByName("secret")
		parameter := params.ByName("parameter")

		response, err := a.app.GetParameter(ctx, namespace, secret, parameter)

		err = shared.HttpWriteResponse(ctx, rw, response, err)
		if err != nil {
			log.Ctx(ctx).Err(err).Msg("failed to write response")
		}
	})
}
