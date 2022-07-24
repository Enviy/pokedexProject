package util

import (
	"bufio"
	"strings"
	"runtime"
	"os/exec"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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
func Banner() string {
	const banner = `
	__ ___         _            __     ___      ___
	| '_  \       | |    /_     | |    \  \    /  /
	| |_) |  ___  | | _____  ___| |___  \  \  /  /
	| .___/ / _ \ | |/ / _ \/  _  | _ \  \  \/  /
	| |    | (_) ||   <| __/| (_) | __/  /  /\  \
	| |     \___/ |_|\_\___/\_____|___/ /  /  \  \
	|_|                                /__/    \__\
	`
	return banner
}

// OakLine set up macOS 'say' command
func OakLine(message string) {
	if runtime.GOOS == "darwin" {
		cmd := exec.Command("say", message)
		cmd.Run()
	}
}

// WordWrap wraps text using given length maintaining formatting.
func WordWrap(text string, lineWidth int) string {
        wrappedFacts := ""
        // Nested function checks words in line for length.
        wrapper := func(text string, lineWidth int) string {
                words := strings.Fields(strings.TrimSpace(text))
                if len(words) == 0 {
                        return text
                }
                wrapped := words[0]
                spaceLeft := lineWidth - len(wrapped)
                for _, word := range words[1:] {
                        if len(word)+1 > spaceLeft {
                                wrapped += "\n" + word
                                spaceLeft = lineWidth - len(word) + 1
                        } else {
                                wrapped += " " + word
                                spaceLeft -= 1 + len(word)
                        }
                }

                return wrapped
        }
        // Check by line first to maintain formatting.
        scanner := bufio.NewScanner(strings.NewReader(text))
        for scanner.Scan() {
                if len(scanner.Text()) > lineWidth {
                        if strings.Contains(scanner.Text(), "yes or no") {
				wrappedFacts += "\n\n" + wrapper(scanner.Text(), lineWidth) + "\n"
                                continue
                        }
                        wrappedFacts += "\n" + wrapper(scanner.Text(), lineWidth)
                } else {
                        wrappedFacts += scanner.Text() + "\n"
                }
        }
        return wrappedFacts
}

// RepeatingKeyPressed returns true when key is pressed.
func RepeatKey(key ebiten.Key) bool {
        const (
                delay    = 30
                interval = 3
        )
        d := inpututil.KeyPressDuration(key)
        if d == 1 {
                return true
        }
        if d >= delay && (d-delay)%interval == 0 {
                return true
        }
        return false
}
