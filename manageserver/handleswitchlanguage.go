package manageserver

import (
	"hangman-web/config"
	"hangman-web/game"
	"net/http"

	"github.com/thedevrems/hangman/configuration"
	"github.com/thedevrems/hangman/managefiles"
)

// Handler to switch the language extension
func HandleSwitchLanguage(data *config.DataHangmanWeb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			language := r.FormValue("switchLanguage")

			var fileMap map[int]string
			if language == "EN" {
				data.DataConfigHangman.Language = "en"
				data.TranslationHangman = configuration.DataTranslationHangmanEn()
				data.HangmanWebTranslations = config.DataTranslationHangmanWebEn()

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
				data.DataConfigHangman.Language = "fr"
				data.TranslationHangman = configuration.DataTranslationHangmanFr()
				data.HangmanWebTranslations = config.DataTranslationHangmanWebFr()

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

			// Set the file name based on difficulty, with a default value if difficulty is not found
			if data.DataConfigHangman.EnableDifficulty {
				if file, ok := fileMap[data.GameData.IdDifficulty]; ok {
					data.GameData.NameFile = file
				} else {
					data.GameData.NameFile = fileMap[3] // Default file for level 3
				}
			} else {
				data.GameData.NameFile = fileMap[3] // Default file for level 3
			}

			// Additional configurations
			managefiles.CaseManagementFiles(
				data.DataFiles, data.DataConfigHangman,
				data.DataGameHangman, data.DataError,
				data.TranslationHangman,
			)
			data.GameData.AddFormForTheExtensionLanguage = game.ExtensionSwitchLanguage(data)
			data.GameData.AddButtonForUseJoker = game.AddButtonJoker(data)
			data.GameData.DataInputField = game.AddInputField(data, data.GameData.StateInputField)
			data.GameData.AddFormForTheExtensionAddWordAfterGame = game.ExtensionInsertWord(data)
			data.GameData.AddFormForTheExtensionLanguage = game.ExtensionSwitchLanguage(data)
			data.GameData.AddFormForTheExtensionEnableDifficulty = game.ExtensionDifficulty(data)
			data.GameData.AddFormForTheExtensionEnableJokers = game.ExtensionJokers(data)
			data.GameData.AddFormForTheExtensionVictoryCounter = game.ExtensionCountVictoryAndLoose(data)

			Redirect(w, r, "/configuration.html")
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
