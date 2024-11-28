package manageserver

import (
	"hangman-web/config"
	"hangman-web/game"
	"net/http"

	"github.com/thedevrems/hangman/managefiles"
)

// HandleSubmitInsertWord returns an HTTP handler function that handles submitting a new word to be added
// to the Hangman game word list. When a POST request is made, it checks if the word input is empty or not.
// If the word is not empty, it appends the word to a file. If the word input is empty, it returns an error message
// indicating that the word must be provided. After the operation, it redirects the user to the game page.
//
// Parameters:
// - data (*config.DataHangmanWeb): The game data containing configuration, game state, and translation data.
//
// Returns:
// - http.HandlerFunc: An HTTP handler function that processes word submissions for the Hangman game.
func HandleSubmitInsertWord(data *config.DataHangmanWeb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check if the request method is POST
		if r.Method == "POST" {
			// Get the word submitted by the user
			word := r.FormValue("addWord")

			// If the word is empty, set an error message
			if word == "" {
				data.GameData.PrintInfo = game.PrintInfo(data, data.TranslationHangman.RespectFormulation)
			} else {
				// If the word is not empty, append it to the word list file
				managefiles.AppendWordToFile(data.DataError, data.TranslationHangman, data.DataFiles, data.GameData.NameFile, word)
			}

			// Redirect the user to the game page after processing the word
			Redirect(w, r, "/game.html")
		} else {
			// If the method is not POST, return a "Method Not Allowed" error
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
