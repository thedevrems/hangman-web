package manageserver

import (
	"hangman-web/config"
	"hangman-web/game"
	"net/http"
	"strconv"
)

// Handler to switch difficulty levels
func HandleSwitchDifficultyLevel(data *config.DataHangmanWeb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			difficultyLevel, err := strconv.Atoi(r.FormValue("difficultyLevel"))
			if err != nil {
				difficultyLevel = 3 // Default to Normal
			}

			// Define difficulty levels configuration based on language
			difficultyConfig := map[int]struct {
				file       string
				difficulty string
			}{}

			if data.DataConfigHangman.Language == "EN" {
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

			// Apply the selected difficulty level configuration or fall back to default
			config, exists := difficultyConfig[difficultyLevel]
			if !exists {
				config = difficultyConfig[3] // Default to Normal if level is unknown
			}
			data.GameData.NameFile = config.file
			data.GameData.NameDifficulty = config.difficulty

			// Set the remaining game data properties
			data.GameData.IdDifficulty = difficultyLevel
			data.GameData.AddFormForTheExtensionEnableDifficulty = game.ExtensionDifficulty(data)

			Redirect(w, r, "/configuration.html")
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
