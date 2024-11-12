package manageserver

import (
	"hangman-web/config"
	"hangman-web/game"
	"net/http"
)

// Handler for submitting a letter
func HandleSubmitLetter(data *config.DataHangmanWeb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			letter := r.FormValue("letter")
			info := game.IntegrationHangman(data, letter)
			data.GameData.PrintInfo = game.PrintInfo(data, info)
			Redirect(w, r, "/game.html")
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
