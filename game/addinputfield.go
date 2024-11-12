package game

import (
	"hangman-web/config"
	"html/template"
)

func AddInputField(dataHangmanWeb *config.DataHangmanWeb, enable bool) template.HTML {
	var data string
	var placeholder string = dataHangmanWeb.HangmanWebTranslations.PlaceHolderInputField
	if enable {
		data = `
			<form action="/submit-letter-word" method="POST">
				<div class="input-section">
					<input type="text" name="letter" placeholder="` + placeholder + `">
				</div>
			</form>
		`
	}
	return template.HTML(data)
}
