package manageserver

import (
	"hangman-web/config"
	"hangman-web/game"
	"net/http"
)

// Handler for using a joker
func HandleUseJoker(data *config.DataHangmanWeb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			info := game.UseJoker(data)
			data.GameData.PrintInfo = game.PrintInfo(data, info)
			Redirect(w, r, "/game.html")
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
