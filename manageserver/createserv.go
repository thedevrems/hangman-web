package manageserver

import (
	"hangman-web/config"
	"hangman-web/game"
	"log"
	"net/http"
	"path/filepath"
)

func CreateServer(dataHangmanWeb *config.DataHangmanWeb) {
	dataConfigServer := dataHangmanWeb.ServerConfig
	Path := dataConfigServer.Path
	staticPath := filepath.Join(Path, "assets")

	// Define file paths for templates
	filePaths := map[string]string{
		"index":  filepath.Join(Path, dataConfigServer.IndexPage),
		"config": filepath.Join(Path, dataConfigServer.ConfigPage),
		"game":   filepath.Join(Path, dataConfigServer.GamePage),
	}

	// Initialize game state
	game.InitGameState(dataHangmanWeb)

	// Define HTTP handlers
	http.HandleFunc(dataConfigServer.Slash, func(w http.ResponseWriter, r *http.Request) {
		ParseAndExecuteTemplate(w, filePaths["index"], dataHangmanWeb)
	})
	http.HandleFunc(dataConfigServer.Slash+dataConfigServer.GamePage, func(w http.ResponseWriter, r *http.Request) {
		ParseAndExecuteTemplate(w, filePaths["game"], dataHangmanWeb)
	})
	http.HandleFunc(dataConfigServer.Slash+dataConfigServer.ConfigPage, func(w http.ResponseWriter, r *http.Request) {
		ParseAndExecuteTemplate(w, filePaths["config"], dataHangmanWeb)
	})

	// Define game action handlers
	http.HandleFunc("/submit-letter-word", HandleSubmitLetter(dataHangmanWeb))
	http.HandleFunc("/reset-game", HandleResetGame(dataHangmanWeb))
	http.HandleFunc("/use-joker", HandleUseJoker(dataHangmanWeb))

	// Define extension handlers
	http.HandleFunc("/enable-extension-victorycounter", HandleToggleExtensionVictoryCounter(dataHangmanWeb))
	http.HandleFunc("/enable-extension-insertword", HandleToggleExtensionInsertWord(dataHangmanWeb))
	http.HandleFunc("/submit-insertword", HandleSubmitInsertWord(dataHangmanWeb))
	http.HandleFunc("/enable-extension-switchlanguage", HandleSwitchLanguage(dataHangmanWeb))
	http.HandleFunc("/enable-extension-enablejokers", HandleToggleExtensionEnableJokers(dataHangmanWeb))
	http.HandleFunc("/enable-extension-enabledifficulty", HandleToggleExtensionEnableDifficulty(dataHangmanWeb))
	http.HandleFunc("/enable-extension-switchdifficulty", HandleSwitchDifficultyLevel(dataHangmanWeb))

	// Serve static assets
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir(staticPath))))

	// Start server and open in browser
	go OpenBrowser(dataHangmanWeb.ServerConfig)
	log.Fatal(http.ListenAndServe(dataHangmanWeb.ServerConfig.Port, nil))
}
