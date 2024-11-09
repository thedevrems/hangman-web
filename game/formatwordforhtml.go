package game

import (
	"html/template"
	"strings"
)

// FormatWordForHTML formats a given word for display in HTML, wrapping each letter in a span element.
// If the letter is an underscore (indicating a letter that has not yet been guessed),
// it also wraps the underscore in a span element for consistent styling.
//
// Parameters:
// - word: a string representing the word to format for HTML display.
//
// Returns:
//   - template.HTML: the formatted HTML representation of the word, where each letter is enclosed in
//     a span with the class "letter" for styling purposes.
func FormatWordForHTML(word string) template.HTML {
	var formattedWord strings.Builder

	// Iterate through each character in the word.
	for _, letter := range word {
		// If the letter is not an underscore, wrap it in a span element.
		if letter != '_' {
			formattedWord.WriteString(`<span class="letter">` + string(letter) + `</span>`)
		} else {
			// If the letter is an underscore, wrap it in a span element as well.
			formattedWord.WriteString(`<span class="letter">_</span>`)
		}
	}
	// Return the constructed HTML as a template.HTML type for safe rendering.
	return template.HTML(formattedWord.String())
}
