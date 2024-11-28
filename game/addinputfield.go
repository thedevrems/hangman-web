package game

import (
	"hangman-web/config"
	"html/template"
)

// AddInputField generates an HTML input field for users to submit letters or words
// in the hangman game, if the `enable` parameter is set to true.
//
// Parameters:
//   - dataHangmanWeb (*config.DataHangmanWeb): a structure containing configuration
//     and translation data for the Hangman-Web application.
//   - enable (bool): a boolean flag indicating whether the input field should be displayed.
//     If set to true, the input field is generated; otherwise, an empty HTML string is returned.
//
// Returns:
//   - template.HTML: an HTML block containing a form with an input field to submit letters
//     or words. If `enable` is false, the returned HTML string is empty.
func AddInputField(dataHangmanWeb *config.DataHangmanWeb, enable bool) template.HTML {
	var data string = ""
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
