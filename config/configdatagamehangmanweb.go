package config

import "html/template"

func ConfigDataGameHangmanWeb() *GameData {
	return &GameData{
		TargetWord:                             "",
		CurrentWord:                            "",
		CurrentWordFormattedHTML:               template.HTML(""),
		DifficultyLevel:                        "",
		NbrVictory:                             0,
		NbrLoose:                               0,
		NbrJokers:                              0,
		NameDifficulty:                         "",
		NameFile:                               "",
		IdDifficulty:                           3,
		StageHangman:                           0,
		StageHangmanFormattedHTML:              template.HTML(""),
		NumOfPossibility:                       10,
		TabSelectedLetter:                      []string{},
		TabSelectedWord:                        []string{},
		PrintNbrVictoryAndLoose:                template.HTML(""),
		PrintInfo:                              template.HTML(""),
		PrintNbrJokers:                         template.HTML(""),
		PrintLetterAndWord:                     template.HTML(""),
		DataInputField:                         template.HTML(""),
		StateInputField:                        true,
		StateInputFieldForAddWord:              false,
		DataInputFieldForAddWord:               template.HTML(""),
		AddButtonForUseJoker:                   template.HTML(""),
		AddFormForTheExtensionLanguage:         template.HTML(""),
		AddFormForTheExtensionVictoryCounter:   template.HTML(""),
		AddFormForTheExtensionEnableDifficulty: template.HTML(""),
		AddFormForTheExtensionAddWordAfterGame: template.HTML(""),
		AddFormForTheExtensionEnableJokers:     template.HTML(""),
	}
}
