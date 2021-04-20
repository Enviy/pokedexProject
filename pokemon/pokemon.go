package pokemon

import (
	"io/ioutil"
	"net/http"

	"go.uber.org/zap"
)

const baseURL = "https://pokeapi.co/api/v2/pokemon/"

// Pokemon interface implements methods interacting with API
type Pokemon interface {
	GetBase(name string) ([]byte, error)
}

type pokemon struct {
	logger zap.SugaredLogger
}

// New returns Pokemon interface for API
func New(logger zap.SugaredLogger) Pokemon {
	return &pokemon{logger: logger}
}

// GetBase collects base pokemon data by name
func (p *pokemon) GetBase(name string) ([]byte, error) {
	resp, err := http.Get(baseURL + name)
	if err != nil {
		p.logger.Error("[!] Get failure in GetBase()", zap.Error(err))
		return nil, err
	}
	if resp.StatusCode == 404 {
		p.logger.Error("404 response for pokemon name")
		// call decide()
	}
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		p.logger.Error("failure to read resp.Body in GetBase()", zap.Error(err))
		return nil, err
	}
	return contents, nil
}
