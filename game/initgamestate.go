package game

import "hangman-web/config"

// Initializes the game state with default values
func InitGameState(data *config.DataHangmanWeb) {
	data.GameData.PrintNbrVictoryAndLoose = PrintNbrVictoryAndLoose(data)
	data.GameData.PrintNbrJokers = PrintNbrJokers(data)
	data.GameData.StateInputField = true
	data.GameData.DataInputField = AddInputField(data, data.GameData.StateInputField)
	data.GameData.DataInputFieldForAddWord = AddInputField(data, data.GameData.StateInputFieldForAddWord)
	data.GameData.CurrentWordFormattedHTML = InitWord(data)
	data.GameData.StageHangmanFormattedHTML = FormatDrawHangmanForHTML(data)
	data.GameData.AddButtonForUseJoker = AddButtonJoker(data)
	data.GameData.AddFormForTheExtensionAddWordAfterGame = ExtensionInsertWord(data)
	data.GameData.AddFormForTheExtensionLanguage = ExtensionSwitchLanguage(data)
	data.GameData.AddFormForTheExtensionEnableDifficulty = ExtensionDifficulty(data)
	data.GameData.AddFormForTheExtensionEnableJokers = ExtensionJokers(data)
	data.GameData.AddFormForTheExtensionVictoryCounter = ExtensionCountVictoryAndLoose(data)
}
