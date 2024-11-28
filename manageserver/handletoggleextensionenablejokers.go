package manageserver

import (
	"hangman-web/config"
	"hangman-web/game"
	"net/http"
)

// HandleToggleExtensionEnableJokers returns an HTTP handler function that enables or disables the "Jokers" extension in the Hangman game.
// When a POST request is made with a value indicating whether the Jokers extension should be enabled or not, the handler updates the game configuration accordingly.
// If the Jokers extension is enabled, it modifies the game state to reflect the addition of Jokers, including displaying the button to use a Joker and updating the number of available Jokers.
// The handler redirects the user to the configuration page after processing the request.
//
// Parameters:
// - data (*config.DataHangmanWeb): The game data containing configuration, game state, and translation data.
//
// Returns:
// - http.HandlerFunc: An HTTP handler function that processes enabling/disabling the Jokers extension in the Hangman game.
func HandleToggleExtensionEnableJokers(data *config.DataHangmanWeb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Ensure the request method is POST
		if r.Method == "POST" {
			// Retrieve the value of the "enableJokers" form field
			// If the value is "NO", disable the Jokers extension; otherwise, enable it
			data.DataConfigHangman.EnableJokers = r.FormValue("enableJokers") != "NO"

			// Update the game state with the form data related to Jokers
			data.GameData.AddFormForTheExtensionEnableJokers = game.ExtensionJokers(data)
			data.GameData.AddButtonForUseJoker = game.AddButtonJoker(data)
			data.GameData.PrintNbrJokers = game.PrintNbrJokers(data)

			// Redirect to the configuration page after processing the request
			Redirect(w, r, "/configuration.html")
		} else {
			// If the method is not POST, return a "Method Not Allowed" error
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
