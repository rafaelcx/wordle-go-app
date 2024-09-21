package sanitizer

import "strings"

type caseSanitizer struct{}

func (sanitizer caseSanitizer) sanitize(s string) string {
	return strings.ToUpper(s)
}
