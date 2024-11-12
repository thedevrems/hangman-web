package manageserver

import (
	"net/http"
	"text/template"
)

// Helper function to parse and execute templates
func ParseAndExecuteTemplate(w http.ResponseWriter, filePath string, data interface{}) {
	tmpl := template.Must(template.ParseFiles(filePath))
	tmpl.Execute(w, data)
}
