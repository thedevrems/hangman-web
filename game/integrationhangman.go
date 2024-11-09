package game

import (
	"fmt"
	"hangman-web/config"

	"github.com/thedevrems/hangman/manageerror"
	"github.com/thedevrems/hangman/managegame"
)

// Cette fonction aura pour but d'intéragir avec le hangman normal et le site
func IntegrationHangman(dataHangmanWeb *config.DataHangmanWeb, selectedLetter string) (data string) {
	var youWin bool = false

	if dataHangmanWeb.TranslationHangman.VeryEasy == dataHangmanWeb.GameData.NameDifficulty {
		dataHangmanWeb.GameData.NumOfPossibility = 15
	}

	if dataHangmanWeb.GameData.StageHangman == dataHangmanWeb.GameData.NumOfPossibility {
		dataHangmanWeb.GameData.DataInputField = AddInputField(dataHangmanWeb, false)
		return "José est mort, merci de cliquer sur le button suivant :"
	}

	runeLetter := []rune(selectedLetter)

	if len(runeLetter) == 0 {

		manageerror.PrintError(dataHangmanWeb.DataError, dataHangmanWeb.TranslationHangman.TittleError, dataHangmanWeb.TranslationHangman.MissingCharacterError)

		if dataHangmanWeb.GameData.NameDifficulty == dataHangmanWeb.TranslationHangman.Hacker {
			dataHangmanWeb.GameData.StageHangman++
			dataHangmanWeb.GameData.StageHangmanFormattedHTML = FormatDrawHangmanForHTML(dataHangmanWeb)
		}

		return dataHangmanWeb.TranslationHangman.MissingCharacterError

	} else if len(runeLetter) == 1 {
		if !managegame.IsLetter(runeLetter[0]) {

			manageerror.PrintError(dataHangmanWeb.DataError, dataHangmanWeb.TranslationHangman.TittleError, dataHangmanWeb.TranslationHangman.NotALetterError)

			if dataHangmanWeb.GameData.NameDifficulty == dataHangmanWeb.TranslationHangman.Hacker {
				dataHangmanWeb.GameData.StageHangman++
				dataHangmanWeb.GameData.StageHangmanFormattedHTML = FormatDrawHangmanForHTML(dataHangmanWeb)
			}

			return dataHangmanWeb.TranslationHangman.NotALetterError

		} else if managegame.AlreadySelected(dataHangmanWeb.GameData.TabSelectedLetter, managegame.GetLetterConversion(runeLetter[0])) {
			manageerror.PrintError(dataHangmanWeb.DataError, dataHangmanWeb.TranslationHangman.TittleError, dataHangmanWeb.TranslationHangman.AlreadySelectedLetterError)

			if dataHangmanWeb.GameData.NameDifficulty == dataHangmanWeb.TranslationHangman.Hacker {
				dataHangmanWeb.GameData.StageHangman++
				dataHangmanWeb.GameData.StageHangmanFormattedHTML = FormatDrawHangmanForHTML(dataHangmanWeb)
			}

			return dataHangmanWeb.TranslationHangman.AlreadySelectedLetterError

		} else {
			letterConversion := managegame.GetLetterConversion(runeLetter[0])
			dataHangmanWeb.GameData.TabSelectedLetter = append(dataHangmanWeb.GameData.TabSelectedLetter, letterConversion)
			dataHangmanWeb.GameData.PrintLetterAndWord = PrintLetterAndWord(dataHangmanWeb)

			if managegame.TestLetterInWord(dataHangmanWeb.GameData.TargetWord, letterConversion) {
				dataHangmanWeb.GameData.CurrentWord = managegame.RevealWord(dataHangmanWeb.GameData.TargetWord, dataHangmanWeb.GameData.CurrentWord, letterConversion)
				if dataHangmanWeb.GameData.CurrentWord == dataHangmanWeb.GameData.TargetWord {
					youWin = true
				}
				dataHangmanWeb.GameData.CurrentWordFormattedHTML = FormatWordForHTML(dataHangmanWeb.GameData.CurrentWord)
			} else {
				dataHangmanWeb.GameData.StageHangman++
				dataHangmanWeb.GameData.StageHangmanFormattedHTML = FormatDrawHangmanForHTML(dataHangmanWeb)
			}
		}

	} else if len(selectedLetter) == len(dataHangmanWeb.GameData.TargetWord) {
		if managegame.AlreadySelected(dataHangmanWeb.GameData.TabSelectedWord, selectedLetter) {

			manageerror.PrintError(dataHangmanWeb.DataError, dataHangmanWeb.TranslationHangman.TittleError, dataHangmanWeb.TranslationHangman.AlreadySelectedWordError)

			if dataHangmanWeb.GameData.NameDifficulty == dataHangmanWeb.TranslationHangman.Hacker {
				dataHangmanWeb.GameData.StageHangman++
				dataHangmanWeb.GameData.StageHangmanFormattedHTML = FormatDrawHangmanForHTML(dataHangmanWeb)
			}

			return dataHangmanWeb.TranslationHangman.AlreadySelectedWordError

		} else {

			dataHangmanWeb.GameData.TabSelectedWord = append(dataHangmanWeb.GameData.TabSelectedWord, selectedLetter)
			dataHangmanWeb.GameData.PrintLetterAndWord = PrintLetterAndWord(dataHangmanWeb)
			if selectedLetter == dataHangmanWeb.GameData.TargetWord {
				youWin = true
			} else {
				dataHangmanWeb.GameData.StageHangman++
				dataHangmanWeb.GameData.StageHangmanFormattedHTML = FormatDrawHangmanForHTML(dataHangmanWeb)
			}
		}
	} else {

		manageerror.PrintError(dataHangmanWeb.DataError, dataHangmanWeb.TranslationHangman.TittleError, dataHangmanWeb.TranslationHangman.SingleLetterOrWordError)

		if dataHangmanWeb.GameData.NameDifficulty == dataHangmanWeb.TranslationHangman.Hacker {
			dataHangmanWeb.GameData.StageHangman++
			dataHangmanWeb.GameData.StageHangmanFormattedHTML = FormatDrawHangmanForHTML(dataHangmanWeb)
		}

		return dataHangmanWeb.TranslationHangman.SingleLetterOrWordError
	}

	if youWin {
		fmt.Println(dataHangmanWeb.TranslationHangman.Congratulations)
		fmt.Println(dataHangmanWeb.TranslationHangman.GameWonMessage + dataHangmanWeb.GameData.TargetWord)
		dataHangmanWeb.GameData.DataInputField = AddInputField(dataHangmanWeb, false)
		return dataHangmanWeb.TranslationHangman.Congratulations + dataHangmanWeb.TranslationHangman.GameWonMessage + dataHangmanWeb.GameData.TargetWord
	}

	return ""
}
