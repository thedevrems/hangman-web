package config

// data de la configuration du serveur
func ConfigServer() *ServerConfig {
	return &ServerConfig{
		Path:       "templates",
		IndexPage:  "index.html",
		ConfigPage: "configuration.html",
		GamePage:   "game.html",
		URL:        "http://localhost",
		Port:       ":8000",
		Slash:      "/",
	}
}
