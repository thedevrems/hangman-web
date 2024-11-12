package game

import "hangman-web/config"

// Resets game values to their initial state
func ResetGame(data *config.DataHangmanWeb) {
	data.GameData.TabSelectedLetter = []string{}
	data.GameData.TabSelectedWord = []string{}
	data.GameData.PrintLetterAndWord = PrintLetterAndWord(data)
	data.GameData.StageHangman = 0
	data.GameData.PrintInfo = PrintInfo(data, "")
	data.GameData.CurrentWordFormattedHTML = InitWord(data)
	data.GameData.StageHangmanFormattedHTML = FormatDrawHangmanForHTML(data)
	data.GameData.StateInputField = true
	data.GameData.DataInputField = AddInputField(data, data.GameData.StateInputField)
	data.GameData.StateInputFieldForAddWord = false
	data.GameData.DataInputFieldForAddWord = AddInputFieldForAddWord(data, data.GameData.StateInputFieldForAddWord)

	if data.DataConfigHangman.VictoryCounter {
		if data.GameData.CurrentWord != data.GameData.TargetWord {
			data.GameData.NbrLoose++
		} else {
			data.GameData.NbrVictory++
		}
		data.GameData.PrintNbrVictoryAndLoose = PrintNbrVictoryAndLoose(data)
	}
}
