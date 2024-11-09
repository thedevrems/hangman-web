package game

import (
	"hangman-web/config"
	"html/template"
)

func AddButtonJoker(dataHangmanWeb *config.DataHangmanWeb) template.HTML {
	var data template.HTML
	if dataHangmanWeb.DataConfigHangman.EnableJokers {
		data = `
			<form action="/use-joker" method="POST">
 				<button type="submit" class="button-style">Utiliser un joker</button>
			</form>
		`
	} else {
		data = template.HTML("")
	}
	return data
}
