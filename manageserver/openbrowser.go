package manageserver

import (
	"hangman-web/config"
	"log"
	"os/exec"
	"runtime"
)

// OpenBrowser attempts to open the game server's URL in the default web browser based on the operating system.
//
// This function determines the operating system of the host machine (Linux, Windows, etc.) and uses the
// appropriate command to open the provided URL in the default web browser. The URL is formed by appending
// the port from the `dataConfigServer` configuration to the server's base URL.
//
// Parameters:
//   - dataConfigServer (*config.ServerConfig): The server configuration data, which contains the URL and port
//     for the server to be opened in a browser.
//
// Behavior:
// - On Linux, the function uses the `xdg-open` command to open the URL.
// - On Windows, it uses the `rundll32` command with `url.dll` to open the URL.
// - If the system is not Linux or Windows, the function will not attempt to open the browser and will log an error.
//
// If there is an error executing the command to open the browser, the function logs the error and exits the program.
func OpenBrowser(dataConfigServer *config.ServerConfig) {
	var err error
	// Construct the URL by combining the server base URL and the port
	var url string = dataConfigServer.URL + dataConfigServer.Port

	// Determine the operating system and run the appropriate command to open the URL in a browser
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	}

	// If there was an error executing the command, log it and exit the program
	if err != nil {
		log.Fatal(err)
	}
}
