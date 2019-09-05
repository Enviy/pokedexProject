package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("pokemon: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimFunc(name, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
	url := "https://pokeapi.co/api/v2/pokemon/" + name
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("[!] Something went wrong!", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
