package game

import (
	"fmt"
	"hangman-web/config"
	"html/template"
)

// ExtensionInsertWord generates an HTML form with a checkbox to enable or disable
// the feature that allows adding a word after the game ends.
//
// Parameters:
//   - dataHangmanWeb (*config.DataHangmanWeb): a structure containing configuration
//     and translation data for the Hangman-Web application.
//
// Returns:
//   - template.HTML: an HTML block containing a form with a checkbox to enable or disable
//     the "add word after game" feature.
func ExtensionInsertWord(dataHangmanWeb *config.DataHangmanWeb) template.HTML {
	checked := IfChecked(dataHangmanWeb.DataConfigHangman.AddWordAfterGame)

	data := fmt.Sprintf(`
        <form action="enable-extension-insertword" method="POST" class="config-texte-align">
            <label for="insertword">%s</label>
            <div class="button r" id="button-4">
                <input type="checkbox" id="insertword" name="insertWord" onchange="this.form.submit()" value="NO" class="checkbox"%s/>
                <div class="knobs"></div>
                <div class="layer"></div>
            </div>
        </form>`, dataHangmanWeb.HangmanWebTranslations.TitleExtensionAddWordAfterGame, checked)

	return template.HTML(data)
}
