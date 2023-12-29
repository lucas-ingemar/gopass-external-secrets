package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/lucas-ingemar/gopass-external-secrets/internal/app"
	"github.com/lucas-ingemar/gopass-external-secrets/internal/config"
	"github.com/lucas-ingemar/gopass-external-secrets/internal/pass"
	"github.com/lucas-ingemar/gopass-external-secrets/internal/shared"
	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	if config.DEVMODE {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	zerolog.SetGlobalLevel(config.LOG_LEVEL)

	if !shared.IsCommandAvailable("gopass") {
		log.Fatal().Msgf("gopass is not installed. Visit %s for instructions", shared.GOPASS_GITHUB_URL)
	}

	if config.AUTH_ACTIVE {
		if config.AUTH_USER == "" {
			log.Fatal().Msg("env AUTH_USER not set")
		}

		if config.AUTH_PASSWORD == "" {
			log.Fatal().Msg("env AUTH_PASSWORD not set")
		}
	}
}

func main() {
	appObj := app.NewApp(pass.GoPass{})
	appObj.SyncGit(context.Background())

	c := cron.New()
	// Define the Cron job schedule
	c.AddFunc(config.GIT_PULL_CRON, func() {
		err := appObj.SyncGit(context.Background())
		if err != nil {
			log.Err(err).Msg("git cron error")
		}
	})
	// Start the Cron job scheduler
	c.Start()

	handler := app.Router(*app.NewApi(appObj))
	log.Info().Msgf("opening on port :%s", config.API_PORT)
	log.Fatal().Err(http.ListenAndServe(fmt.Sprintf(":%s", config.API_PORT), handler))
}
