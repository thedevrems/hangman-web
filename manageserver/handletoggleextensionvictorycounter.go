package manageserver

import (
	"hangman-web/config"
	"hangman-web/game"
	"net/http"
)

func HandleToggleExtensionVictoryCounter(data *config.DataHangmanWeb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			data.DataConfigHangman.VictoryCounter = r.FormValue("victoryCounter") != "NO"
			data.GameData.AddFormForTheExtensionVictoryCounter = game.ExtensionCountVictoryAndLoose(data)
			Redirect(w, r, "/configuration.html")
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
