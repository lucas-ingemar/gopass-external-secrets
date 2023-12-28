package pass

import (
	"context"
	"strings"

	"github.com/lucas-ingemar/gopass-external-secrets/internal/shared"
	"github.com/lucas-ingemar/gopass-external-secrets/internal/system"
)

type Pass interface {
	GetSecret(ctx context.Context, secretPath string) (shared.PassSecret, error)
	ParseSecret(ctx context.Context, rawSecret string) (shared.PassSecret, error)
	SyncSecrets(ctx context.Context) error
	ListSecrets(ctx context.Context) ([]string, error)
}

type GoPass struct {
}

func (gp GoPass) GetSecret(ctx context.Context, secretPath string) (shared.PassSecret, error) {
	res, err := system.Call().Cmd("gopass").Args([]string{"show", secretPath}).Exec(ctx)
	if err != nil || res.ExitCode != 0 {
		return nil, shared.MapError(res.ExitCode, res.Stderr)
	}
	return gp.ParseSecret(ctx, res.Stdout)
}

func (gp GoPass) ParseSecret(ctx context.Context, rawSecret string) (shared.PassSecret, error) {
	secret := shared.PassSecret{}
	lines := strings.Split(rawSecret, "\n")
	secret["password"] = lines[0]

	for _, l := range lines[1:] {
		vals := strings.Split(l, ":")
		if len(vals) < 2 {
			continue
		}
		secret[strings.TrimSpace(vals[0])] = strings.TrimSpace(strings.Join(vals[1:], ":"))
	}

	return secret, nil
}

func (gp GoPass) SyncSecrets(ctx context.Context) error {
	res, err := system.Call().Cmd("gopass").Args([]string{"git", "pull"}).Exec(ctx)
	if err != nil || res.ExitCode != 0 {
		return shared.MapError(res.ExitCode, res.Stderr)
	}
	return nil
}

func (gp GoPass) ListSecrets(ctx context.Context) ([]string, error) {
	res, err := system.Call().Cmd("gopass").Args([]string{"ls", "-f"}).Exec(ctx)
	if err != nil || res.ExitCode != 0 {
		return nil, shared.MapError(res.ExitCode, res.Stderr)
	}
	return strings.Split(strings.TrimSpace(res.Stdout), "\n"), nil
}
