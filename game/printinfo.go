package game

import (
	"hangman-web/config"
	"html/template"
)

// PrintInfo generates and returns HTML content that wraps a message in a <p> tag.
//
// Parameters:
// - dataHangmanWeb (*config.DataHangmanWeb): The game data structure (although not used here directly, it can be useful for further extensions).
// - message (string): The message that needs to be wrapped in the <p> tag.
//
// Returns:
// - template.HTML: A template-safe HTML representation of the message wrapped inside a <p> element.
func PrintInfo(dataHangmanWeb *config.DataHangmanWeb, message string) template.HTML {
	// Return the HTML for the message wrapped in <p> tags.
	return template.HTML("<p>" + message + "</p>")
}
