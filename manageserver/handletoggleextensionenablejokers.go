package manageserver

import (
	"hangman-web/config"
	"hangman-web/game"
	"net/http"
)

// Handler to enable or disable the "Jokers" extension
func HandleToggleExtensionEnableJokers(data *config.DataHangmanWeb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			data.DataConfigHangman.EnableJokers = r.FormValue("enableJokers") != "NO"
			data.GameData.AddFormForTheExtensionEnableJokers = game.ExtensionJokers(data)
			data.GameData.AddButtonForUseJoker = game.AddButtonJoker(data)
			data.GameData.PrintNbrJokers = game.PrintNbrJokers(data)
			Redirect(w, r, "/configuration.html")
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
