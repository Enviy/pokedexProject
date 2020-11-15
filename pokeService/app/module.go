package app

import (
	"github.com/Enviy/pokedexProject/pokeService/config"
	"github.com/Enviy/pokedexProject/pokeService/controller"
	"github.com/Enviy/pokedexProject/pokeService/db"
	"github.com/Enviy/pokedexProject/pokeService/gateway"
	"github.com/Enviy/pokedexProject/pokeService/handler"
	"github.com/Enviy/pokedexProject/pokeService/observability"
	"github.com/Enviy/pokedexProject/pokeService/secrets"
	"go.uber.org/fx"
)

// Module gets exported to Fx
var Module = fx.Options(
	observability.Module,
	handler.Module,
	config.Module,
	controller.Module,
	db.Module,
	gateway.Module,
	secrets.Module,
)
