package inserthtml

import (
	"fmt"
	"hangman-web/config"
	"html/template"
)

func ExtensionSwitchLanguage(dataHangmanWeb *config.DataHangmanWeb) template.HTML {
	checked := ""

	if dataHangmanWeb.DataConfigHangman.Language == "en" {
		checked = " checked"
	}

	data := fmt.Sprintf(`
          <form action="enable-extension-switchlanguage" method="POST" class="config-texte-align">
              <label for="language-select">%s</label>
              <div class="button b2" id="button-10">
                  <input type="checkbox" id="language-select" name="language" onchange="this.form.submit()" value="EN" class="checkbox" %s />
                  <div class="knobs">
                    <span>FR</span>
                  </div>
                  <div class="layer"></div>
              </div>
          </form>`, dataHangmanWeb.HangmanWebTranslations.TitleExtensionLanguage, checked)

	return template.HTML(data)
}
