package config

func DataTranslationHangmanWebFr() *HangmanWebTranslations {
	return &HangmanWebTranslations{
		Title:           "Hangman Web",
		Description:     "Bienvenue dans le Hangame, un jeu d'agilité et de stratégie où seuls tes réflexes comptent ! Reproduis ou déjoue les gestes de tes adversaires avec précision et rapidité. Simple à jouer, mais difficile à maîtriser, ce jeu mettra à l'épreuve ta vitesse et ta concentration. Prêt à prouver que tes mains sont imbattables ? À toi de jouer !",
		StartButtonText: "Commencer",
		NavAbout:        "À propos",
		NavGame:         "Jeu",
		NavConfig:       "Configuration",
		NavLogoTitle:    "HANG • MAN • WEB",
	}
}

func DataTranslationHangmanWebEn() *HangmanWebTranslations {
	return &HangmanWebTranslations{
		Title:           "Hangman Web",
		Description:     "Welcome to Hangame, a game of agility and strategy where only your reflexes matter! Reproduce or counteract your opponents' moves with precision and speed. Simple to play but challenging to master, this game will test your speed and focus. Ready to prove your hands are unbeatable? It's your turn!",
		StartButtonText: "Get Started",
		NavAbout:        "About",
		NavGame:         "Game",
		NavConfig:       "Settings",
		NavLogoTitle:    "HANG • MAN • WEB",
	}
}
