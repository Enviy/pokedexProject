package pokemon

import (
	"fmt"
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
		fmt.Println("[!] No information available. Maybe it's spelled wrong?")
		return nil, false, nil
	}
	if err != nil {
		fmt.Printf("[!] Error in api get request: %v", err)
		return nil, false, err
	}
	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("[!] Error reading response body: %v", err)
		return nil, false, err
	}
	return respBytes, true, nil
}

// GetArt collects pokemon ascii art
func GetArt(artBytes []byte) string {
	var imgs models.SpriteURLs
	err := json.Unmarshal(artBytes, &imgs)
	if err != nil {
		fmt.Printf("[!] Unable to unmarshal artBytes: %+v", err)
	}
	imgURL := imgs.Sprites.FrontDefault
	response, err := http.Get(imgURL)
	if err != nil {
		fmt.Printf("[!] Error in Get(imgURL): %+v", err)
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("[!] Unable to read response body: %+v", err)
	}
	// initialize convert options
	dops := convert.DefaultOptions
	options := &convert.Options{
		Ratio: dops.Ratio,
		FixedWidth: dops.FixedWidth,
		FixedHeight: dops.FixedHeight,
		FitScreen: dops.FitScreen,
		StretchedScreen: dops.StretchedScreen,
		Colored: dops.Colored,
		Reversed: dops.Reversed,
	}
	converter := convert.NewImageConverter()
	r := bytes.NewReader(contents)
	img, _, err := image.Decode(r)
	if err != nil {
		fmt.Printf("[!] Unable to decode image: %+v", err)
	}
	return converter.Image2ASCIIString(img, options)
}

// GetInfo parses and prints pokemon facts
func GetInfo(baseBytes []byte) {
	var basics models.Basics
	err := json.Unmarshal(baseBytes, &basics)
	if err != nil {
		fmt.Printf("[!] Unable to unmarshal baseBytes: %+v", err)
	}
	fmt.Printf("Ah, the %s\n", basics.Name)
	util.OakLine("Ah, the " + basics.Name)
	URL := basics.Species.URL
	fmt.Println("It's been observed with these base stats:")
	for _, BasicEntry := range basics.Stats {
		fmt.Println(BasicEntry.Stat.Name, BasicEntry.BaseStat)
	}
	// get flavor text
	response, err := http.Get(URL)
	if err != nil {
		fmt.Printf("[!] Error in flavortext GET: %+v", err)
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("[!] Unable to read flavortext body: %+v", err)
	}
	var flavor models.Flavor
	err = json.Unmarshal(contents, &flavor)
	if err != nil {
		fmt.Printf("[!] Unable to unmarshal flavortext: %+v", err)
	}
	flavorArray := deduplicate(flavor)
	fmt.Println("\nField Notes:")
	selectThree(flavorArray)
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
func selectThree(flavorArray []string) {
	var finalFlavor []string
	rand.Seed(time.Now().UnixNano())
	randomize := rand.Perm(len(flavorArray))
	for _, v := range randomize[:3] {
		if util.FlavorSeen(flavorArray[v], finalFlavor) == false {
			finalFlavor = append(finalFlavor, flavorArray[v])
		}
	}
	for _, finalEntry := range finalFlavor {
		fmt.Println("\n", finalEntry)
		util.OakLine(finalEntry)
	}
}
