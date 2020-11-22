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
	Artwork(w http.ResponseWriter, r *http.Request)
	Facts(w http.ResponseWriter, r *http.Request)
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
	type combinedResponse struct {
		Art   string
		Facts *gateway.Response
	}
	if r.Method != http.MethodGet {
		io.WriteString(w, "[!] Sorry, this isn't a GET request.")
		h.logger.Info("[!] Exiting h.Pokemon, didn't get a GET request.")
		return
	}
	keys, ok := r.URL.Query()["name"]
	// log affirmation of request recieved for additional testing
	h.logger.Info("Request recieved.", keys)
	if !ok || len(keys[0]) < 1 {
		h.logger.Info("URL parameter 'name' is missing.")
		return
	}
	p, err := gateway.NewAPI(h.logger)
	if err != nil {
		h.logger.Error("h.Pokemon() Error in NewAPI() call", err)
	}
	pokemon, err := p.GetFacts(keys[0])
	if err != nil {
		h.logger.Error("h.Pokemon() Error in p.GetFacts()", err)
		return
	}
	art, err := p.GetArt(keys[0])
	if err != nil {
		h.logger.Error("h.Pokemon() Error in p.GetArt()", err)
		return
	}
	cr := combinedResponse{
		Art:   art,
		Facts: pokemon,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cr)
}

// Artwork collects a png, converts to ascii art
func (h *handler) Artwork(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		io.WriteString(w, "[!] Sorry, this isn't a GET request.")
		h.logger.Info("[!] Exiting h.Pokemon(), didn't get a GET request.")
		return
	}
	keys, ok := r.URL.Query()["name"]
	if !ok || len(keys[0]) < 1 {
		h.logger.Info("URL parameter 'name' is missing.")
		return
	}
	p, err := gateway.NewAPI(h.logger)
	if err != nil {
		h.logger.Error("h.Artwork() Error in gateway.NewAPI()", err)
	}
	art, err := p.GetArt(keys[0])
	if err != nil {
		h.logger.Error("h.Artwork() Error in p.GetArt()", err)
		return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(art)
}

// Facts collects a name, base stats, flavor text
func (h *handler) Facts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		io.WriteString(w, "[!] Sorry, this isn't a GET request.")
		h.logger.Info("[!] Exiting h.Pokemon(), didn't get a GET request.")
		return
	}
	keys, ok := r.URL.Query()["name"]
	if !ok || len(keys[0]) < 1 {
		h.logger.Info("URL parameter 'name' is missing.")
		return
	}
	p, err := gateway.NewAPI(h.logger)
	if err != nil {
		h.logger.Error("h.Facts() Error in gateway.NewAPI()", err)
	}
	facts, err := p.GetFacts(keys[0])
	if err != nil {
		h.logger.Error("h.Facts() Error p.GetFacts()", err)
		return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(facts)
}
