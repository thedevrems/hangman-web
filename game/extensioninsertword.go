package game

import (
	"fmt"
	"hangman-web/config"
	"html/template"
)

func ExtensionInsertWord(dataHangmanWeb *config.DataHangmanWeb) template.HTML {
	// Détermine l'état "checked"
	checked := ""
	if !dataHangmanWeb.DataConfigHangman.AddWordAfterGame {
		checked = " checked"
	}

	// Crée et retourne la chaîne HTML
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
