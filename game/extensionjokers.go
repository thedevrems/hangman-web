package game

import (
	"fmt"
	"hangman-web/config"
	"html/template"
)

// ExtensionJokers generates an HTML form with a checkbox to enable or disable
// the jokers feature.
// Parameters:
//   - dataHangmanWeb (*config.DataHangmanWeb): a structure containing configuration
//     and translation data for the Hangman-Web application.
//
// Returns:
//   - template.HTML: an HTML block containing a form with a checkbox to enable or disable
//     the jokers feature.
func ExtensionJokers(dataHangmanWeb *config.DataHangmanWeb) template.HTML {
	checked := IfChecked(dataHangmanWeb.DataConfigHangman.EnableJokers)

	data := fmt.Sprintf(`
        <form action="enable-extension-enablejokers" method="POST" class="config-texte-align">
            <label for="enablejokers">%s</label>
            <div class="button r" id="button-4">
                <input type="checkbox" id="enablejokers" name="enableJokers" onchange="this.form.submit()" value="NO" class="checkbox"%s/>
                <div class="knobs"></div>
                <div class="layer"></div>
            </div>
        </form>`, dataHangmanWeb.HangmanWebTranslations.TitleExtensionJokers, checked)

	return template.HTML(data)
}
