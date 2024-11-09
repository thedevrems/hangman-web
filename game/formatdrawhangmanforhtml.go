package game

import (
	"hangman-web/config"
	"html/template"
)

func FormatDrawHangmanForHTML(dataHangmanWeb *config.DataHangmanWeb) template.HTML {
	pos := dataHangmanWeb.GameData.StageHangman

	// Crée des divs dynamiques en fonction de l'étape actuelle du jeu
	var hangmanHTML string
	hangmanHTML += `<div class="gallows">`
	hangmanHTML += `<div class="hangman">`
	hangmanHTML += `<div class="pole"></div>`
	hangmanHTML += `<div class="rope"></div>`

	// On ajoute les parties du pendu en fonction de la valeur de `pos`
	if pos >= 1 {
		hangmanHTML += `<div class="head"></div>`
	}
	if pos >= 2 {
		hangmanHTML += `<div class="body"></div>`
	}
	if pos >= 3 {
		hangmanHTML += `<div class="left-arm"></div>`
	}
	if pos >= 4 {
		hangmanHTML += `<div class="right-arm"></div>`
	}
	if pos >= 5 {
		hangmanHTML += `<div class="left-leg"></div>`
	}
	if pos >= 6 {
		hangmanHTML += `<div class="right-leg"></div>`
	}

	// Fermeture des divs principales
	hangmanHTML += `</div></div>`

	// Enregistrement de la position du hangman
	dataHangmanWeb.GameData.StageHangmanFormattedHTML = template.HTML(hangmanHTML)
	// On retourne le HTML en tant que template.HTML pour être sûr qu'il est interprété comme du code HTML
	return template.HTML(hangmanHTML)
}
