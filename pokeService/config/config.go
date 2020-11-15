package config

import (
	"os"

	"github.com/Enviy/pokedexProject/pokeService/constants"
	"github.com/Enviy/pokedexProject/pokeService/entity"
	"github.com/Enviy/pokedexProject/pokeService/utils/goutils"
	"go.uber.org/config"
	"go.uber.org/fx"
)

var (
	serverEnvVariable = "SERVER_ENV"

	baseYAMLPath       = "./config/base.yaml"
	productionYAMLPath = "./config/production.yaml"
)

// Module exports all config resources to Fx at startup
var Module = fx.Provide(
	New,
)

// New creates a new Config for the server
func New() (entity.Config, error) {
	// If we have the SERVER_ENV variable set, and its production, we are in production, else development
	env := constants.EnvDevelopment
	yamlPath := baseYAMLPath
	if os.Getenv(serverEnvVariable) == constants.EnvProduction {
		env = constants.EnvProduction
		yamlPath = productionYAMLPath
	}

	provider, err := config.NewYAML(config.File(yamlPath))
	if err != nil {
		return entity.Config{}, goutils.ErrWrap(err)
	}

	var conf entity.Config
	if err := provider.Get("").Populate(&conf); err != nil {
		return entity.Config{}, goutils.ErrWrap(err)
	}

	conf.Env = env
	return conf, nil
}
