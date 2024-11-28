package game

import (
	"fmt"
	"hangman-web/config"
	"html/template"
)

// ExtensionDifficulty generates an HTML form with a checkbox to enable or disable
// the difficulty extension and a dropdown menu to select the difficulty level, if enabled.
//
// Parameters:
// - dataHangmanWeb (*config.DataHangmanWeb): a structure containing configuration
//   and translation data for the Hangman-Web application, including the difficulty levels,
//   whether the difficulty extension is enabled, and the current selected difficulty level.
//
// Returns:
// - template.HTML: an HTML block containing a form with a checkbox to enable or disable
//   the difficulty extension, and a dropdown menu to select the difficulty level if the
//   extension is enabled.

func ExtensionDifficulty(dataHangmanWeb *config.DataHangmanWeb) template.HTML {
	// Set the title for the difficulty extension from translations
	title := dataHangmanWeb.HangmanWebTranslations.TitleExtensionEnableDifficulty

	// Define the difficulty levels based on translations
	levels := []string{
		dataHangmanWeb.HangmanWebTranslations.ExtensionDifficultyLevel1,
		dataHangmanWeb.HangmanWebTranslations.ExtensionDifficultyLevel2,
		dataHangmanWeb.HangmanWebTranslations.ExtensionDifficultyLevel3,
		dataHangmanWeb.HangmanWebTranslations.ExtensionDifficultyLevel4,
		dataHangmanWeb.HangmanWebTranslations.ExtensionDifficultyLevel5,
		dataHangmanWeb.HangmanWebTranslations.ExtensionDifficultyLevel6,
	}

	// Initialize the string for the dropdown options
	var difficultyOptions string

	// If difficulty is enabled, build the dropdown options for selecting difficulty levels
	if dataHangmanWeb.DataConfigHangman.EnableDifficulty {
		for i, level := range levels {
			// Check if the current level should be selected based on the game's current difficulty
			selected := ""
			if i == dataHangmanWeb.GameData.IdDifficulty-1 { // Default level 3 is selected (IdDifficulty starts at 1)
				selected = " selected"
			}
			// Add each difficulty level option to the dropdown
			difficultyOptions += fmt.Sprintf(`<option value="%d"%s>%s</option>`, i+1, selected, level)
		}
	}

	// Initialize the HTML for the difficulty dropdown form (only created if difficulty extension is enabled)
	dataswitchdifficulty := ""
	if dataHangmanWeb.DataConfigHangman.EnableDifficulty {
		dataswitchdifficulty = fmt.Sprintf(`
            <form action="enable-extension-switchdifficulty" method="post">
                <select name="difficultyLevel" onchange="this.form.submit()">
                    %s
                </select>
            </form>
        `, difficultyOptions)
	}

	// Generate the HTML for the main form with the checkbox to enable the difficulty extension
	// and include the difficulty dropdown if enabled
	data := fmt.Sprintf(`
        <form action="enable-extension-enabledifficulty" method="post" class="config-texte-align">
            <label for="enabledifficulty">%s</label>
            <div class="button r" id="button-4">
                <input type="checkbox" id="enabledifficulty" name="enableDifficulty" onchange="this.form.submit()" value="NO" class="checkbox"%s  />
                <div class="knobs"></div>
                <div class="layer"></div>
            </div>
        </form>
        %s
    `, title, IfChecked(dataHangmanWeb.DataConfigHangman.EnableDifficulty), dataswitchdifficulty)

	// Return the generated HTML block
	return template.HTML(data)
}
