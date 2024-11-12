package game

import (
	"hangman-web/config"
	"hangman-web/inserthtml"

	"github.com/thedevrems/hangman/manageerror"
	"github.com/thedevrems/hangman/managegame"
)

func UseJoker(dataHangmanWeb *config.DataHangmanWeb) string {
	tempActualWord, charJoker := managegame.EnableJokers(dataHangmanWeb.DataGameHangman, dataHangmanWeb.GameData.CurrentWord, dataHangmanWeb.GameData.TargetWord, dataHangmanWeb.GameData.TabSelectedLetter)

	if tempActualWord == dataHangmanWeb.DataGameHangman.ErrorJokers {
		manageerror.PrintError(dataHangmanWeb.DataError, dataHangmanWeb.TranslationHangman.TittleError, dataHangmanWeb.TranslationHangman.JokerLimitMessage)
		return dataHangmanWeb.TranslationHangman.JokerLimitMessage
	} else {
		dataHangmanWeb.GameData.NbrJokers--
		dataHangmanWeb.GameData.CurrentWord = tempActualWord
		dataHangmanWeb.GameData.CurrentWordFormattedHTML = inserthtml.FormatWordForHTML(dataHangmanWeb.GameData.CurrentWord)
		dataHangmanWeb.GameData.TabSelectedLetter = append(dataHangmanWeb.GameData.TabSelectedLetter, charJoker)
		dataHangmanWeb.GameData.PrintLetterAndWord = inserthtml.PrintLetterAndWord(dataHangmanWeb)

		return dataHangmanWeb.TranslationHangman.JokerUsedMessage + charJoker
	}
}
