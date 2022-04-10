package theme

import (
	"fmt"
	"github.com/TwiN/go-color"
)

const (
	reset        = "\033[0m"
	bold         = "\033[1m"
	red          = "\033[31m"
	green        = "\033[32m"
	yellow       = "\033[33m"
	blue         = "\033[34m"
	purple       = "\033[35m"
	cyan         = "\033[36m"
	gray         = "\033[37m"
	white        = "\033[97m"
	InputRequest = blue
	ErrorMessage = red

	RepositoryTitle = green
	PullState       = green
	PullTitle       = gray
	PullUrl         = blue
	UserLogin       = yellow
	DebugMessage    = gray
	LogLevel        = yellow
)

const debug = false

func ColorString(messageColor string, message string) (msg string) {
	msg = color.Ize(messageColor, fmt.Sprintf(message))
	return
}

func PrintColorString(color string, message string) {
	fmt.Print(ColorString(color, message))
}

// PrintDebug TODO: Move to some logging package
func PrintDebug(message string) {
	if debug {
		PrintColorString(color.Ize(yellow, "[debug] ")+DebugMessage, message)
	}
}
