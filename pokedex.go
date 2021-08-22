package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/Enviy/pokedexProject/util"
	"github.com/Enviy/pokedexProject/pokemon"
)

func main() {
	console()
}

// Pokemon collect pokemon art and info
func Pokemon(name string) {
	// Call Pokemon API
	contents, found, err := pokemon.GetBase(name)
	if err != nil {
		fmt.Printf("[!] Error in GET request: %+v", err)
	}
	if !found {
		decide()
	}
	// get .png URL and generate ascii art
	fmt.Println(pokemon.GetArt(contents))

	// Extracts & prints Pokemon name & stats
	pokemon.GetInfo(contents)

	// Ask user to choos other search or exit
	decide()
}

// Setup primary starting point for pokedex searches
func console() {
	util.CallClear()
	util.Banner()
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please Search a Pokemon's Name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	if name == "\n" {
		console()
	}
	name = strings.TrimFunc(name, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
	Pokemon(strings.ToLower(name))
}

// Setup decision point to continue or exit after initial search
func decide() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\nWould you like to search for another Pokemon? (yes/no)")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	if strings.Compare("yes", text) == 0 {
		console()
	} else if strings.Compare("no", text) == 0 {
		util.CallClear()
		os.Exit(0)
	} else {
		util.CallClear()
		util.Banner()
		decide()
	}
}
