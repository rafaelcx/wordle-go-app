package sanitizer

type sanitizer interface {
	sanitize(s string) string
}

func Sanitize(s string) string {
	for _, sl := range getSanitizerList() {
		s = sl.sanitize(s)
	}
	return s
}

func getSanitizerList() []sanitizer {
	return []sanitizer{
		caseSanitizer{},
		newLinesSanitizer{},
		emptySpacesSanitizer{},
	}
}
