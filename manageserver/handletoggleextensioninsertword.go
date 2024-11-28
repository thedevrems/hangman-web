package manageserver

import (
	"hangman-web/config"
	"hangman-web/game"
	"net/http"
)

// HandleToggleExtensionInsertWord returns an HTTP handler function that enables or disables the "Insert Word After Game" extension in the Hangman game.
// When a POST request is made with a value indicating whether the "Insert Word" feature should be enabled or not, the handler updates the game configuration accordingly.
// If the feature is enabled, it modifies the game state to reflect the ability to insert words after the game ends.
// The handler redirects the user to the configuration page after processing the request.
//
// Parameters:
// - data (*config.DataHangmanWeb): The game data containing configuration, game state, and translation data.
//
// Returns:
// - http.HandlerFunc: An HTTP handler function that processes enabling/disabling the "Insert Word After Game" extension in the Hangman game.
func HandleToggleExtensionInsertWord(data *config.DataHangmanWeb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Ensure the request method is POST
		if r.Method == "POST" {
			// Retrieve the value of the "insertWord" form field
			// If the value is "NO", disable the Insert Word feature; otherwise, enable it
			data.DataConfigHangman.AddWordAfterGame = r.FormValue("insertWord") != "NO"

			// Update the game state with the form data related to inserting words after the game
			data.GameData.AddFormForTheExtensionAddWordAfterGame = game.ExtensionInsertWord(data) // Add form for Insert Word feature

			// Redirect to the configuration page after processing the request
			Redirect(w, r, "/configuration.html")
		} else {
			// If the method is not POST, return a "Method Not Allowed" error
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
