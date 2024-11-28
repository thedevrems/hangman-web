package config

import (
	"html/template"

	"github.com/thedevrems/hangman/configuration"
)

// ServerConfig holds configuration settings specific to the server setup.
type ServerConfig struct {
	Path       string // Root path for the server.
	IndexPage  string // Path or name of the index page.
	ConfigPage string // Path or name of the configuration page.
	GamePage   string // Path or name of the game page.
	URL        string // URL where the server will be hosted.
	Port       string // Port number on which the server will listen.
	Slash      string // Character or sequence representing a URL slash.
}

// HangmanWebTranslations holds translations specific to the Hangman web interface.
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

// GameData holds the current state and data for an ongoing game session.
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

// DataHangmanWeb aggregates all configuration, translation, and game data structures for the Hangman web interface.
type DataHangmanWeb struct {
	*ServerConfig                     // Server configuration data.
	*HangmanWebTranslations           // Translations for web interface elements.
	*GameData                         // Game session data and current state.
	*configuration.DataError          // Error messages and error configuration data.
	*configuration.DataGameHangman    // Main configuration for the Hangman game.
	*configuration.DataFiles          // File data for word lists and other resources.
	*configuration.DataConfigHangman  // Default Hangman game settings.
	*configuration.TranslationHangman // Game-specific translations.
}
