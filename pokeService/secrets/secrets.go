package secrets

import (
	"github.com/Enviy/pokedexProject/pokeService/entity"
	"github.com/Enviy/pokedexProject/pokeService/utils/goutils"
	"go.uber.org/config"
	"go.uber.org/fx"
)

var (
	secretsYAMLPath = "./secrets/secrets.yaml"
)

// Module exports all secret resources to Fx at startup
var Module = fx.Provide(
	New,
)

// New creates a new secrets entity for Fx
func New() (entity.Secrets, error) {
	provider, err := config.NewYAML(config.File(secretsYAMLPath))
	if err != nil {
		return entity.Secrets{}, goutils.ErrWrap(err)
	}

	var secrets entity.Secrets
	if err := provider.Get("").Populate(&secrets); err != nil {
		return entity.Secrets{}, goutils.ErrWrap(err)
	}

	return secrets, nil
}
