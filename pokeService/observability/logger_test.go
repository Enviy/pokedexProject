package observability

import (
	"testing"

	"github.com/kubernetes/utils/pointer"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/Enviy/pokedexProject/pokeService/constants"

	"github.com/stretchr/testify/assert"

	"github.com/Enviy/pokedexProject/pokeService/entity"
)

var (
	errFoo = errors.New("errored")
)

func setupFakeProdLogger() func() {
	original := newProductionLogger
	newProductionLogger = func(options ...zap.Option) (*zap.Logger, error) {
		return nil, errFoo
	}

	return func() {
		newProductionLogger = original
	}
}

func setupFakeDevLogger() func() {
	original := newDevelopmentLogger
	newDevelopmentLogger = func(options ...zap.Option) (*zap.Logger, error) {
		return nil, errFoo
	}

	return func() {
		newDevelopmentLogger = original
	}
}

func Test_NewLogger(t *testing.T) {
	tests := []struct {
		Name   string
		Config entity.Config

		ErrOnCreateProdLogger *bool
		ErrOnCreateDevLogger  *bool

		Err error
	}{
		{
			Name: "Success development",
			Config: entity.Config{
				Env: constants.EnvDevelopment,
			},
		},
		{
			Name: "Success production",
			Config: entity.Config{
				Env: constants.EnvProduction,
			},
		},
		{
			Name: "Err development",
			Config: entity.Config{
				Env: constants.EnvDevelopment,
			},
			ErrOnCreateDevLogger: pointer.BoolPtr(true),
			Err:                  errFoo,
		},
		{
			Name: "Success production",
			Config: entity.Config{
				Env: constants.EnvProduction,
			},
			ErrOnCreateProdLogger: pointer.BoolPtr(true),
			Err:                   errFoo,
		},
	}

	for i := 0; i < len(tests); i++ {
		currTest := tests[i]
		t.Run(currTest.Name, func(t *testing.T) {

			if currTest.ErrOnCreateDevLogger != nil && *currTest.ErrOnCreateDevLogger == true {
				cleanup := setupFakeDevLogger()
				defer cleanup()
			}
			if currTest.ErrOnCreateProdLogger != nil && *currTest.ErrOnCreateProdLogger == true {
				cleanup := setupFakeProdLogger()
				defer cleanup()
			}

			logger, err := NewLogger(currTest.Config)
			if currTest.Err != nil {
				assert.Error(t, err)
				assert.Nil(t, logger)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, logger)
			}
		})
	}
}
