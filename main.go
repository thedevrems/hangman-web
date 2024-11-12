package main

import (
	"hangman-web/config"
	"hangman-web/manageserver"

	"github.com/thedevrems/hangman/configuration"
	"github.com/thedevrems/hangman/managefiles"
)

func main() {
	// Retrieve the server configuration data.
	// The ConfigServer() function loads the specific server configuration.
	dataConfigServer := config.ConfigServer()

	// Retrieve the configuration for the files used in the game.
	dataFiles := configuration.ConfigFiles()

	// Retrieve the configuration for errors, which may contain specific error messages for the game.
	dataError := configuration.ConfigError()

	// Load the default configuration for the Hangman game.
	dataConfigHangman := configuration.DefaultConfiguration()

	// Retrieve the specific configuration for the Hangman game
	dataGameHangman := configuration.ConfigGame()

	// Retrieve the configuration for the Hangman game for the web.
	dataGameHangmanWeb := config.ConfigDataGameHangmanWeb()

	// Declare variables for the Hangman game translations and for the web interface translations.
	var dataTranslationHangman *configuration.TranslationHangman
	var dataTranslationHangmanWeb *config.HangmanWebTranslations

	// Select which language translation to load based on the language set in the Hangman game configuration.
	if dataConfigHangman.Language == dataGameHangman.Frensh {
		// If the language is French, load the French translations.
		dataTranslationHangman = configuration.DataTranslationHangmanFr()
		dataTranslationHangmanWeb = config.DataTranslationHangmanWebFr()
	} else {
		// Otherwise, load the English translations.
		dataTranslationHangman = configuration.DataTranslationHangmanEn()
		dataTranslationHangmanWeb = config.DataTranslationHangmanWebEn()
	}

	// Initialize the main data structure for the Hangman Web interface, containing all configuration and translation data.
	dataHangmanWeb := &config.DataHangmanWeb{
		ServerConfig:           dataConfigServer,          // Server configuration.
		HangmanWebTranslations: dataTranslationHangmanWeb, // Web interface translations.
		GameData:               dataGameHangmanWeb,        // Hangman game data for the web.
		DataError:              dataError,                 // Error data for the game.
		DataGameHangman:        dataGameHangman,           // General Hangman game configuration.
		DataFiles:              dataFiles,                 // Files needed for the game (e.g., word lists).
		DataConfigHangman:      dataConfigHangman,         // General Hangman configuration.
		TranslationHangman:     dataTranslationHangman,    // Hangman game translations.
	}

	// If the game's difficulty configuration is empty, set it to the default value (Normal).
	if dataHangmanWeb.GameData.NameDifficulty == "" {
		dataHangmanWeb.GameData.NameDifficulty = dataHangmanWeb.TranslationHangman.Normal
		if dataHangmanWeb.DataConfigHangman.Language == "en" {
			dataHangmanWeb.GameData.NameFile = dataHangmanWeb.DataFiles.NameFilesConfigEnglishWordsDefault
		} else {
			dataHangmanWeb.GameData.NameFile = dataHangmanWeb.DataFiles.NameFilesConfigFrenchWordsDefault
		}
	}

	// Manage the files needed for the game to function properly.
	managefiles.CaseManagementFiles(dataFiles, dataConfigHangman, dataGameHangman, dataError, dataTranslationHangman)

	// Create the web server to host the Hangman game using the initialized configuration and translation data.
	manageserver.CreateServer(dataHangmanWeb)
}
