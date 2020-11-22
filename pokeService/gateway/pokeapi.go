package gateway

import (
	"bytes"
	"encoding/json"
	"errors"
	"image"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
	"unicode"

	// Support decode jpeg image
	_ "image/jpeg"
	// Support deocde the png image
	_ "image/png"

	"github.com/Enviy/pokedexProject/pokeService/maps"
	"github.com/qeesung/image2ascii/convert"
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
	GetFacts(pokemon string) (*Response, error)
	GetArt(url string) (string, error)
	GetFlavor(url string) ([]string, error)
}

type pokeapi struct {
	logger *zap.SugaredLogger
}

// NewAPI creates new PokeAPI interface
func NewAPI(logger *zap.SugaredLogger) (PokeAPI, error) {
	return &pokeapi{logger: logger}, nil
}

// Response defines output of GetFacts()
type Response struct {
	Name           string
	Health         int
	Attack         int
	Defense        int
	SpecialAttack  int
	SpecialDefense int
	Speed          int
	FlavorText     []string
}

// GetFacts collects pokemon facts from pokemon api
func (p *pokeapi) GetFacts(pokemon string) (*Response, error) {
	content, err := getContent(pokeURL+pokemon, p)
	if err != nil {
		p.logger.Error("GetFacts() Error in getContent()", err)
	}
	f := maps.Basics{}
	_ = json.Unmarshal(content, &f)
	// second call; collect flavor text
	flavor, err := p.GetFlavor(f.Species.URL)
	if err != nil {
		p.logger.Error("GetFacts() Error in GetFlavor()", err)
		return nil, err
	}
	response := Response{
		Name:           f.Name,
		Health:         f.Stats[0].BaseStat,
		Attack:         f.Stats[1].BaseStat,
		Defense:        f.Stats[2].BaseStat,
		SpecialAttack:  f.Stats[3].BaseStat,
		SpecialDefense: f.Stats[4].BaseStat,
		Speed:          f.Stats[5].BaseStat,
		FlavorText:     flavor,
	}
	return &response, nil
}

// GetArt collects a png from the pokemon api
func (p *pokeapi) GetArt(pokemon string) (string, error) {
	content, err := getContent(pokeURL+pokemon, p)
	if err != nil {
		p.logger.Error("GetArt() Error in getContent()", err)
		return "", err
	}
	f := maps.SpriteURL{}
	err = json.Unmarshal(content, &f)
	if err != nil {
		p.logger.Error("GetArt() Error in json.Unmarshal()", err)
		return "", err
	}
	contents, err := getContent(f.Sprites.FrontDefault, p)
	if err != nil {
		p.logger.Error("GetArt Error in imgURL getContent()", err)
		return "", err
	}
	// log imgURL for separate testing
	p.logger.Info("imgURL:", f.Sprites.FrontDefault)
	art, err := getArt(contents, p)
	if err != nil {
		p.logger.Error("GetArt() Error in controller.GenArt()", err)
		return "", err
	}
	return art, nil
}

// GetFlavor collects flavor text for pokemon
func (p *pokeapi) GetFlavor(url string) ([]string, error) {
	content, err := getContent(url, p)
	if err != nil {
		p.logger.Error("GetFlavor() Error in getContent()", err)
		return []string{}, err
	}
	f := maps.Flavor{}
	err = json.Unmarshal(content, &f)
	if err != nil {
		p.logger.Error("GetFlavor() Error in json.Unmarshal()", err)
	}
	var s []string
	for _, flavor := range f.FlavorTextEntries {
		if strings.EqualFold(flavor.Language.Name, "en") {
			if seen(flavor.FlavorText, s) == false {
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
		if seen(s[v], finalFlavor) == false {
			finalFlavor = append(finalFlavor, s[v])
		}
	}
	return finalFlavor, nil
}

// een check if string in array is duplicate
func seen(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// getContent base initial get for Pokemon, Facts, Artwork
func getContent(url string, p *pokeapi) ([]byte, error) {
	resp, err := http.Get(url)
	if resp.StatusCode == 404 {
		p.logger.Error("[!] No information found for query", err)
		return nil, errors.New("getContent() Error in base Get() 404 response")
	}
	if err != nil {
		p.logger.Error("getContent() Error in base Get()", err)
		return nil, err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	return content, nil
}

// getArt png to ascii conversion
func getArt(content []byte, p *pokeapi) (string, error) {
	r := bytes.NewReader(content)
	img, _, err := image.Decode(r)
	if err != nil {
		p.logger.Error("getArt() Error in image.Decode()", err)
		return "", err
	}
	convertOptions := convert.DefaultOptions
	convertOptions.FixedWidth = 100
	convertOptions.FixedHeight = 40
	converter := convert.NewImageConverter()

	art := converter.Image2ASCIIString(img, &convertOptions)
	p.logger.Info(art)
	return art, nil
}
