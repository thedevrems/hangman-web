package inserthtml

import (
	"fmt"
	"hangman-web/config"
	"html/template"
)

func ExtensionDifficulty(dataHangmanWeb *config.DataHangmanWeb) template.HTML {
	// Variables de texte
	title := dataHangmanWeb.HangmanWebTranslations.TitleExtensionEnableDifficulty
	levels := []string{
		dataHangmanWeb.HangmanWebTranslations.ExtensionDifficultyLevel1,
		dataHangmanWeb.HangmanWebTranslations.ExtensionDifficultyLevel2,
		dataHangmanWeb.HangmanWebTranslations.ExtensionDifficultyLevel3,
		dataHangmanWeb.HangmanWebTranslations.ExtensionDifficultyLevel4,
		dataHangmanWeb.HangmanWebTranslations.ExtensionDifficultyLevel5,
		dataHangmanWeb.HangmanWebTranslations.ExtensionDifficultyLevel6,
	}

	// Construction de la liste déroulante de niveaux de difficulté si activé
	var difficultyOptions string
	if dataHangmanWeb.DataConfigHangman.EnableDifficulty {
		for i, level := range levels {
			selected := ""
			if i == 2 { // Niveau 3 est sélectionné par défaut
				selected = " selected"
			}
			difficultyOptions += fmt.Sprintf(`<option value="%d"%s>%s</option>`, i+1, selected, level)
		}
	}

	// Formulaire pour l'activation de la difficulté
	dataswitchdifficulty := ""
	if dataHangmanWeb.DataConfigHangman.EnableDifficulty {
		dataswitchdifficulty = fmt.Sprintf(`
            <form action="enable-extension-switchdifficulty" method="post">
                <select name="difficultyLevel">
                    %s
                </select>
            </form>
        `, difficultyOptions)
	}

	// Formulaire principal avec le checkbox
	data := fmt.Sprintf(`
        <form action="enable-extension-enabledifficulty" method="post" class="config-texte-align">
            <label for="enabledifficulty">%s</label>
            <div class="button r" id="button-4">
                <input type="checkbox" id="enabledifficulty" name="enableDifficulty" class="checkbox"%s />
                <div class="knobs"></div>
                <div class="layer"></div>
            </div>
        </form>
        %s
    `, title, ifChecked(dataHangmanWeb.DataConfigHangman.EnableDifficulty), dataswitchdifficulty)

	return template.HTML(data)
}

// Fonction pour renvoyer l'attribut "checked" si activé
func ifChecked(enabled bool) string {
	if !enabled {
		return " checked"
	}
	return ""
}
