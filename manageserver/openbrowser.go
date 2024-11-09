package manageserver

import (
	"hangman-web/config"
	"log"
	"os/exec"
	"runtime"
)

func OpenBrowser(dataConfigServer *config.ServerConfig) {
	var err error
	// Récupération de l'url de la page
	var url string = dataConfigServer.URL + dataConfigServer.Port

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	}

	if err != nil {
		log.Fatal(err)
	}
}
