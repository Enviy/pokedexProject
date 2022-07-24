package pokemon

import (
	"fmt"
	"log"
	"bytes"
	"image"
	"time"
	"regexp"
	"strings"
	"unicode"
	"math/rand"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"github.com/Enviy/pokedexProject/models"
	"github.com/Enviy/pokedexProject/convert"
	"github.com/Enviy/pokedexProject/util"
)

// GetBase initial call to api
func GetBase(name string) ([]byte, bool, error) {
	url := "https://pokeapi.co/api/v2/pokemon/" + name
	resp, err := http.Get(url)
	if resp.StatusCode == 404 {
		return nil, false, nil
	}
	if err != nil {
		return nil, false, err
	}
	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, false, err
	}
	return respBytes, true, nil
}

// GetArt collects pokemon ascii art
func GetArt(artBytes []byte) (string, error) {
	var imgs models.SpriteURLs
	err := json.Unmarshal(artBytes, &imgs)
	if err != nil {
		return "", err
	}
	imgURL := imgs.Sprites.FrontDefault
	response, err := http.Get(imgURL)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	// initialize convert options
	dops := convert.DefaultOptions
	options := &convert.Options{
		Ratio: dops.Ratio,
		FixedWidth: dops.FixedWidth,
		FixedHeight: dops.FixedHeight,
		FitScreen: dops.FitScreen,
		StretchedScreen: dops.StretchedScreen,
		Colored: false,
		Reversed: dops.Reversed,
	}
	converter := convert.NewImageConverter()
	r := bytes.NewReader(contents)
	img, _, err := image.Decode(r)
	if err != nil {
		return "", err
	}
	return converter.Image2ASCIIString(img, options), nil
}

// GetInfo parses and prints pokemon facts
func GetInfo(baseBytes []byte) string {
	responseString := ""
	var basics models.Basics
	err := json.Unmarshal(baseBytes, &basics)
	if err != nil {
		log.Fatal(err)
	}
	responseString = fmt.Sprintf("Ah, the %s\n", basics.Name)
	URL := basics.Species.URL
	responseString += "It's been observed with these base stats:\n"
	for _, BasicEntry := range basics.Stats {
		responseString += fmt.Sprintf("%s: %v\n", BasicEntry.Stat.Name, BasicEntry.BaseStat)
	}
	// get flavor text
	response, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var flavor models.Flavor
	err = json.Unmarshal(contents, &flavor)
	if err != nil {
		log.Fatal(err)
	}
	flavorArray := deduplicate(flavor)
	responseString += "\nField Notes:\n"
	responseString += selectThree(flavorArray)
	responseString += "Would you like to search for another Pokemon? (yes or no)"
	return responseString
}

// deduplicate removes duplicate flavor text entries
func deduplicate(flavorText models.Flavor) []string {
	var deduped []string
	for _, flavor := range flavorText.FlavorTextEntries {
		if strings.EqualFold(flavor.Language.Name, "en") {
			if util.FlavorSeen(flavor.FlavorText, deduped) == false {
				strippedFlavor := strings.TrimFunc(flavor.FlavorText, func(r rune) bool {
					return !unicode.IsLetter(r) && !unicode.IsNumber(r)
				})
				regexed := regexp.MustCompile(`\s+`)
				strippedFlavor = regexed.ReplaceAllString(strippedFlavor, " ")
				deduped = append(deduped, strippedFlavor)
			}
		}
	}
	return deduped
}

// selectThree collects three sumbissions from the deduped flavor text entries
func selectThree(flavorArray []string) string {
	responseString := ""
	var finalFlavor []string
	rand.Seed(time.Now().UnixNano())
	randomize := rand.Perm(len(flavorArray))
	for _, v := range randomize[:2] {
		if util.FlavorSeen(flavorArray[v], finalFlavor) == false {
			finalFlavor = append(finalFlavor, flavorArray[v])
		}
	}
	for _, finalEntry := range finalFlavor {
		responseString += finalEntry + "\n"
	}
	return responseString
}
