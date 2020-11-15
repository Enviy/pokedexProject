package db

import (
	"github.com/Enviy/pokedexProject/pokeService/db/postgres"
	"go.uber.org/fx"
)

// Module exports the DB resources to Fx at startup
var Module = fx.Provide(
	postgres.New,
)
