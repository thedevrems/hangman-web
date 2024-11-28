package manageserver

import (
	"hangman-web/config"
	"hangman-web/game"
	"net/http"

	"github.com/thedevrems/hangman/configuration"
	"github.com/thedevrems/hangman/managefiles"
)

// HandleSwitchLanguage returns an HTTP handler function that handles switching the language extension in the Hangman game.
// When a POST request is made with the selected language (either "EN" for English or "FR" for French), the handler updates the game configuration based on the chosen language.
// The handler adjusts various game settings, including translations, difficulty file mappings, and additional configurations, and redirects the user to the configuration page.
//
// Parameters:
// - data (*config.DataHangmanWeb): The game data containing configuration, game state, and translation data.
//
// Returns:
// - http.HandlerFunc: An HTTP handler function that processes the language change in the Hangman game.
func HandleSwitchLanguage(data *config.DataHangmanWeb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Ensure the request method is POST
		if r.Method == "POST" {
			// Retrieve the selected language from the form (either "EN" for English or "FR" for French)
			language := r.FormValue("switchLanguage")

			// Define a variable to hold the file map based on the language
			var fileMap map[int]string

			// Switch language settings based on the selected language
			if language == "EN" {
				// Set game language to English
				data.DataConfigHangman.Language = "en"
				data.TranslationHangman = configuration.DataTranslationHangmanEn() // Set English translations for the game
				data.HangmanWebTranslations = config.DataTranslationHangmanWebEn() // Set English web translations

				// Define the file map for English difficulty levels
				fileMap = map[int]string{
					1: data.DataFiles.NameFilesConfigEnglishWordsVeryEasy,
					2: data.DataFiles.NameFilesConfigEnglishWordsEasy,
					3: data.DataFiles.NameFilesConfigEnglishWordsDefault,
					4: data.DataFiles.NameFilesConfigEnglishWordsHard,
					5: data.DataFiles.NameFilesConfigEnglishWordsExpert,
					6: data.DataFiles.NameFilesConfigEnglishWordsHacker,
				}
			} else {
				// Set game language to French
				data.DataConfigHangman.Language = "fr"
				data.TranslationHangman = configuration.DataTranslationHangmanFr() // Set French translations for the game
				data.HangmanWebTranslations = config.DataTranslationHangmanWebFr() // Set French web translations

				// Define the file map for French difficulty levels
				fileMap = map[int]string{
					1: data.DataFiles.NameFilesConfigFrenchWordsVeryEasy,
					2: data.DataFiles.NameFilesConfigFrenchWordsEasy,
					3: data.DataFiles.NameFilesConfigFrenchWordsDefault,
					4: data.DataFiles.NameFilesConfigFrenchWordsHard,
					5: data.DataFiles.NameFilesConfigFrenchWordsExpert,
					6: data.DataFiles.NameFilesConfigFrenchWordsHacker,
				}
			}

			// Set the file name based on the current difficulty level
			if data.DataConfigHangman.EnableDifficulty {
				// Ensure the file for the current difficulty level exists in the file map
				if file, ok := fileMap[data.GameData.IdDifficulty]; ok {
					data.GameData.NameFile = file
				} else {
					data.GameData.NameFile = fileMap[3] // Default to file for level 3 (Normal) if difficulty is not found
				}
			} else {
				// Default to file for level 3 (Normal) if difficulty is not enabled
				data.GameData.NameFile = fileMap[3]
			}

			// Perform case management for files
			managefiles.CaseManagementFiles(
				data.DataFiles, data.DataConfigHangman,
				data.DataGameHangman, data.DataError,
				data.TranslationHangman,
			)

			// Set additional game configurations based on language
			data.GameData.AddFormForTheExtensionLanguage = game.ExtensionSwitchLanguage(data)
			data.GameData.AddButtonForUseJoker = game.AddButtonJoker(data)
			data.GameData.DataInputField = game.AddInputField(data, data.GameData.StateInputField)
			data.GameData.AddFormForTheExtensionAddWordAfterGame = game.ExtensionInsertWord(data)
			data.GameData.AddFormForTheExtensionLanguage = game.ExtensionSwitchLanguage(data)
			data.GameData.AddFormForTheExtensionEnableDifficulty = game.ExtensionDifficulty(data)
			data.GameData.AddFormForTheExtensionEnableJokers = game.ExtensionJokers(data)
			data.GameData.AddFormForTheExtensionVictoryCounter = game.ExtensionCountVictoryAndLoose(data)

			// Redirect to the configuration page after applying language changes
			Redirect(w, r, "/configuration.html")
		} else {
			// If the method is not POST, return a "Method Not Allowed" error
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
