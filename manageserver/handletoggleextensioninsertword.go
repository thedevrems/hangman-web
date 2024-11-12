package manageserver

import (
	"hangman-web/config"
	"hangman-web/game"
	"net/http"
)

func HandleToggleExtensionInsertWord(data *config.DataHangmanWeb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			data.DataConfigHangman.AddWordAfterGame = r.FormValue("insertWord") != "NO"
			data.GameData.AddFormForTheExtensionAddWordAfterGame = game.ExtensionInsertWord(data)
			Redirect(w, r, "/configuration.html")
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
