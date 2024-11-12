package game

import (
	"hangman-web/config"
	"html/template"
	"strconv"
)

func PrintNbrVictoryAndLoose(dataHangmanWeb *config.DataHangmanWeb) template.HTML {
	var message string = ""

	if dataHangmanWeb.DataConfigHangman.VictoryCounter && (dataHangmanWeb.GameData.NbrVictory > 0 || dataHangmanWeb.GameData.NbrLoose > 0) {
		message = dataHangmanWeb.TranslationHangman.YourCurrentlyOn + strconv.Itoa(dataHangmanWeb.GameData.NbrVictory) + dataHangmanWeb.TranslationHangman.Victory + dataHangmanWeb.TranslationHangman.And + strconv.Itoa(dataHangmanWeb.GameData.NbrLoose) + dataHangmanWeb.TranslationHangman.Defeat
	}

	return template.HTML("<p>" + message + "</p>")
}
