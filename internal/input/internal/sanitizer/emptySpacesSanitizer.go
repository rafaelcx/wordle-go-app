package sanitizer

import "strings"

type emptySpacesSanitizer struct{}

func (sanitizer emptySpacesSanitizer) sanitize(s string) string {
	return strings.ReplaceAll(s, " ", "")
}
