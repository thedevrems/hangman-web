package game

import (
	"fmt"
	"hangman-web/config"
	"html/template"
)

// ExtensionCountVictoryAndLoose generates an HTML form with a checkbox to toggle
// the victory counter extension in the hangman game.
//
// Parameters:
//   - dataHangmanWeb (*config.DataHangmanWeb): a structure containing configuration
//     and translation data for the Hangman-Web application, including information
//     about whether the victory counter is enabled and the translated label text for
//     the checkbox.
//
// Returns:
//   - template.HTML: an HTML block containing a form with a checkbox to enable or disable
//     the victory counter. The checkbox is pre-checked if the victory counter is currently disabled.
func ExtensionCountVictoryAndLoose(dataHangmanWeb *config.DataHangmanWeb) template.HTML {
	checked := IfChecked(dataHangmanWeb.DataConfigHangman.VictoryCounter)

	data := fmt.Sprintf(`
        <form action="/enable-extension-victorycounter" method="POST" class="config-texte-align">
            <label for="victorycounter">%s</label>
            <div class="button r" id="button-4">
                <input type="checkbox" id="victorycounter" name="victoryCounter" onchange="this.form.submit()" value="NO" class="checkbox" %s/>
                <div class="knobs"></div>
                <div class="layer"></div>
            </div>
        </form>`, dataHangmanWeb.HangmanWebTranslations.TitleExtensionCountVictory, checked)

	return template.HTML(data)
}
