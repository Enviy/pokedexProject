package server

import (
	"context"
	"strconv"

	"github.com/Enviy/pokedexProject/pokeService/constants"

	"go.uber.org/zap"

	"github.com/Enviy/pokedexProject/pokeService/entity"
	"github.com/Enviy/pokedexProject/pokeService/handler"

	"net/http"

	"go.uber.org/fx"
)

// Server creates the http server through FX
var Server = fx.Invoke(
	server,
)

// Server is the HTTP server that gets registered
func server(
	lifecycle fx.Lifecycle,
	handler handler.Handler,
	conf entity.Config,
	logger *zap.SugaredLogger,
) (entity.ServerProperties, error) {
	mux := http.NewServeMux()
	server := http.Server{
		Addr:    buildServerPort(conf.ServerPort),
		Handler: mux,
	}

	mux.HandleFunc("/", handler.Root)
	mux.HandleFunc("/pokemon", handler.Pokemon)

	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				go func() {
					logger.With(
						zap.String(constants.ServerPortField, buildServerPort(conf.ServerPort)),
					).Info("Started HTTP Server.")

					server.ListenAndServe()
				}()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				return server.Shutdown(ctx)
			},
		},
	)

	return entity.ServerProperties{
		Address: buildServerPort(conf.ServerPort),
	}, nil
}

func buildServerPort(port int) string {
	return ":" + strconv.Itoa(port)
}
