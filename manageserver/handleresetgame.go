package manageserver

import (
	"hangman-web/config"
	"hangman-web/game"
	"net/http"
)

func HandleResetGame(data *config.DataHangmanWeb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			game.ResetGame(data)
			Redirect(w, r, "/game.html")
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
