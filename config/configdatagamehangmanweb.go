package config

import "html/template"

func ConfigDataGameHangmanWeb() *GameData {
	return &GameData{
		TargetWord:                "",
		CurrentWord:               "",
		CurrentWordFormattedHTML:  template.HTML(""),
		DifficultyLevel:           "",
		NbrVictory:                0,
		NbrLoose:                  0,
		NbrJokers:                 0,
		NameDifficulty:            "",
		StageHangman:              0,
		StageHangmanFormattedHTML: template.HTML(""),
		NumOfPossibility:          10,
		TabSelectedLetter:         []string{},
		TabSelectedWord:           []string{},
		PrintNbrVictoryAndLoose:   template.HTML(""),
		PrintInfo:                 template.HTML(""),
		PrintNbrJokers:            template.HTML(""),
		PrintLetterAndWord:        template.HTML(""),
		DataInputField:            template.HTML(""),
		AddButtonForUseJoker:      template.HTML(""),
	}
}
