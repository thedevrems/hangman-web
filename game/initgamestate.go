package game

import "hangman-web/config"

// InitGameState initializes the game state with default values and generates various components
// for the game interface. This function populates the `data.GameData` structure with the HTML
// elements and values required for the game to function properly.
//
// Parameters:
// - data (*config.DataHangmanWeb): A reference to the game data structure that holds
//   the current game state, including the game's HTML components, translations, and configuration.
//
// This function does not return any value. It modifies the `data.GameData` structure in place.
func InitGameState(data *config.DataHangmanWeb) {
	// Initialize the display of the number of victories and losses
	data.GameData.PrintNbrVictoryAndLoose = PrintNbrVictoryAndLoose(data)

	// Initialize the display of the number of jokers
	data.GameData.PrintNbrJokers = PrintNbrJokers(data)

	// Set the initial state of the input field (enabled)
	data.GameData.StateInputField = true

	// Add the input field for guessing the word or letter based on the state of data.GameData.StateInputField
	data.GameData.DataInputField = AddInputField(data, data.GameData.StateInputField)

	// Add the input field for adding a new word after the game based on the state of data.GameData.StateInputFieldForAddWord
	data.GameData.DataInputFieldForAddWord = AddInputFieldForAddWord(data, data.GameData.StateInputFieldForAddWord)

	// Initialize the display of the current word (empty or partially revealed)
	data.GameData.CurrentWordFormattedHTML = InitWord(data)

	// Generate the Hangman figure based on the current game state
	data.GameData.StageHangmanFormattedHTML = FormatDrawHangmanForHTML(data)

	// Add the "Use Joker" button to the interface based on the current configuration
	data.GameData.AddButtonForUseJoker = AddButtonJoker(data)

	// Add the form for inserting a new word after the game ends
	data.GameData.AddFormForTheExtensionAddWordAfterGame = ExtensionInsertWord(data)

	// Add the form to switch between English and French language
	data.GameData.AddFormForTheExtensionLanguage = ExtensionSwitchLanguage(data)

	// Add the form to enable difficulty settings
	data.GameData.AddFormForTheExtensionEnableDifficulty = ExtensionDifficulty(data)

	// Add the form to enable jokers in the game
	data.GameData.AddFormForTheExtensionEnableJokers = ExtensionJokers(data)

	// Add the form to enable the victory counter extension
	data.GameData.AddFormForTheExtensionVictoryCounter = ExtensionCountVictoryAndLoose(data)
}
