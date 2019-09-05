package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
	"unicode"
)

func main() {
	console()
}

func pokemon(name string) {
	// Call Pokemon API
	url := "https://pokeapi.co/api/v2/pokemon/" + name
	resp, err := http.Get(url)
	if resp.StatusCode == 404 {
		fmt.Println("[!] No information found for that query :(")
		time.Sleep(1 * time.Second)
		decide()
	}
	if err != nil {
		fmt.Println("[!] There was an error!", err)
	}
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)

	type Basics struct {
		Name    string `json:"name"`
		Order   int    `json:"order"`
		Species struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"species"`
		Stats []struct {
			BaseStat int `json:"base_stat"`
			Effort   int `json:"effort"`
			Stat     struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"stat"`
		} `json:"stats"`
	}

	// Extracts & prints Pokemon name & stats
	f := Basics{}
	_ = json.Unmarshal(contents, &f)
	fmt.Println("Ah, the", f.Name)
	url2 := f.Species.URL

	fmt.Println("It reportedly has the base stats:")
	for _, BasicEntry := range f.Stats {
		fmt.Println(BasicEntry.Stat.Name, BasicEntry.BaseStat)
	}

	// 2nd API call for flavor text
	response, _ := http.Get(url2)
	defer response.Body.Close()
	contents2, _ := ioutil.ReadAll(response.Body)

	type Flav struct {
		FlavorTextEntries []struct {
			FlavorText string `json:"flavor_text"`
			Language   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"language"`
		} `json:"flavor_text_entries"`
	}

	fl := Flav{}
	_ = json.Unmarshal(contents2, &fl)
	fmt.Println("\nField Notes:")

	// Deduplicate flavor text, append to array of string
	var s []string
	for _, flavorEntry := range fl.FlavorTextEntries {
		if strings.EqualFold(flavorEntry.Language.Name, "en") {
			if flavSeen(flavorEntry.FlavorText, s) == false {
				strippedFlavorText := strings.TrimFunc(flavorEntry.FlavorText, func(r rune) bool {
					return !unicode.IsLetter(r) && !unicode.IsNumber(r)
				})
				s = append(s, strippedFlavorText)
			}
		}
	}
	// Randomly select 3 flavor texts that are not the same to print
	var finalS []string
	rand.Seed(time.Now().UnixNano())
	randomize := rand.Perm(len(s))
	for _, v := range randomize[:3] {
		if flavSeen(s[v], finalS) == false {
			finalS = append(finalS, s[v])
		}
	}
	for _, finalEntry := range finalS {
		fmt.Println("\n", finalEntry)
	}

	// Ask user to choos other search or exit
	fmt.Println("\n")
	decide()
}

// Check if a string in an array has already been seen
func flavSeen(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// Setup banner again
func banner() {
	const a = `
   __ ___         _            __     ___      ___
   | '_  \       | |           | |    \  \    /  /
   | |_) |  ___  | | _____  ___| |___  \  \  /  /
   | .___/ / _ \ | |/ / _ \/  _  | _ \  \  \/  /
   | |    | (_) ||   <| __/| (_) | __/  /  /\  \
   | |     \___/ |_|\_\___/\_____|___/ /  /  \  \
   |_|                                /__/    \__\
`
	fmt.Println(a)
}

// Setup clear function for different OS'
var clear map[string]func()

func init() {
	clear = make(map[string]func())
	clear["darwin"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func callClear() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("[!] Looks like an unsupported OS sucka, not clearing this screen.")
	}
}

// Setup primary starting point for pokedex searches
func console() {
	callClear()
	banner()
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please Enter a Pokemon Name to Search: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimFunc(name, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
	pokemon(name)
}

// Setup decision point to continue or exit after initial search
func decide() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Would you like to search another Pokemon? (yes/no)")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	if strings.Compare("yes", text) == 0 {
		console()
	} else if strings.Compare("no", text) == 0 {
		callClear()
		os.Exit(0)
	}
}
