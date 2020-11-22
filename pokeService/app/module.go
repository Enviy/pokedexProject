package app

import (
	"github.com/Enviy/pokedexProject/pokeService/config"
	"github.com/Enviy/pokedexProject/pokeService/gateway"
	"github.com/Enviy/pokedexProject/pokeService/handler"
	"github.com/Enviy/pokedexProject/pokeService/observability"
	"go.uber.org/fx"
)

// Module gets exported to Fx
var Module = fx.Options(
	observability.Module,
	handler.Module,
	config.Module,
	gateway.Module,
)
