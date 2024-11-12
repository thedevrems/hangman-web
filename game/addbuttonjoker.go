package game

import (
	"hangman-web/config"
	"html/template"
)

func AddButtonJoker(dataHangmanWeb *config.DataHangmanWeb) template.HTML {
	var data string
	var title string = dataHangmanWeb.HangmanWebTranslations.UseJoker
	if dataHangmanWeb.DataConfigHangman.EnableJokers {
		data = `
			<form action="/use-joker" method="POST">
 				<button type="submit" class="button-style">` + title + `</button>
			</form>
		`
	}
	return template.HTML(data)
}
