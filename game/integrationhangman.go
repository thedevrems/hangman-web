package game

import (
	"fmt"
	"hangman-web/config"

	"github.com/thedevrems/hangman/manageerror"
	"github.com/thedevrems/hangman/managegame"
)

// IntegrationHangman handles the main logic of the Hangman game, including checking for valid input,
// updating the game state based on the user's guess, and determining when the game ends.
//
// Parameters:
// - dataHangmanWeb (*config.DataHangmanWeb): the game data structure, which holds all the game state and configurations.
// - selectedLetter (string): the user's input, which could be a single letter or a full word.
//
// Returns:
// - string: a message to be displayed to the user, indicating the result of their guess or any error encountered.
func IntegrationHangman(dataHangmanWeb *config.DataHangmanWeb, selectedLetter string) string {
	// Set the number of attempts based on difficulty level (VeryEasy level gives 15 attempts)
	if dataHangmanWeb.TranslationHangman.VeryEasy == dataHangmanWeb.GameData.NameDifficulty {
		dataHangmanWeb.GameData.NumOfPossibility = 15
	}

	// Check if the game is over due to reaching the maximum number of attempts
	if dataHangmanWeb.GameData.StageHangman >= dataHangmanWeb.GameData.NumOfPossibility {
		IncrementStage(dataHangmanWeb)
		EndGame(dataHangmanWeb, false)
		return dataHangmanWeb.TranslationHangman.MessageJoseDead
	}

	// Handle input errors and game logic
	runeLetter := []rune(selectedLetter)
	if len(runeLetter) == 0 {
		return HandleInputError(dataHangmanWeb, dataHangmanWeb.TranslationHangman.MissingCharacterError)
	}

	// Handle single letter guesses
	if len(runeLetter) == 1 {
		return HandleSingleLetter(dataHangmanWeb, runeLetter[0])
	}

	// Handle full word guesses
	if len(selectedLetter) == len(dataHangmanWeb.GameData.TargetWord) {
		return HandleWordGuess(dataHangmanWeb, selectedLetter)
	}

	// If input is invalid (not a letter or word)
	return HandleInputError(dataHangmanWeb, dataHangmanWeb.TranslationHangman.SingleLetterOrWordError)
}

// EndGame stops the game, either marking it as a win or loss, and updates the UI accordingly.
// It also provides the option to add a new word to the game after a win.
//
// Parameters:
// - dataHangmanWeb (*config.DataHangmanWeb): the game data structure.
// - youWin (bool): whether the player won the game or not.
func EndGame(dataHangmanWeb *config.DataHangmanWeb, youWin bool) {
	dataHangmanWeb.GameData.StateInputField = false
	dataHangmanWeb.GameData.DataInputField = AddInputField(dataHangmanWeb, dataHangmanWeb.GameData.StateInputField)
	// If the player won, display the win message and possibly offer to add a new word
	if youWin {
		fmt.Println(dataHangmanWeb.TranslationHangman.Congratulations)
		fmt.Println(dataHangmanWeb.TranslationHangman.GameWonMessage + dataHangmanWeb.GameData.TargetWord)
		if dataHangmanWeb.DataConfigHangman.AddWordAfterGame {
			dataHangmanWeb.GameData.StateInputFieldForAddWord = true
			dataHangmanWeb.GameData.DataInputFieldForAddWord = AddInputFieldForAddWord(dataHangmanWeb, true)
		}
	}
}

// HandleInputError logs the error and returns the appropriate error message.
//
// Parameters:
// - dataHangmanWeb (*config.DataHangmanWeb): the game data structure.
// - errorMsg (string): the error message to be returned to the user.
//
// Returns:
// - string: the error message to display to the user.
func HandleInputError(dataHangmanWeb *config.DataHangmanWeb, errorMsg string) string {
	manageerror.PrintError(dataHangmanWeb.DataError, dataHangmanWeb.TranslationHangman.TittleError, errorMsg)

	// If the player is on the "Hacker" difficulty, automatically increment the Hangman stage on error
	if dataHangmanWeb.GameData.NameDifficulty == dataHangmanWeb.TranslationHangman.Hacker {
		IncrementStage(dataHangmanWeb)
	}
	return errorMsg
}

// HandleSingleLetter processes the user's guess when it's a single letter.
//
// Parameters:
// - dataHangmanWeb (*config.DataHangmanWeb): the game data structure.
// - letter (rune): the guessed letter.
//
// Returns:
// - string: an error message or an empty string if the guess was valid.
func HandleSingleLetter(dataHangmanWeb *config.DataHangmanWeb, letter rune) string {
	// Check if the letter is a valid character
	if !managegame.IsLetter(letter) {
		return HandleInputError(dataHangmanWeb, dataHangmanWeb.TranslationHangman.NotALetterError)
	}

	// Convert letter and check if it has been selected before
	letterConversion := managegame.GetLetterConversion(letter)
	if managegame.AlreadySelected(dataHangmanWeb.GameData.TabSelectedLetter, letterConversion) {
		return HandleInputError(dataHangmanWeb, dataHangmanWeb.TranslationHangman.AlreadySelectedLetterError)
	}

	// Add the selected letter to the list of chosen letters
	dataHangmanWeb.GameData.TabSelectedLetter = append(dataHangmanWeb.GameData.TabSelectedLetter, letterConversion)
	dataHangmanWeb.GameData.PrintLetterAndWord = PrintLetterAndWord(dataHangmanWeb)

	// If the letter is in the target word, reveal it in the current word
	if managegame.TestLetterInWord(dataHangmanWeb.GameData.TargetWord, letterConversion) {
		dataHangmanWeb.GameData.CurrentWord = managegame.RevealWord(dataHangmanWeb.GameData.TargetWord, dataHangmanWeb.GameData.CurrentWord, letterConversion)
		dataHangmanWeb.GameData.CurrentWordFormattedHTML = FormatWordForHTML(dataHangmanWeb.GameData.CurrentWord)

		// If the word is fully revealed, the player wins
		if dataHangmanWeb.GameData.CurrentWord == dataHangmanWeb.GameData.TargetWord {
			EndGame(dataHangmanWeb, true)
			return dataHangmanWeb.TranslationHangman.Congratulations + dataHangmanWeb.TranslationHangman.GameWonMessage + dataHangmanWeb.GameData.TargetWord
		}
	} else {
		// If the letter is incorrect, increment the Hangman stage
		IncrementStage(dataHangmanWeb)
	}
	return ""
}

// HandleWordGuess processes the user's guess when they try to guess the full word.
//
// Parameters:
// - dataHangmanWeb (*config.DataHangmanWeb): the game data structure.
// - guessedWord (string): the guessed word.
//
// Returns:
// - string: an error message or an empty string if the guess was valid.
func HandleWordGuess(dataHangmanWeb *config.DataHangmanWeb, guessedWord string) string {
	// Check if the word has been guessed before
	if managegame.AlreadySelected(dataHangmanWeb.GameData.TabSelectedWord, guessedWord) {
		return HandleInputError(dataHangmanWeb, dataHangmanWeb.TranslationHangman.AlreadySelectedWordError)
	}

	// Add the guessed word to the list of selected words
	dataHangmanWeb.GameData.TabSelectedWord = append(dataHangmanWeb.GameData.TabSelectedWord, guessedWord)
	dataHangmanWeb.GameData.PrintLetterAndWord = PrintLetterAndWord(dataHangmanWeb)

	// If the guessed word is correct, the player wins
	if guessedWord == dataHangmanWeb.GameData.TargetWord {
		EndGame(dataHangmanWeb, true)
		return dataHangmanWeb.TranslationHangman.Congratulations + dataHangmanWeb.TranslationHangman.GameWonMessage + dataHangmanWeb.GameData.TargetWord
	}

	// If the guessed word is incorrect, increment the Hangman stage
	IncrementStage(dataHangmanWeb)
	return ""
}

// IncrementStage increments the stage of the Hangman figure when the player makes an incorrect guess.
//
// Parameters:
// - dataHangmanWeb (*config.DataHangmanWeb): the game data structure.
func IncrementStage(dataHangmanWeb *config.DataHangmanWeb) {
	// Increment the Hangman stage and update the HTML representation
	dataHangmanWeb.GameData.StageHangman++
	dataHangmanWeb.GameData.StageHangmanFormattedHTML = FormatDrawHangmanForHTML(dataHangmanWeb)
}
