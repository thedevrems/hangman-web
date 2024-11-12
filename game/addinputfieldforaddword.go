package game

import (
	"hangman-web/config"
	"html/template"
)

func AddInputFieldForAddWord(dataHangmanWeb *config.DataHangmanWeb, enable bool) template.HTML {
	var data string
	var placeholder string = dataHangmanWeb.HangmanWebTranslations.PlaceHolderInputFieldForAddWord

	if enable {
		data = `
			<form action="/submit-insertword" method="POST">
				<div class="input-section">
					<input type="text" name="addWord" placeholder="` + placeholder + `">
				</div>
			</form>
		`
	}
	return template.HTML(data)
}
