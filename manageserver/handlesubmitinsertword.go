package manageserver

import (
	"hangman-web/config"
	"hangman-web/game"
	"net/http"

	"github.com/thedevrems/hangman/managefiles"
)

// Handler for submitting a letter
func HandleSubmitInsertWord(data *config.DataHangmanWeb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			word := r.FormValue("addWord")
			if word == "" {
				data.GameData.PrintInfo = game.PrintInfo(data, data.TranslationHangman.RespectFormulation)
			} else {
				managefiles.AppendWordToFile(data.DataError, data.TranslationHangman, data.DataFiles, data.GameData.NameFile, word)
			}
			Redirect(w, r, "/game.html")
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
