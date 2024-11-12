package game

import (
	"hangman-web/config"
	"html/template"
)

func PrintInfo(dataHangmanWeb *config.DataHangmanWeb, message string) template.HTML {
	return template.HTML("<p>" + message + "</p>")
}
