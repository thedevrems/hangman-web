package game

import (
	"hangman-web/config"
	"html/template"
)

func FormatDrawHangmanForHTML(dataHangmanWeb *config.DataHangmanWeb) template.HTML {
	// Retrieve the current stage of the Hangman figure from the game data
	pos := dataHangmanWeb.GameData.StageHangman
	idLevel := dataHangmanWeb.GameData.IdDifficulty

	// Initialize an empty string to hold the HTML structure for the Hangman figure
	var hangmanHTML string

	// Only add the gallows div and its contents if pos is at least 1
	if pos >= 1 {
		hangmanHTML += `<div class="hangman">`
		hangmanHTML += `<div class="gallows">`
		if pos >= 2 {
			hangmanHTML += `<div class="beam"></div>` // Add beam as the second stage
		}
		if pos >= 3 {
			hangmanHTML += `<div class="rope"></div>` // Add rope as the third stage
		}
		if pos >= 4 {
			hangmanHTML += `<div class="support"></div>` // Add support as the fourth stage
		}
		// Add triangle at position 5 if GameData == 1
		if idLevel == 1 && pos >= 5 {
			hangmanHTML += `<div class="triangle-right"></div>` // Add triangle at position 5
		}
		if idLevel == 1 && pos >= 6 {
			hangmanHTML += `<div class="triangle-left"></div>` // Add triangle at position 5
		}
		hangmanHTML += `</div>` // Close gallows div
	}

	// Start the man div, and add parts based on the game stage (pos)
	hangmanHTML += `<div class="man">`

	// Adjust stages to account for triangle being added at position 5
	if idLevel == 1 {
		// Shift body parts to positions starting from 6 onward
		if pos >= 7 {
			hangmanHTML += `<div class="head"></div>` // Add head as the sixth stage
		}
		if pos >= 8 {
			hangmanHTML += `<div class="body"></div>` // Add body as the seventh stage
		}
		if pos >= 9 {
			hangmanHTML += `<div class="arm left"></div>` // Add left arm as the eighth stage
		}
		if pos >= 10 {
			hangmanHTML += `<div class="arm right"></div>` // Add right arm as the ninth stage
		}
		if pos >= 11 {
			hangmanHTML += `<div class="leg left"></div>` // Add left leg as the tenth stage
		}
		if pos >= 12 {
			hangmanHTML += `<div class="leg right"></div>` // Add right leg as the eleventh stage
		}
		// Add eyes and mouth at new positions (12, 13, 14)
		if pos >= 13 {
			hangmanHTML += `<div class="eye left"></div>` // Add left eye as the twelfth stage
		}
		if pos >= 14 {
			hangmanHTML += `<div class="eye right"></div>` // Add right eye as the thirteenth stage
		}
		if pos >= 15 {
			hangmanHTML += `<div class="mouth"></div>` // Add mouth as the fourteenth stage
		}
	} else {
		// For gameData != 1, no shift in positions, standard order of body parts
		if pos >= 5 {
			hangmanHTML += `<div class="head"></div>` // Add head as the fifth stage
		}
		if pos >= 6 {
			hangmanHTML += `<div class="body"></div>` // Add body as the sixth stage
		}
		if pos >= 7 {
			hangmanHTML += `<div class="arm left"></div>` // Add left arm as the seventh stage
		}
		if pos >= 8 {
			hangmanHTML += `<div class="arm right"></div>` // Add right arm as the eighth stage
		}
		if pos >= 9 {
			hangmanHTML += `<div class="leg left"></div>` // Add left leg as the ninth stage
		}
		if pos >= 10 {
			hangmanHTML += `<div class="leg right"></div>` // Add right leg as the tenth stage
		}
		// Add eyes and mouth at positions 11, 12, 13
		if pos >= 11 {
			hangmanHTML += `<div class="eye left"></div>` // Add left eye as the eleventh stage
		}
		if pos >= 12 {
			hangmanHTML += `<div class="eye right"></div>` // Add right eye as the twelfth stage
		}
		if pos >= 13 {
			hangmanHTML += `<div class="mouth"></div>` // Add mouth as the thirteenth stage
		}
	}

	if pos >= 1 {
		hangmanHTML += `</div>` // Close man div
		hangmanHTML += `</div>` // Close hangman div
	}

	// Store the generated HTML in the game data for future reference
	dataHangmanWeb.GameData.StageHangmanFormattedHTML = template.HTML(hangmanHTML)

	// Return the generated Hangman HTML as a template.HTML, ensuring it's rendered as HTML
	return template.HTML(hangmanHTML)
}
