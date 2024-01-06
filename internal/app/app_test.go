package app

import (
	"context"
	"testing"
	"time"

	"github.com/lucas-ingemar/gopass-external-secrets/internal/mock"
	"github.com/stretchr/testify/assert"
)

func TestSyncGit(t *testing.T) {
	startTime := time.Now()
	app := App{
		knownSecrets: []string{},
		lastGitSync:  startTime,
		pass:         mock.NewMockPass("namespace1", "secret1", []string{"username", "password", "something"}),
	}

	err := app.SyncGit(context.Background())
	assert.Nil(t, err, "no error")
	assert.Equal(t, []string{"secret1", "secret2", "secret3"}, app.knownSecrets, "list secrets")
	assert.Greater(t, app.lastGitSync, startTime, "sync time")
}

func TestGetParameter(t *testing.T) {
	startTime := time.Now()
	app := App{
		knownSecrets: []string{},
		lastGitSync:  startTime,
		pass:         mock.NewMockPass("namespace1", "secret1", []string{"username", "password", "something"}),
	}

	resp, err := app.GetParameter(context.Background(), "namespace1", "secret1", "username")
	assert.Nil(t, err, "no error")
	assert.Equal(t, "username_value", resp.Value, "return value")
}
