package manageserver

import "net/http"

// Redirect performs an HTTP redirect to the specified path with a "See Other" status code (303).
//
// This function is used to redirect the client to another URL, typically after handling a form submission or
// completing a game action. It sends an HTTP response with a `303 See Other` status code, which tells the client
// to perform a GET request to the specified path.
//
// Parameters:
// - w (http.ResponseWriter): The response writer used to send the redirect response.
// - r (http.Request): The incoming HTTP request, which is needed to handle the redirection properly.
// - path (string): The URL path to redirect the client to.
//
// Behavior:
// - The client is redirected to the provided `path` with a status code `303 See Other`.
// - This is useful for avoiding form resubmission errors after a POST request.
//
// Example Usage:
//   Redirect(w, r, "/game.html")
func Redirect(w http.ResponseWriter, r *http.Request, path string) {
	http.Redirect(w, r, path, http.StatusSeeOther)
}
