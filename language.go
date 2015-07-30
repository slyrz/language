// Package language provides primitives to detect the language of text.
package language

// Language stores information about the language of the analyzed text.
type Language struct {
	Code     string // ISO 639 language code like "en", "de", ...
	Name     string // human-readable language name
	Reliable bool   // true if the detection is reliable
}

// Detect detects the language of text.
func Detect(text string) *Language {
	return detectLanguage(text, true)
}

// DetectHTML detects the language of text that might contain HTML tags and
// entities.
func DetectHTML(text string) *Language {
	return detectLanguage(text, false)
}
