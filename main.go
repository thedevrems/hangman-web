package main

import (
	"hangman-web/config"
	"hangman-web/manageserver"

	"github.com/thedevrems/hangman/configuration"
	"github.com/thedevrems/hangman/managefiles"
)

func main() {
	// Récupération des données de la configuration du serveur (voir config/configserver.go)
	dataConfigServer := config.ConfigServer()
	dataFiles := configuration.ConfigFiles()
	dataError := configuration.ConfigError()
	dataConfigHangman := configuration.DefaultConfiguration()
	dataGameHangman := configuration.ConfigGame()
	dataGameHangmanWeb := config.ConfigDataGameHangmanWeb()

	var dataTranslationHangman *configuration.TranslationHangman
	var dataTranslationHangmanWeb *config.HangmanWebTranslations

	// Charger les traductions en fonction de la langue
	if dataConfigHangman.Language == dataGameHangman.Frensh {
		dataTranslationHangman = configuration.DataTranslationHangmanFr()
		dataTranslationHangmanWeb = config.DataTranslationHangmanWebFr()
	} else {
		dataTranslationHangman = configuration.DataTranslationHangmanEn()
		dataTranslationHangmanWeb = config.DataTranslationHangmanWebEn()
	}

	// Initialiser la structure de traduction principale
	dataHangmanWeb := &config.DataHangmanWeb{
		ServerConfig:           dataConfigServer,
		HangmanWebTranslations: dataTranslationHangmanWeb,
		GameData:               dataGameHangmanWeb,
		DataError:              dataError,
		DataGameHangman:        dataGameHangman,
		DataFiles:              dataFiles,
		DataConfigHangman:      dataConfigHangman,
		TranslationHangman:     dataTranslationHangman,
	}

	if dataHangmanWeb.GameData.NameDifficulty == "" {
		dataHangmanWeb.GameData.NameDifficulty = dataHangmanWeb.TranslationHangman.Normal
	}

	managefiles.CaseManagementFiles(dataFiles, dataConfigHangman, dataGameHangman, dataError, dataTranslationHangman)

	manageserver.CreateServer(dataHangmanWeb)
}
