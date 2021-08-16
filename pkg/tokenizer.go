package pkg

import (
	"strings"
	"unicode"
)

func Tokenize(text string) []string {
	return strings.FieldsFunc(text, func(r rune) bool {
		// Split on any character that is not a letter or a number.
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
}

func Analyze(text string) []string {
	tokens := Tokenize(text)
	tokens = LowercaseFilter(tokens)
	tokens = StopwordFilter(tokens)
	tokens = StemmerFilter(tokens)
	return tokens
}
