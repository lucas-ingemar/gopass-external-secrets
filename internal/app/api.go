package app

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"github.com/lucas-ingemar/gopass-external-secrets/internal/middleware"
	"github.com/lucas-ingemar/gopass-external-secrets/internal/shared"
	"github.com/rs/zerolog/log"
)

func Router(api Api) http.Handler {
	r := httprouter.New()

	// mw := alice.New(middleware.ContextLogger)
	mw := alice.New()

	// "message":"GET /v1/parameter/namespace1.secret.password"}
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
		// if err != nil {
		// 	// FIXME: Maybe should be done in httpWriter func
		// 	log.Ctx(ctx).Err(err).Msg("Need to map errors to codes")
		// 	// return
		// }

		err = shared.HttpWriteResponse(ctx, rw, response, err)
		if err != nil {
			log.Ctx(ctx).Err(err).Msg("failed to write response")
		}

		// fmt.Println(namespace)
		// fmt.Println(secret)
		// fmt.Println(parameter)
		// _, err = rw.Write([]byte(`{"value": "hejsan"}`))
		// if err != nil {
		// 	log.Ctx(ctx).Err(err).Msg("writing response")
		// }
	})
}
