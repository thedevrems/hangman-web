package manageserver

import (
	"hangman-web/config"
	"hangman-web/game"
	"net/http"
)

// HandleResetGame returns an HTTP handler function that handles the resetting of the Hangman game state.
// When a POST request is made to this handler, it calls the `ResetGame` function to reset the game data
// and then redirects the user to the game page. If the request method is not POST, it returns a "Method Not Allowed" error.
//
// Parameters:
// - data (*config.DataHangmanWeb): The game data that holds the current state of the Hangman game and its configuration.
//
// Returns:
// - http.HandlerFunc: An HTTP handler function that processes requests to reset the game.
func HandleResetGame(data *config.DataHangmanWeb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// If the request method is POST, reset the game
		if r.Method == "POST" {
			// Reset the game state
			game.ResetGame(data)

			// Redirect to the game page after resetting the game
			Redirect(w, r, "/game.html")
		} else {
			// If the method is not POST, return a "Method Not Allowed" error
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
