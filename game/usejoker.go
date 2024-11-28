package game

import (
	"hangman-web/config"

	"github.com/thedevrems/hangman/manageerror"
	"github.com/thedevrems/hangman/managegame"
)

// UseJoker allows the player to use a joker if they have any remaining.
// A joker helps the player by revealing a letter from the target word without guessing it.
// The number of available jokers is decremented each time a joker is used. If no jokers are left,
// the function returns a message indicating that the player can't use a joker.
//
// Parameters:
// - dataHangmanWeb (*config.DataHangmanWeb): The game state object that holds all the game data and configurations.
//
// Returns:
//   - string: A message indicating the outcome of using the joker.
//     It can either be a success message (showing the letter revealed) or an error message (if no jokers are left or if there is an issue using a joker).
func UseJoker(dataHangmanWeb *config.DataHangmanWeb) string {
	// Check if the player has jokers remaining
	if dataHangmanWeb.GameData.NbrJokers > 0 {
		// Attempt to use a joker and get the updated word and the letter revealed by the joker
		tempActualWord, charJoker := managegame.EnableJokers(dataHangmanWeb.DataGameHangman, dataHangmanWeb.GameData.CurrentWord, dataHangmanWeb.GameData.TargetWord, dataHangmanWeb.GameData.TabSelectedLetter)

		// Check if the joker usage resulted in an error
		if tempActualWord == dataHangmanWeb.DataGameHangman.ErrorJokers {
			// If there was an error (e.g., joker can't be used), print the error and return the message
			manageerror.PrintError(dataHangmanWeb.DataError, dataHangmanWeb.TranslationHangman.TittleError, dataHangmanWeb.TranslationHangman.JokerLimitMessage)
			return dataHangmanWeb.TranslationHangman.JokerLimitMessage
		} else {
			// Successfully used a joker: decrement the joker count, update the current word, and show the new state
			dataHangmanWeb.GameData.NbrJokers--                                                                       // Decrease the number of jokers available
			dataHangmanWeb.GameData.CurrentWord = tempActualWord                                                      // Update the current word with the revealed letter
			dataHangmanWeb.GameData.CurrentWordFormattedHTML = FormatWordForHTML(dataHangmanWeb.GameData.CurrentWord) // Format the word for display
			dataHangmanWeb.GameData.TabSelectedLetter = append(dataHangmanWeb.GameData.TabSelectedLetter, charJoker)  // Add the joker's letter to the selected letters
			dataHangmanWeb.GameData.PrintLetterAndWord = PrintLetterAndWord(dataHangmanWeb)                           // Update the displayed selected letters and words

			// Return a message indicating the joker was used successfully
			return dataHangmanWeb.TranslationHangman.JokerUsedMessage + charJoker
		}
	} else {
		// If no jokers are left, return a message saying the player can't use a joker
		return dataHangmanWeb.HangmanWebTranslations.CantUseJoker
	}
}
