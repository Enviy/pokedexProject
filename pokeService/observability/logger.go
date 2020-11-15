package observability

import (
	"github.com/Enviy/pokedexProject/pokeService/constants"
	"github.com/Enviy/pokedexProject/pokeService/entity"
	"github.com/Enviy/pokedexProject/pokeService/utils/goutils"
	"go.uber.org/zap"
)

// For better testing coverage
var (
	newProductionLogger  = zap.NewProduction
	newDevelopmentLogger = zap.NewDevelopment
)

// NewLogger creates a new SugaredLogger for the server
func NewLogger(
	config entity.Config,
) (*zap.SugaredLogger, error) {
	var logger *zap.Logger
	if config.Env == constants.EnvProduction {
		var err error
		logger, err = newProductionLogger()
		if err != nil {
			return nil, goutils.ErrWrap(err)
		}
	} else {
		var err error
		logger, err = newDevelopmentLogger()
		if err != nil {
			return nil, goutils.ErrWrap(err)
		}
	}

	defer logger.Sync()
	return logger.With(
		zap.String(constants.AppIDField, config.AppID),
	).Sugar(), nil
}
