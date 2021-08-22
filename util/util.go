package util

import (
	"os"
	"fmt"
	"strings"
	"runtime"
	"os/exec"
)

// FlavorSeen check if the flavor text has been seen
func FlavorSeen(value string, list []string) bool {
	for _, b := range list {
		if strings.ToLower(b) == strings.ToLower(value) {
			return true
		}
	}
	return false
}

// Banner outputs pokedex banner
func Banner() {
	const banner = `
	__ ___         _            __     ___      ___
	| '_  \       | |           | |    \  \    /  /
	| |_) |  ___  | | _____  ___| |___  \  \  /  /
	| .___/ / _ \ | |/ / _ \/  _  | _ \  \  \/  /
	| |    | (_) ||   <| __/| (_) | __/  /  /\  \
	| |     \___/ |_|\_\___/\_____|___/ /  /  \  \
	|_|                                /__/    \__\
	`
	fmt.Println(banner)
}

// CallClear clears terminal in different OSs
func CallClear() {
	var clear map[string]func()
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
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		fmt.Println("[!] This is an unsupported OS, sorry about it.")
	}
}

// OakLine set up macOS 'say' command
func OakLine(message string) {
	if runtime.GOOS == "darwin" {
		cmd := exec.Command("say", message)
		cmd.Run()
	}
}
