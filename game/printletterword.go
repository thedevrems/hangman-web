package game

import (
	"hangman-web/config"
	"html/template"
)

// PrintLetterAndWord generates an HTML string that displays the list of already selected letters and words.
//
// Parameters:
// - dataHangmanWeb (*config.DataHangmanWeb): The game state, which holds the selected letters and words.
//
// Returns:
// - template.HTML: An HTML-safe string that contains two sections:
//  1. A paragraph displaying the letters already used.
//  2. A paragraph displaying the words already guessed.
func PrintLetterAndWord(dataHangmanWeb *config.DataHangmanWeb) template.HTML {
	var data string

	// If there are selected letters, display them
	if len(dataHangmanWeb.GameData.TabSelectedLetter) > 0 {
		data += "<p>" + dataHangmanWeb.TranslationHangman.LettersAlreadyUsed
		for _, letter := range dataHangmanWeb.GameData.TabSelectedLetter {
			data += letter + " "
		}
		// Remove the extra space at the end
		data = data[:len(data)-1]
		data += "</p>"
	}

	// If there are selected words, display them
	if len(dataHangmanWeb.GameData.TabSelectedWord) > 0 {
		data += "<p>" + dataHangmanWeb.TranslationHangman.WordsAlreadyUsed
		for _, word := range dataHangmanWeb.GameData.TabSelectedWord {
			data += word + " "
		}
		// Remove the extra space at the end
		data = data[:len(data)-1]
		data += "</p>"
	}

	// Return the HTML-safe string for rendering
	return template.HTML(data)
}
