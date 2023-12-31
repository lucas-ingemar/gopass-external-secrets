package app

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/lucas-ingemar/gopass-external-secrets/internal/config"
	"github.com/lucas-ingemar/gopass-external-secrets/internal/pass"
	"github.com/lucas-ingemar/gopass-external-secrets/internal/shared"
	"github.com/rs/zerolog/log"
	"github.com/samber/lo"
)

func NewApp(pass pass.Pass) AppFace {
	return &App{
		lastGitSync: time.Now(),
		pass:        pass,
	}
}

type App struct {
	knownSecrets []string
	lastGitSync  time.Time
	pass         pass.Pass
}

type AppFace interface {
	GetParameter(ctx context.Context, namespace, secret, parameter string) (shared.Response, error)
	SyncGit(ctx context.Context) error
}

func (a *App) SyncGit(ctx context.Context) error {
	log.Info().Msg("syncing git")
	err := a.pass.SyncSecrets(ctx)
	if err != nil {
		return err
	}

	secrets, err := a.pass.ListSecrets(ctx)
	if err != nil {
		return err
	}
	a.knownSecrets = secrets
	a.lastGitSync = time.Now()
	return nil
}

func (a *App) GetParameter(ctx context.Context, namespace, secret, parameter string) (shared.Response, error) {
	secretsPath := fmt.Sprintf("%s/%s/%s", config.GOPASS_PREFIX, namespace, secret)
	lastSyncWithCooldown := a.lastGitSync.Add(time.Duration(config.GIT_COOLDOWN) * time.Minute)
	if !lo.Contains(a.knownSecrets, secretsPath) && time.Now().After(lastSyncWithCooldown) {
		err := a.SyncGit(ctx)
		if err != nil {
			return shared.Response{}, err
		}
	}
	s, err := a.pass.GetSecret(ctx, secretsPath)
	if err != nil {
		return shared.Response{}, err
	}

	val, ok := s[parameter]
	if !ok {
		return shared.Response{}, shared.GenericError{
			ErrorMsg:       fmt.Sprintf("parameter %s not found", parameter),
			HttpStatusCode: http.StatusNotFound,
		}
	}

	return shared.Response{
		Value: val,
	}, err
}
