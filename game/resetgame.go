package game

import "hangman-web/config"

// ResetGame resets the game state to its initial values, including clearing selected letters, words, and resetting the hangman stage.
// It also updates various game-related attributes like victory/defeat counters, input fields, and the formatted word.
//
// Parameters:
// - data (*config.DataHangmanWeb): The game state object that contains all the game data and configuration settings.
//
// Function Effects:
// - Resets the selected letters and words to empty slices.
// - Resets the hangman stage to 0.
// - Clears and reinitializes various game attributes such as the current word, victory/defeat counters, input fields, and more.
func ResetGame(data *config.DataHangmanWeb) {
	// Reset the list of selected letters and words
	data.GameData.TabSelectedLetter = []string{}
	data.GameData.TabSelectedWord = []string{}

	// Update the display of the selected letters and words
	data.GameData.PrintLetterAndWord = PrintLetterAndWord(data)

	// Reset the hangman stage to 0 (no parts drawn)
	data.GameData.StageHangman = 0

	// Reset any information displayed (i.e., empty message)
	data.GameData.PrintInfo = PrintInfo(data, "")

	// Initialize the current word in the game
	data.GameData.CurrentWordFormattedHTML = InitWord(data)

	// Recreate the HTML representation of the hangman (no parts drawn)
	data.GameData.StageHangmanFormattedHTML = FormatDrawHangmanForHTML(data)

	// Set the state for the input field for the next guess
	data.GameData.StateInputField = true
	data.GameData.DataInputField = AddInputField(data, data.GameData.StateInputField)

	// Disable the input field for adding a word after the game
	data.GameData.StateInputFieldForAddWord = false
	data.GameData.DataInputFieldForAddWord = AddInputFieldForAddWord(data, data.GameData.StateInputFieldForAddWord)

	// If victory counter is enabled, update the counters based on the result of the last game
	if data.DataConfigHangman.VictoryCounter {
		// Check if the current word matches the target word and update the victory or loss counter
		if data.GameData.CurrentWord != data.GameData.TargetWord {
			data.GameData.NbrLoose++ // Increment loss counter if the word does not match
		} else {
			data.GameData.NbrVictory++ // Increment victory counter if the word matches
		}
		// Update the display of victories and losses
		data.GameData.PrintNbrVictoryAndLoose = PrintNbrVictoryAndLoose(data)
	}
}
