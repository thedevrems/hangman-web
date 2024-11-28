package game

import (
	"fmt"
	"hangman-web/config"
	"html/template"
)

// ExtensionSwitchLanguage generates an HTML form with a checkbox to toggle between
// English and French languages.
//
// Parameters:
//   - dataHangmanWeb (*config.DataHangmanWeb): a structure containing configuration
//     and translation data for the Hangman-Web application.
//
// Returns:
// - template.HTML: an HTML block containing a form with a checkbox to switch between English and French languages.
func ExtensionSwitchLanguage(dataHangmanWeb *config.DataHangmanWeb) template.HTML {
	checked := IfChecked(dataHangmanWeb.DataConfigHangman.Language != "en")

	data := fmt.Sprintf(`
          <form action="enable-extension-switchlanguage" method="POST" class="config-texte-align">
              <label for="switchlanguage">%s</label>
              <div class="button b2" id="button-10">
                  <input type="checkbox" id="switchlanguage" name="switchLanguage" onchange="this.form.submit()" value="EN" class="checkbox" %s />
                  <div class="knobs">
                    <span>FR</span>
                  </div>
                  <div class="layer"></div>
              </div>
          </form>`, dataHangmanWeb.HangmanWebTranslations.TitleExtensionLanguage, checked)

	return template.HTML(data)
}
