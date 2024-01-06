package mock

import (
	"context"
	"errors"
	"fmt"

	"github.com/lucas-ingemar/gopass-external-secrets/internal/config"
	"github.com/lucas-ingemar/gopass-external-secrets/internal/shared"
)

type Pass interface {
	GetSecret(ctx context.Context, secretPath string) (shared.PassSecret, error)
	ParseSecret(ctx context.Context, rawSecret string) (shared.PassSecret, error)
	SyncSecrets(ctx context.Context) error
	ListSecrets(ctx context.Context) ([]string, error)
}

type MockPass struct {
	namespace  string
	secret     string
	parameters []string
	// Add any fields or dependencies your struct may need
}

func NewMockPass(namespace string, secret string, parameters []string) *MockPass {
	return &MockPass{
		namespace:  namespace,
		secret:     secret,
		parameters: parameters,
	}
}

func (mp *MockPass) GetSecret(ctx context.Context, secretPath string) (shared.PassSecret, error) {
	if secretPath != fmt.Sprintf("%s/%s/%s", config.GOPASS_PREFIX, mp.namespace, mp.secret) {
		return nil, fmt.Errorf("secret path %s not found", secretPath)
	}
	ps := shared.PassSecret{}
	for _, p := range mp.parameters {
		ps[p] = fmt.Sprintf("%s_value", p)
	}
	return ps, nil
}

func (mp *MockPass) ParseSecret(ctx context.Context, rawSecret string) (shared.PassSecret, error) {
	return shared.PassSecret{}, errors.New("ParseSecret method not implemented")
}

func (mp *MockPass) SyncSecrets(ctx context.Context) error {
	return nil
}

func (mp *MockPass) ListSecrets(ctx context.Context) ([]string, error) {
	return []string{"secret1", "secret2", "secret3"}, nil
}
