package config

import (
	"html/template"

	"github.com/thedevrems/hangman/configuration"
)

// ServerConfig représente la configuration du serveur
type ServerConfig struct {
	Path       string
	IndexPage  string
	ConfigPage string
	GamePage   string
	URL        string
	Port       string
	Slash      string
}

// HangmanWebTranslations contient les traductions spécifiques à l'interface web du jeu du pendu
type HangmanWebTranslations struct {
	Title                           string
	Description                     string
	StartButtonText                 string
	NavLogoTitle                    string
	NavAbout                        string
	NavGame                         string
	NavConfig                       string
	TitleExtensionAddWordAfterGame  string
	TitleExtensionCountVictory      string
	TitleExtensionJokers            string
	TitleExtensionLanguage          string
	TitleExtensionEnableDifficulty  string
	ExtensionDifficultyLevel1       string
	ExtensionDifficultyLevel2       string
	ExtensionDifficultyLevel3       string
	ExtensionDifficultyLevel4       string
	ExtensionDifficultyLevel5       string
	ExtensionDifficultyLevel6       string
	CantUseJoker                    string
	RestartGameButton               string
	PlaceHolderInputField           string
	UseJoker                        string
	PlaceHolderInputFieldForAddWord string
}

// GameData contient les données spécifiques au jeu pour une session
type GameData struct {
	TargetWord                string
	CurrentWord               string
	CurrentWordFormattedHTML  template.HTML
	DifficultyLevel           string
	NbrVictory                int
	NbrLoose                  int
	NbrJokers                 int
	NameDifficulty            string
	NameFile                  string
	IdDifficulty              int
	StageHangman              int
	StageHangmanFormattedHTML template.HTML
	NumOfPossibility          int
	TabSelectedLetter         []string
	TabSelectedWord           []string
	PrintNbrVictoryAndLoose   template.HTML
	PrintInfo                 template.HTML
	PrintNbrJokers            template.HTML
	PrintLetterAndWord        template.HTML
	DataInputField            template.HTML
	StateInputField           bool
	StateInputFieldForAddWord bool
	AddButtonForUseJoker      template.HTML
	DataInputFieldForAddWord  template.HTML

	AddFormForTheExtensionLanguage         template.HTML
	AddFormForTheExtensionVictoryCounter   template.HTML
	AddFormForTheExtensionEnableDifficulty template.HTML
	AddFormForTheExtensionAddWordAfterGame template.HTML
	AddFormForTheExtensionEnableJokers     template.HTML
}

// DataHangmanWeb regroupe toute les structure
type DataHangmanWeb struct {
	*ServerConfig
	*HangmanWebTranslations
	*GameData
	*configuration.DataError
	*configuration.DataGameHangman
	*configuration.DataFiles
	*configuration.DataConfigHangman
	*configuration.TranslationHangman
}
