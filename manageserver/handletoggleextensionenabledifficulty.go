package manageserver

import (
	"hangman-web/config"
	"hangman-web/game"
	"net/http"

	"github.com/thedevrems/hangman/managefiles"
)

// HandleToggleExtensionEnableDifficulty returns an HTTP handler function that toggles the "Difficulty" extension in the Hangman game.
// When a POST request is made with a value indicating whether the difficulty extension should be enabled or not, the handler updates the game configuration accordingly.
// If the difficulty extension is enabled, it sets the appropriate difficulty level and word file, and manages game files.
// The handler redirects the user to the configuration page after processing the request.
//
// Parameters:
// - data (*config.DataHangmanWeb): The game data containing configuration, game state, and translation data.
//
// Returns:
// - http.HandlerFunc: An HTTP handler function that processes enabling/disabling the difficulty extension in the Hangman game.
func HandleToggleExtensionEnableDifficulty(data *config.DataHangmanWeb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Ensure the request method is POST
		if r.Method == "POST" {
			// Retrieve the value of the "enableDifficulty" form field
			// If the value is "NO", disable the difficulty extension; otherwise, enable it
			data.DataConfigHangman.EnableDifficulty = r.FormValue("enableDifficulty") != "NO"
			data.GameData.NameDifficulty = data.TranslationHangman.Normal // Set default difficulty to "Normal" (level 3)

			// Based on the current language, set the corresponding word file for the difficulty level
			if data.DataConfigHangman.Language == "en" {
				data.GameData.NameFile = data.DataFiles.NameFilesConfigEnglishWordsDefault // English default file for difficulty 3 (Normal)
			} else {
				data.GameData.NameFile = data.DataFiles.NameFilesConfigFrenchWordsDefault // French default file for difficulty 3 (Normal)
			}

			// Set the default difficulty level to 3 (Normal)
			data.GameData.IdDifficulty = 3
			// Set the form for the "Enable Difficulty" extension
			data.GameData.AddFormForTheExtensionEnableDifficulty = game.ExtensionDifficulty(data)

			// If the difficulty extension is enabled, manage game files accordingly
			if data.DataConfigHangman.EnableDifficulty {
				managefiles.CaseManagementFiles(
					data.DataFiles, data.DataConfigHangman,
					data.DataGameHangman, data.DataError,
					data.TranslationHangman,
				)
			}

			// Redirect the user to the configuration page after applying the changes
			Redirect(w, r, "/configuration.html")
		} else {
			// If the method is not POST, return a "Method Not Allowed" error
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
