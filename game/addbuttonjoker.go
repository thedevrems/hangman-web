package game

import (
	"hangman-web/config"
	"html/template"
)

// AddButtonJoker generates an HTML button that allows players to use a joker
// in the hangman game if the joker option is enabled in the configuration.
//
// Parameters:
//   - dataHangmanWeb (*config.DataHangmanWeb): a structure containing configuration
//     and translation data for the Hangman-Web application.
//
// Returns:
//   - template.HTML: an HTML block containing a form with a button to use a joker.
//     If jokers are not enabled, the returned HTML string is empty.
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
