package manageserver

import (
	"hangman-web/config"
	"hangman-web/game"
	"net/http"
)

// HandleToggleExtensionVictoryCounter returns an HTTP handler function that enables or disables the "Victory Counter" extension in the Hangman game.
// When a POST request is made with a value indicating whether the "Victory Counter" feature should be enabled or not, the handler updates the game configuration accordingly.
// If the feature is enabled, it modifies the game state to reflect the ability to count victories and defeats in the game.
// The handler redirects the user to the configuration page after processing the request.
//
// Parameters:
// - data (*config.DataHangmanWeb): The game data containing configuration, game state, and translation data.
//
// Returns:
// - http.HandlerFunc: An HTTP handler function that processes enabling/disabling the "Victory Counter" extension in the Hangman game.
func HandleToggleExtensionVictoryCounter(data *config.DataHangmanWeb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Ensure the request method is POST
		if r.Method == "POST" {
			// Retrieve the value of the "victoryCounter" form field
			// If the value is "NO", disable the Victory Counter feature; otherwise, enable it
			data.DataConfigHangman.VictoryCounter = r.FormValue("victoryCounter") != "NO"

			// Update the game state with the form data related to the Victory Counter extension
			data.GameData.AddFormForTheExtensionVictoryCounter = game.ExtensionCountVictoryAndLoose(data)

			// Redirect to the configuration page after processing the request
			Redirect(w, r, "/configuration.html")
		} else {
			// If the method is not POST, return a "Method Not Allowed" error
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
