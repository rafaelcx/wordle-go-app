package sanitizer

import "strings"

type newLinesSanitizer struct{}

func (sanitizer newLinesSanitizer) sanitize(s string) string {
	return strings.TrimSuffix(s, "\n")
}
