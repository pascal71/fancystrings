// fancystrings/fancystrings.go
package fancystrings

// InsertSpaces inserts a space after each character in the input string.
func InsertSpaces(s string) string {
	if s == "" {
		return s
	}

	var spacedStr string
	for _, r := range s {
		spacedStr += string(r) + " "
	}

	return spacedStr[:len(spacedStr)-1]
}
