package manageserver

import (
	"fmt"
	"hangman-web/config"
	"hangman-web/game"
	"hangman-web/inserthtml"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

func CreateServer(dataHangmanWeb *config.DataHangmanWeb) {
	dataConfigServer := dataHangmanWeb.ServerConfig
	// Récupération du path
	Path := dataConfigServer.Path
	// Récupération des path des fichiers des pages
	filePathIndex := filepath.Join(Path, dataConfigServer.IndexPage)
	filePathConfiguration := filepath.Join(Path, dataConfigServer.ConfigPage)
	filePathGame := filepath.Join(Path, dataConfigServer.GamePage)

	// Fonction pour servir la page d'accueil
	pageIndex := func(w http.ResponseWriter, r *http.Request) {
		tmp1 := template.Must(template.ParseFiles(filePathIndex))
		tmp1.Execute(w, dataHangmanWeb)
	}

	// Initialise l'affichage de données d'informations
	dataHangmanWeb.GameData.PrintNbrVictoryAndLoose = inserthtml.PrintNbrVictoryAndLoose(dataHangmanWeb)
	dataHangmanWeb.GameData.PrintNbrJokers = inserthtml.PrintNbrJokers(dataHangmanWeb)
	// Initialise le input
	dataHangmanWeb.GameData.DataInputField = inserthtml.AddInputField(dataHangmanWeb, true)
	// Initialise le jeu avec un mot à trouver
	dataHangmanWeb.GameData.CurrentWordFormattedHTML = game.InitWord(dataHangmanWeb)
	// Initialise le drawhangman
	dataHangmanWeb.GameData.StageHangmanFormattedHTML = inserthtml.FormatDrawHangmanForHTML(dataHangmanWeb)
	// Initialise l'utilisation de jokers
	dataHangmanWeb.GameData.AddButtonForUseJoker = inserthtml.AddButtonJoker(dataHangmanWeb)

	// Initialise les extensions :
	dataHangmanWeb.GameData.AddFormForTheExtensionAddWordAfterGame = inserthtml.ExtensionInsertWord(dataHangmanWeb)
	dataHangmanWeb.GameData.AddFormForTheExtensionLanguage = inserthtml.ExtensionSwitchLanguage(dataHangmanWeb)
	dataHangmanWeb.GameData.AddFormForTheExtensionEnableDifficulty = inserthtml.ExtensionDifficulty(dataHangmanWeb)
	dataHangmanWeb.GameData.AddFormForTheExtensionEnableJokers = inserthtml.ExtensionJokers(dataHangmanWeb)
	dataHangmanWeb.GameData.AddFormForTheExtensionVictoryCounter = inserthtml.ExtensionCountVictoryAndLoose(dataHangmanWeb)

	// Fonction pour servir la page du jeu
	pageGame := func(w http.ResponseWriter, r *http.Request) {
		tmp2 := template.Must(template.ParseFiles(filePathGame))
		tmp2.Execute(w, dataHangmanWeb)
	}

	// Fonction pour servir la page de configuration
	pageConfiguration := func(w http.ResponseWriter, r *http.Request) {
		tmp3 := template.Must(template.ParseFiles(filePathConfiguration))
		tmp3.Execute(w, dataHangmanWeb)
	}

	// Fonction pour gérer la soumission de lettres
	submitLetter := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			letter := r.FormValue("letter")
			// Intéraction avec le hangman
			info := game.IntegrationHangman(dataHangmanWeb, letter)
			dataHangmanWeb.GameData.PrintInfo = inserthtml.PrintInfo(dataHangmanWeb, info)
			fmt.Println(dataHangmanWeb.GameData.StageHangman)
			// Logique de traitement de la lettre reçue (ex: mise à jour du jeu)
			http.Redirect(w, r, "/game.html", http.StatusSeeOther)
		} else {
			http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		}
	}

	resetGame := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			// Réinitialise leu jeu
			dataHangmanWeb.GameData.TabSelectedLetter = []string{}
			dataHangmanWeb.GameData.TabSelectedWord = []string{}
			dataHangmanWeb.GameData.PrintLetterAndWord = inserthtml.PrintLetterAndWord(dataHangmanWeb)
			dataHangmanWeb.GameData.StageHangman = 0
			dataHangmanWeb.GameData.PrintInfo = inserthtml.PrintInfo(dataHangmanWeb, "")
			dataHangmanWeb.GameData.CurrentWordFormattedHTML = game.InitWord(dataHangmanWeb)
			dataHangmanWeb.GameData.StageHangmanFormattedHTML = inserthtml.FormatDrawHangmanForHTML(dataHangmanWeb)
			// Réactive le input
			dataHangmanWeb.GameData.DataInputField = inserthtml.AddInputField(dataHangmanWeb, true)

			if dataHangmanWeb.DataConfigHangman.VictoryCounter && dataHangmanWeb.GameData.CurrentWord != dataHangmanWeb.GameData.TargetWord {
				dataHangmanWeb.GameData.NbrLoose++
				dataHangmanWeb.GameData.PrintNbrVictoryAndLoose = inserthtml.PrintNbrVictoryAndLoose(dataHangmanWeb)
			} else if dataHangmanWeb.DataConfigHangman.VictoryCounter {
				dataHangmanWeb.GameData.NbrVictory++
				dataHangmanWeb.GameData.PrintNbrVictoryAndLoose = inserthtml.PrintNbrVictoryAndLoose(dataHangmanWeb)
			}
			// Redirigez vers la page de jeu
			http.Redirect(w, r, "/game.html", http.StatusSeeOther)
		} else {
			http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		}
	}

	useJoker := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			info := game.UseJoker(dataHangmanWeb)
			dataHangmanWeb.GameData.PrintInfo = inserthtml.PrintInfo(dataHangmanWeb, info)
			// Redirigez vers la page de jeu
			http.Redirect(w, r, "/game.html", http.StatusSeeOther)
		} else {
			http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		}
	}

	FormExtensionCountVictoryAndLoose := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method)
		if r.Method == "POST" {
			testData := r.FormValue("victoryCounter")
			fmt.Println(testData)
			if testData != "YES" {
				dataHangmanWeb.DataConfigHangman.VictoryCounter = true
			} else {
				dataHangmanWeb.DataConfigHangman.VictoryCounter = false
				dataHangmanWeb.GameData.NbrLoose = 0
				dataHangmanWeb.GameData.NbrVictory = 0
				dataHangmanWeb.GameData.PrintNbrVictoryAndLoose = inserthtml.PrintNbrVictoryAndLoose(dataHangmanWeb)
			}

			dataHangmanWeb.GameData.AddFormForTheExtensionVictoryCounter = inserthtml.ExtensionCountVictoryAndLoose(dataHangmanWeb)

			http.Redirect(w, r, "/configuration.html", http.StatusSeeOther)
		} else {
			http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		}
	}

	FormExtensionInsertWord := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			testData := r.FormValue("insertWord")

			if testData != "YES" {
				dataHangmanWeb.DataConfigHangman.AddWordAfterGame = true
			} else {
				dataHangmanWeb.DataConfigHangman.AddWordAfterGame = false
			}

			dataHangmanWeb.GameData.AddFormForTheExtensionAddWordAfterGame = inserthtml.ExtensionInsertWord(dataHangmanWeb)
			http.Redirect(w, r, "/configuration.html", http.StatusSeeOther)
		} else {
			http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		}
	}

	FormExtensionSwitchLanguage := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			fmt.Println(r.Method)
			testData := r.FormValue("language")
			fmt.Println(testData)
			if testData == "en" {
				dataHangmanWeb.DataConfigHangman.Language = "en"
			} else {
				dataHangmanWeb.DataConfigHangman.Language = "fr"
			}

			dataHangmanWeb.GameData.AddFormForTheExtensionLanguage = inserthtml.ExtensionSwitchLanguage(dataHangmanWeb)

			http.Redirect(w, r, "/configuration.html", http.StatusSeeOther)
		} else {
			http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		}
	}

	// Gestion des routes
	http.HandleFunc(dataConfigServer.Slash, pageIndex)
	http.HandleFunc(dataConfigServer.Slash+dataConfigServer.GamePage, pageGame)
	http.HandleFunc(dataConfigServer.Slash+dataConfigServer.ConfigPage, pageConfiguration)

	http.HandleFunc("/submit-letter-word", submitLetter) // Route pour soumettre la lettre
	http.HandleFunc("/reset-game", resetGame)            // Route pour réinitialiser la partie
	http.HandleFunc("/use-joker", useJoker)              // Route pour utiliser un joker

	// route pour les extensions
	http.HandleFunc("/enable-extension-victorycounter", FormExtensionCountVictoryAndLoose)
	http.HandleFunc("/enable-extension-insertword", FormExtensionInsertWord)
	http.HandleFunc("/enable-extension-switchlanguage", FormExtensionSwitchLanguage)

	staticPath := filepath.Join(Path, "assets")
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir(staticPath))))

	// Ouvre la Première du Site
	go OpenBrowser(dataHangmanWeb.ServerConfig)

	log.Fatal(http.ListenAndServe(dataHangmanWeb.ServerConfig.Port, nil))
}
