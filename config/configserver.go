package config

// Initialization of values by default of the ServerConfig structure
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
