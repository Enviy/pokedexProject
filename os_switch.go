package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("This is a test, it should print the current OS.")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("It's MacOS!")
	case "linux":
		fmt.Println("It's listed as Linux some crazy how.")
	default:
		fmt.Println("%s.\n", os)
	}
}
