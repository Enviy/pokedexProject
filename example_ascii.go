package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Would you like to search another Pokemon? (yes/no)")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	if strings.Compare("yes", text) == 0 {
		console()
	} else if strings.Compare("no", text) == 0 {
		os.Exit(0)
	}
}

func banner() {
	b, _ := ioutil.ReadFile("dex_ascii.txt")
	fmt.Println(string(b))
}

func console() {
	callClear()
	banner()
	fmt.Println("Please Enter a Pokemon to search: ")
}

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
		panic("[!] Looks like this OS is unsupported, not clearing console.")
	}
}
