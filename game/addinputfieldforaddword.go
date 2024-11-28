package game

import (
	"hangman-web/config"
	"html/template"
)

// AddInputFieldForAddWord generates an HTML input field for users to add a custom word
// in the hangman game, if the `enable` parameter is set to true.
//
// Parameters:
//   - dataHangmanWeb (*config.DataHangmanWeb): a structure containing configuration
//     and translation data for the Hangman-Web application.
//   - enable (bool): a boolean flag indicating whether the input field should be displayed.
//
// Returns:
//   - template.HTML: an HTML block containing a form with an input field to submit a custom word.
//     If `enable` is false, the returned HTML string is empty.
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
