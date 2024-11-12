package game

import (
	"fmt"
	"hangman-web/config"
	"html/template"
)

func ExtensionCountVictoryAndLoose(dataHangmanWeb *config.DataHangmanWeb) template.HTML {
	// Détermine l'état "checked"
	checked := ""
	if !dataHangmanWeb.DataConfigHangman.VictoryCounter {
		checked = " checked"
	}

	// Crée et retourne la chaîne HTML
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
