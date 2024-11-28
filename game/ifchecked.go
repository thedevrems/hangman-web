package game

// IfChecked returns the "checked" attribute for a checkbox input based on the provided boolean value.
//
// Parameters:
// - enabled (bool): A boolean value that determines if the checkbox should be checked or unchecked.
//   - If `enabled` is `true`, it returns an empty string, meaning the checkbox is not checked.
//   - If `enabled` is `false`, it returns the string " checked", meaning the checkbox is checked.
//
// Returns:
// - string: The HTML "checked" attribute (or an empty string) based on the `enabled` value.
func IfChecked(enabled bool) string {
	if !enabled {
		return " checked" // Return "checked" if the feature is not enabled
	}
	return "" // Return an empty string if the feature is enabled
}
