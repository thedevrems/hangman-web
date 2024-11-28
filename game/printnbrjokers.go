package game

import (
	"hangman-web/config"
	"html/template"
	"strconv"
)

// PrintNbrJokers generates an HTML string displaying the number of jokers remaining in the game.
//
// Parameters:
// - dataHangmanWeb (*config.DataHangmanWeb): The game state, which contains the number of jokers and whether the jokers feature is enabled.
//
// Returns:
// - template.HTML: An HTML-safe string containing a message with the number of jokers remaining, or an empty paragraph if jokers are not available or enabled.
func PrintNbrJokers(dataHangmanWeb *config.DataHangmanWeb) template.HTML {
	var data string
	// Check if jokers are enabled and there are jokers remaining
	if dataHangmanWeb.GameData.NbrJokers > 0 && dataHangmanWeb.DataConfigHangman.EnableJokers {
		// Convert the number of jokers to a string and create the message
		stringNbrJokers := strconv.Itoa(dataHangmanWeb.GameData.NbrJokers)
		data = "<p>" + dataHangmanWeb.TranslationHangman.JokersCountMessage + stringNbrJokers + "</p>"
		// Return the HTML-safe string with the number of jokers
		return template.HTML(data)
	} else {
		// Return an empty paragraph if there are no jokers or the feature is disabled
		return template.HTML("<p></p>")
	}
}
