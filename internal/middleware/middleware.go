package middleware

import (
	"fmt"
	"net/http"

	"github.com/lucas-ingemar/gopass-external-secrets/internal/config"
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

func AuthCheck(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			log.Error().Msg("auth on bad format")
			rw.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(rw, "%d error", http.StatusUnauthorized)
			return
		}
		if username != config.AUTH_USER || password != config.AUTH_PASSWORD {
			log.Error().Msg("not authorized")
			rw.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(rw, "%d error", http.StatusUnauthorized)
			return

		}
		next.ServeHTTP(rw, r)
	})
}
