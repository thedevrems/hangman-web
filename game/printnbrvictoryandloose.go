package game

import (
	"hangman-web/config"
	"html/template"
	"strconv"
)

// PrintNbrVictoryAndLoose generates an HTML string displaying the number of victories and defeats the player has accumulated.
//
// Parameters:
// - dataHangmanWeb (*config.DataHangmanWeb): The game state, which contains the number of victories and defeats, and whether the victory counter is enabled.
//
// Returns:
// - template.HTML: An HTML-safe string containing a message displaying the current number of victories and defeats, or an empty paragraph if the counter is disabled or there are no victories or defeats.
func PrintNbrVictoryAndLoose(dataHangmanWeb *config.DataHangmanWeb) template.HTML {
	var message string = ""

	// Check if victory counter is enabled and there are victories or defeats
	if dataHangmanWeb.DataConfigHangman.VictoryCounter && (dataHangmanWeb.GameData.NbrVictory > 0 || dataHangmanWeb.GameData.NbrLoose > 0) {
		// Construct the message with the number of victories and defeats
		message = dataHangmanWeb.TranslationHangman.YourCurrentlyOn + strconv.Itoa(dataHangmanWeb.GameData.NbrVictory) + dataHangmanWeb.TranslationHangman.Victory + dataHangmanWeb.TranslationHangman.And + strconv.Itoa(dataHangmanWeb.GameData.NbrLoose) + dataHangmanWeb.TranslationHangman.Defeat
	}

	// Return the message wrapped in a <p> tag as an HTML-safe string
	return template.HTML("<p>" + message + "</p>")
}
