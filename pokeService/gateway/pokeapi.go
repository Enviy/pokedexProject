package gateway

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
	"unicode"

	"github.com/Enviy/pokedexProject/pokeService/maps"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var pokeURL = "https://pokeapi.co/api/v2/pokemon/"

// Module provides all of the gateways to Fx
var Module = fx.Provide(
	NewAPI,
)

// PokeAPI interacts with poke api
type PokeAPI interface {
	GetFacts(pokemon string) (*response, error)
	GetPng(url string) ([]byte, error)
	GetFlavor(url string) ([]string, error)
}

type pokeapi struct {
	logger *zap.SugaredLogger
}

// NewAPI creates new PokeAPI interface
func NewAPI(logger *zap.SugaredLogger) (PokeAPI, error) {
	return &pokeapi{logger: logger}, nil
}

type response struct {
	Name       string
	FlavorText []string
}

// GetFacts collects pokemon facts from pokemon api
func (p *pokeapi) GetFacts(pokemon string) (*response, error) {
	resp, err := http.Get(pokeURL + pokemon)
	if resp.StatusCode == 404 {
		p.logger.Error("[!] No information found for query", err)
		return nil, errors.New("GetFacts() base GET 404 response")
	}
	if err != nil {
		p.logger.Error("[!] Error in pokeapi GetFacts base Get()", err)
		return nil, err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	f := maps.Basics{}
	_ = json.Unmarshal(content, &f)
	// second call; collect flavor text
	flavor, err := p.GetFlavor(f.Species.URL)
	if err != nil {
		p.logger.Error("[!] Error in p.GetFlavor()", err)
		return nil, err
	}
	Response := response{
		Name:       f.Name,
		FlavorText: flavor,
	}
	return &Response, nil
}

// GetPng collects a png from the pokemon api
func (p *pokeapi) GetPng(url string) ([]byte, error) {
	var s []byte
	return s, nil
}

// GetFlavor collects flavor text for pokemon
func (p *pokeapi) GetFlavor(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		p.logger.Error("[!] Error in pokeapi GetFacts flavorResp Get()", err)
		return []string{}, err
	}
	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	f := maps.Flavor{}
	_ = json.Unmarshal(content, &f)
	var s []string
	for _, flavor := range f.FlavorTextEntries {
		if strings.EqualFold(flavor.Language.Name, "en") {
			if Seen(flavor.FlavorText, s) == false {
				strippedFlavor := strings.TrimFunc(flavor.FlavorText, func(r rune) bool {
					return !unicode.IsLetter(r) && !unicode.IsNumber(r)
				})
				s = append(s, strippedFlavor)
			}
		}
	}
	var finalFlavor []string
	rand.Seed(time.Now().UnixNano())
	randomize := rand.Perm(len(s))
	for _, v := range randomize[:3] {
		if Seen(s[v], finalFlavor) == false {
			finalFlavor = append(finalFlavor, s[v])
		}
	}
	return finalFlavor, nil
}

// Seen check if string in array is duplicate
func Seen(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
