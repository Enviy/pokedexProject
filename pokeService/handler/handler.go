package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Enviy/pokedexProject/pokeService/gateway"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Module exports all handler resources to Fx at startup
var Module = fx.Provide(
	New,
)

// Handler handles server requests
type Handler interface {
	Root(w http.ResponseWriter, r *http.Request)
	Pokemon(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	logger *zap.SugaredLogger
}

// New creates a new Handler
func New(logger *zap.SugaredLogger) (Handler, error) {
	return &handler{logger: logger}, nil
}

// Root handles the root route
func (h *handler) Root(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// Pokemon handles the pokemon route
func (h *handler) Pokemon(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		io.WriteString(w, "[!] Sorry, this isn't a GET request.")
		h.logger.Info("[!] Exiting h.Pokemon, didn't get a GET request.")
		return
	}
	keys, ok := r.URL.Query()["key"]
	if !ok || len(keys[0]) < 1 {
		h.logger.Info("URL parameter 'key' is missing.")
		return
	}
	p, err := gateway.NewAPI(h.logger)
	if err != nil {
		h.logger.Error("h.Pokemon() Error in NewAPI call", err)
	}
	pokemon, err := p.GetFacts(keys[0])
	if err != nil {
		h.logger.Error("Pokeapi GetFacts failed.", err)
		return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(pokemon)
}
