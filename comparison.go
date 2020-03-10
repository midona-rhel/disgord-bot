package disgordbot

import "strings"

// compareString compares the equality of two strings ignoring capitalization
// and trailing whitespaces
func compareString(s, t string) bool {
	s = strings.TrimSpace(s)
	t = strings.TrimSpace(t)
	return strings.EqualFold(s, t)
}
