package game

import (
	"hangman-web/config"
	"html/template"
	"strconv"
)

func PrintNbrJokers(dataHangmanWeb *config.DataHangmanWeb) template.HTML {
	var data string
	if dataHangmanWeb.GameData.NbrJokers > 0 && dataHangmanWeb.DataConfigHangman.EnableJokers {
		stringNbrJokers := strconv.Itoa(dataHangmanWeb.GameData.NbrJokers)
		data = "<p>" + dataHangmanWeb.TranslationHangman.JokersCountMessage + stringNbrJokers + "</p>"
		return template.HTML(data)
	} else {
		return template.HTML("<p></p>")
	}

}
