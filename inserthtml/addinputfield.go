package inserthtml

import (
	"hangman-web/config"
	"html/template"
)

func AddInputField(dataHangmanWeb *config.DataHangmanWeb, enable bool) template.HTML {
	var data template.HTML = ``
	if enable {
		data = `
			<form action="/submit-letter-word" method="POST">
				<div class="input-section">
					<input type="text" name="letter" placeholder="Lettre/mot :">
				</div>
			</form>
		`
	}
	return data
}
