package manageserver

import (
	"hangman-web/config"
	"hangman-web/game"
	"net/http"
	"strconv"
)

// HandleSwitchDifficultyLevel returns an HTTP handler function that handles switching the difficulty level for the Hangman game.
// When a POST request is made with the selected difficulty level, the handler updates the game configuration based on the chosen level.
// The handler adjusts the difficulty of the game, including the corresponding word file and difficulty settings, and redirects the user to the configuration page.
//
// Parameters:
// - data (*config.DataHangmanWeb): The game data containing configuration, game state, and translation data.
//
// Returns:
// - http.HandlerFunc: An HTTP handler function that processes the change of difficulty level in the Hangman game.
func HandleSwitchDifficultyLevel(data *config.DataHangmanWeb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Ensure the request method is POST
		if r.Method == "POST" {
			// Retrieve the selected difficulty level from the form (defaults to 3 if invalid)
			difficultyLevel, err := strconv.Atoi(r.FormValue("difficultyLevel"))
			if err != nil {
				difficultyLevel = 3 // Default to Normal difficulty if conversion fails
			}

			// Define the configuration for difficulty levels based on the selected language
			difficultyConfig := map[int]struct {
				file       string
				difficulty string
			}{}

			// Check the language setting (default to English or French)
			if data.DataConfigHangman.Language == "EN" {
				// English difficulty level configurations
				difficultyConfig = map[int]struct {
					file       string
					difficulty string
				}{
					1: {data.DataFiles.NameFilesConfigEnglishWordsVeryEasy, data.TranslationHangman.VeryEasy},
					2: {data.DataFiles.NameFilesConfigEnglishWordsEasy, data.TranslationHangman.Easy},
					3: {data.DataFiles.NameFilesConfigEnglishWordsDefault, data.TranslationHangman.Normal},
					4: {data.DataFiles.NameFilesConfigEnglishWordsHard, data.TranslationHangman.Hard},
					5: {data.DataFiles.NameFilesConfigEnglishWordsExpert, data.TranslationHangman.Expert},
					6: {data.DataFiles.NameFilesConfigEnglishWordsHacker, data.TranslationHangman.Hacker},
				}
			} else {
				// French difficulty level configurations
				difficultyConfig = map[int]struct {
					file       string
					difficulty string
				}{
					1: {data.DataFiles.NameFilesConfigFrenchWordsVeryEasy, data.TranslationHangman.VeryEasy},
					2: {data.DataFiles.NameFilesConfigFrenchWordsEasy, data.TranslationHangman.Easy},
					3: {data.DataFiles.NameFilesConfigFrenchWordsDefault, data.TranslationHangman.Normal},
					4: {data.DataFiles.NameFilesConfigFrenchWordsHard, data.TranslationHangman.Hard},
					5: {data.DataFiles.NameFilesConfigFrenchWordsExpert, data.TranslationHangman.Expert},
					6: {data.DataFiles.NameFilesConfigFrenchWordsHacker, data.TranslationHangman.Hacker},
				}
			}

			// Apply the selected difficulty configuration or default to Normal if the level is unknown
			config, exists := difficultyConfig[difficultyLevel]
			if !exists {
				config = difficultyConfig[3] // Default to Normal difficulty if the level is invalid
			}

			// Update the game data with the new difficulty settings
			data.GameData.NameFile = config.file
			data.GameData.NameDifficulty = config.difficulty

			// Set the difficulty level ID and extension data
			data.GameData.IdDifficulty = difficultyLevel
			data.GameData.AddFormForTheExtensionEnableDifficulty = game.ExtensionDifficulty(data)

			// Redirect to the configuration page after applying changes
			Redirect(w, r, "/configuration.html")
		} else {
			// If the method is not POST, return a "Method Not Allowed" error
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
