package manageserver

import "net/http"

// Helper function for redirecting
func Redirect(w http.ResponseWriter, r *http.Request, path string) {
	http.Redirect(w, r, path, http.StatusSeeOther)
}
