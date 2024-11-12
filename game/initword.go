package game

import (
	"hangman-web/config"
	"hangman-web/inserthtml"
	"html/template"

	"github.com/thedevrems/hangman/managegame"
)

// InitWord initializes the word to be guessed in the Hangman game based on the difficulty level.
// It uses configuration, translation, and file management data to select an appropriate word
// and reveal random letters to the player.
//
// Parameters:
// - dataError: pointer to the structure containing error information.
// - dataTranslation: pointer to the structure containing translation and game data (config.Structure_Translation).
// - dataFiles: pointer to the structure containing file information (configuration.Structure_files).
// - dataConfig: pointer to the structure containing game configuration parameters (configuration.Structure_configuration).
// - dataGame: pointer to the structure containing game-specific data (configuration.Structure_game).
// - nameDifficulty: string representing the difficulty level chosen by the player.
//
// Returns:
// - template.HTML: the formatted version of the partially revealed word, ready for display in the HTML interface.
func InitWord(dataHangmanWeb *config.DataHangmanWeb) template.HTML {

	// Selects an appropriate word based on the difficulty level using available parameters and translations.
	goodWord := managegame.ChooseWord(dataHangmanWeb.DataError, dataHangmanWeb.TranslationHangman, dataHangmanWeb.DataFiles, dataHangmanWeb.DataConfigHangman, dataHangmanWeb.DataGameHangman, dataHangmanWeb.GameData.NameDifficulty)

	// Stores the selected word in the game data structure for the interface.
	dataHangmanWeb.GameData.TargetWord = goodWord

	// Randomly reveals certain letters of the word to provide a hint to the player.
	revealword := managegame.RevealRandomLettersOfWord(dataHangmanWeb.TranslationHangman, goodWord, dataHangmanWeb.GameData.NameDifficulty)

	// Updates the current displayed word for the player with the revealed letters.
	dataHangmanWeb.GameData.CurrentWord = revealword

	// Formats the word so that it is ready for display in the HTML interface.
	formatWord := inserthtml.FormatWordForHTML(revealword)

	// Returns the formatted word as HTML content for display.
	return formatWord
}
