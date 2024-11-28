package manageserver

import (
	"hangman-web/config"
	"hangman-web/game"
	"log"
	"net/http"
	"path/filepath"
)

// CreateServer initializes and starts an HTTP server to serve the Hangman web game, handling routes for
// various pages, game actions, and extensions like victory counters, word insertion, jokers, and language switching.
//
// This function sets up the following routes:
// - `/` (index page)
// - `/game` (game page)
// - `/config` (config page)
// - `/submit-letter-word` (submit a letter or word guess)
// - `/reset-game` (reset the game)
// - `/use-joker` (use a joker)
// - Various extension routes for enabling/disabling features and switching languages and difficulty levels
// - Static assets are served from the "assets" directory
//
// Parameters:
// - dataHangmanWeb (*config.DataHangmanWeb): The game data that holds all game configuration, server settings, and state.
func CreateServer(dataHangmanWeb *config.DataHangmanWeb) {
	// Retrieve the server configuration from the dataHangmanWeb object
	dataConfigServer := dataHangmanWeb.ServerConfig
	Path := dataConfigServer.Path
	staticPath := filepath.Join(Path, "assets") // Path for static assets

	// Define paths to the templates for each page
	filePaths := map[string]string{
		"index":  filepath.Join(Path, dataConfigServer.IndexPage),  // Path to the index page template
		"config": filepath.Join(Path, dataConfigServer.ConfigPage), // Path to the config page template
		"game":   filepath.Join(Path, dataConfigServer.GamePage),   // Path to the game page template
	}

	// Initialize the game state (e.g., setting up initial values)
	game.InitGameState(dataHangmanWeb)

	// Define HTTP route handlers
	// Root route (index page)
	http.HandleFunc(dataConfigServer.Slash, func(w http.ResponseWriter, r *http.Request) {
		ParseAndExecuteTemplate(w, filePaths["index"], dataHangmanWeb)
	})

	// Game page route
	http.HandleFunc(dataConfigServer.Slash+dataConfigServer.GamePage, func(w http.ResponseWriter, r *http.Request) {
		ParseAndExecuteTemplate(w, filePaths["game"], dataHangmanWeb)
	})

	// Configuration page route
	http.HandleFunc(dataConfigServer.Slash+dataConfigServer.ConfigPage, func(w http.ResponseWriter, r *http.Request) {
		ParseAndExecuteTemplate(w, filePaths["config"], dataHangmanWeb)
	})

	// Define game-related action routes
	http.HandleFunc("/submit-letter-word", HandleSubmitLetter(dataHangmanWeb)) // Submit letter or word guess
	http.HandleFunc("/reset-game", HandleResetGame(dataHangmanWeb))            // Reset game state
	http.HandleFunc("/use-joker", HandleUseJoker(dataHangmanWeb))              // Use a joker

	// Define extension-related action routes
	http.HandleFunc("/enable-extension-victorycounter", HandleToggleExtensionVictoryCounter(dataHangmanWeb))     // Toggle victory counter
	http.HandleFunc("/enable-extension-insertword", HandleToggleExtensionInsertWord(dataHangmanWeb))             // Toggle word insertion feature
	http.HandleFunc("/submit-insertword", HandleSubmitInsertWord(dataHangmanWeb))                                // Submit a new word to be used in the game
	http.HandleFunc("/enable-extension-switchlanguage", HandleSwitchLanguage(dataHangmanWeb))                    // Switch language
	http.HandleFunc("/enable-extension-enablejokers", HandleToggleExtensionEnableJokers(dataHangmanWeb))         // Enable or disable jokers
	http.HandleFunc("/enable-extension-enabledifficulty", HandleToggleExtensionEnableDifficulty(dataHangmanWeb)) // Enable or disable difficulty selection
	http.HandleFunc("/enable-extension-switchdifficulty", HandleSwitchDifficultyLevel(dataHangmanWeb))           // Switch the game difficulty level

	// Serve static assets (e.g., images, CSS, JS files)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir(staticPath))))

	// Start the server in a new goroutine and open it in the browser
	go OpenBrowser(dataHangmanWeb.ServerConfig)

	// Listen and serve on the configured port, logging any errors
	log.Fatal(http.ListenAndServe(dataHangmanWeb.ServerConfig.Port, nil))
}
