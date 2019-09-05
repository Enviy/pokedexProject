package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var possible = []string{
		"exone",
		"extwo",
		"exthree",
		"exfour",
		"exfive",
		"exsix",
		"exseven",
	}
	rand.Seed(time.Now().UnixNano())
	randomize := rand.Perm(len(possible))
	for _, v := range randomize[:3] {
		fmt.Println(possible[v])
	}
}
