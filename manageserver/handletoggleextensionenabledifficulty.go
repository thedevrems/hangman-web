package manageserver

import (
	"hangman-web/config"
	"hangman-web/game"
	"net/http"

	"github.com/thedevrems/hangman/managefiles"
)

// Handler to toggle the "Difficulty" extension
func HandleToggleExtensionEnableDifficulty(data *config.DataHangmanWeb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			data.DataConfigHangman.EnableDifficulty = r.FormValue("enableDifficulty") != "NO"
			data.GameData.NameDifficulty = data.TranslationHangman.Normal

			if data.DataConfigHangman.Language == "en" {
				data.GameData.NameFile = data.DataFiles.NameFilesConfigEnglishWordsDefault
			} else {
				data.GameData.NameFile = data.DataFiles.NameFilesConfigFrenchWordsDefault
			}

			data.GameData.IdDifficulty = 3
			data.GameData.AddFormForTheExtensionEnableDifficulty = game.ExtensionDifficulty(data)

			if data.DataConfigHangman.EnableDifficulty {
				managefiles.CaseManagementFiles(
					data.DataFiles, data.DataConfigHangman,
					data.DataGameHangman, data.DataError,
					data.TranslationHangman,
				)
			}

			Redirect(w, r, "/configuration.html")
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
