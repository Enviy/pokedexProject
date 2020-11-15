package server

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Enviy/pokedexProject/pokeService/entity"
	"go.uber.org/zap"

	fxMocks "github.com/Enviy/pokedexProject/pokeService/.gen/mocks/fx"
	handlerMocks "github.com/Enviy/pokedexProject/pokeService/.gen/mocks/pokeService"
	"github.com/golang/mock/gomock"
)

func Test_Server(t *testing.T) {
	ctrl := gomock.NewController(t)
	lifecycle := fxMocks.NewMockLifecycle(ctrl)
	handlerMock := handlerMocks.NewMockHandler(ctrl)

	lifecycle.EXPECT().Append(gomock.Any()).Times(1)

	properties, err := server(lifecycle, handlerMock, entity.Config{
		ServerPort: 1234,
	}, zap.S())
	assert.NoError(t, err)
	assert.NotZero(t, properties)
	assert.Equal(t, ":1234", properties.Address)
}
