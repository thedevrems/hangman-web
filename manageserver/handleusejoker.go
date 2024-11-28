package manageserver

import (
	"hangman-web/config"
	"hangman-web/game"
	"net/http"
)

// HandleUseJoker returns an HTTP handler function that allows the player to use a "joker" in the Hangman game.
// When a POST request is made, the handler invokes the `game.UseJoker` function to process the use of a joker.
// If the joker is successfully used, the game state is updated accordingly. The handler then redirects the player back to the game page.
// If the method is not POST, a "Method Not Allowed" error is returned.
//
// Parameters:
// - data (*config.DataHangmanWeb): The game data containing the game configuration, state, and translation data.
//
// Returns:
// - http.HandlerFunc: An HTTP handler function that processes the usage of a joker in the Hangman game.
func HandleUseJoker(data *config.DataHangmanWeb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Ensure the request method is POST
		if r.Method == "POST" {
			// Call the UseJoker function to process the joker usage
			info := game.UseJoker(data)

			// Update the game information to reflect the joker's result
			data.GameData.PrintInfo = game.PrintInfo(data, info)

			// Redirect the player back to the game page to see the updated state
			Redirect(w, r, "/game.html")
		} else {
			// If the method is not POST, return a "Method Not Allowed" error
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
