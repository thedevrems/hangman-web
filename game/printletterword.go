package game

import (
	"fmt"
	"hangman-web/config"
	"html/template"
)

func PrintLetterAndWord(dataHangmanWeb *config.DataHangmanWeb) template.HTML {
	var data string

	fmt.Println(dataHangmanWeb.GameData.TabSelectedLetter)
	if len(dataHangmanWeb.GameData.TabSelectedLetter) > 0 {
		data += "<p>" + dataHangmanWeb.TranslationHangman.LettersAlreadyUsed
		for _, letter := range dataHangmanWeb.GameData.TabSelectedLetter {
			data += letter + " "
		}
		// Supprimer l'espace supplémentaire à la fin
		data = data[:len(data)-1]
		data += "</p>"
	}

	if len(dataHangmanWeb.GameData.TabSelectedWord) > 0 {
		data += "<p>" + dataHangmanWeb.TranslationHangman.WordsAlreadyUsed
		for _, word := range dataHangmanWeb.GameData.TabSelectedWord {
			data += word + " "
		}
		// Supprimer l'espace supplémentaire à la fin
		data = data[:len(data)-1]
		data += "</p>"
	}

	return template.HTML(data)
}
