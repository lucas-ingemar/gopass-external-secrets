package middleware

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

func AccessLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		l := log.With().
			Str("method", r.Method).
			Str("url", r.RequestURI).
			Str("host", r.Host).
			Logger()

		l.Debug().Msgf("%s %s", r.Method, r.RequestURI)
		next.ServeHTTP(rw, r)
	})
}
