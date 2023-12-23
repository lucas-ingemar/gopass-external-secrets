package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/lucas-ingemar/gopass-external-secrets/internal/app"
	"github.com/lucas-ingemar/gopass-external-secrets/internal/config"
	"github.com/lucas-ingemar/gopass-external-secrets/internal/shared"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	if config.DEVMODE {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	if !shared.IsCommandAvailable("gopass") {
		log.Fatal().Msgf("gopass is not installed. Visit %s for instructions", shared.GOPASS_GITHUB_URL)
	}
}

func main() {
	log.Info().Msg("hej1")
	// out, err := exec.Command("gopass", "show", "namespace1/janne").Output()
	// if err != nil {
	// 	log.Err(err).Msg("gopass error")
	// }
	// fmt.Println(string(out))
	handler := app.Router()
	log.Info().Msgf("opening on port :%s", config.API_PORT)
	log.Fatal().Err(http.ListenAndServe(fmt.Sprintf(":%s", config.API_PORT), handler))
}
