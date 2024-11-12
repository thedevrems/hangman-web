package inserthtml

import (
	"fmt"
	"hangman-web/config"
	"html/template"
)

func ExtensionJokers(dataHangmanWeb *config.DataHangmanWeb) template.HTML {
	// Détermine l'état "checked" en fonction du compteur de victoire
	checked := ""
	if !dataHangmanWeb.DataConfigHangman.EnableJokers {
		checked = " checked"
	}

	// Crée et retourne la chaîne HTML
	data := fmt.Sprintf(`
        <form action="enable-extension-enablejokers" method="POST" class="config-texte-align">
            <label for="enablejokers">%s</label>
            <div class="button r" id="button-4">
                <input type="checkbox" id="enablejokers" name="enableJokers" onchange="this.form.submit()" value="YES" class="checkbox"%s/>
                <div class="knobs"></div>
                <div class="layer"></div>
            </div>
        </form>`, dataHangmanWeb.HangmanWebTranslations.TitleExtensionJokers, checked)

	return template.HTML(data)
}
