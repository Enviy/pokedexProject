package main

import (
	"github.com/Enviy/pokedexProject/pokeService/app"
	"github.com/Enviy/pokedexProject/pokeService/server"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		app.Module,
		server.Server,
	).Run()
}
