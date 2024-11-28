package manageserver

import (
	"hangman-web/config"
	"hangman-web/game"
	"net/http"
)

// HandleSubmitLetter returns an HTTP handler function that handles submitting a single letter guess in the Hangman game.
// When a POST request is made, it retrieves the letter submitted by the user and processes the guess using the game logic.
// The handler updates the game state and redirects the user back to the game page after processing the letter.
//
// Parameters:
// - data (*config.DataHangmanWeb): The game data containing configuration, game state, and translation data.
//
// Returns:
// - http.HandlerFunc: An HTTP handler function that processes the letter guess in the Hangman game.
func HandleSubmitLetter(data *config.DataHangmanWeb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check if the request method is POST
		if r.Method == "POST" {
			// Get the letter submitted by the user
			letter := r.FormValue("letter")

			// Call the game logic to handle the letter guess
			info := game.IntegrationHangman(data, letter)

			// Update the game state with the result (e.g., success or error message)
			data.GameData.PrintInfo = game.PrintInfo(data, info)

			// Redirect the user back to the game page after processing
			Redirect(w, r, "/game.html")
		} else {
			// If the method is not POST, return a "Method Not Allowed" error
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
