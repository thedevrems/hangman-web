package manageserver

import (
	"net/http"
	"text/template"
)

// ParseAndExecuteTemplate parses an HTML template file and executes it with the provided data.
//
// This function is responsible for loading a template file from the given file path, and then rendering it
// by applying the provided data (usually a data structure containing game state or settings) to the template.
// The rendered HTML is then sent to the response writer to be returned to the client.
//
// Parameters:
// - w (http.ResponseWriter): The response writer to send the rendered template to the client.
// - filePath (string): The path to the template file that needs to be parsed and executed.
// - data (interface{}): The data that will be used to populate the template (e.g., game state, configuration).
//
// Behavior:
// - The template at the specified file path is parsed and executed.
// - If there is an error parsing or executing the template, the application will panic (as it's a critical operation).
//
// Example Usage:
//
//	ParseAndExecuteTemplate(w, "templates/index.html", data)
func ParseAndExecuteTemplate(w http.ResponseWriter, filePath string, data interface{}) {
	tmpl := template.Must(template.ParseFiles(filePath))
	tmpl.Execute(w, data)
}
