package tools

import (
	"strings"
	"unicode"
)

func Title(s string) string {
	chars := []rune(s)

	for i, r := range chars {
		if i == 0 || chars[i-1] == ' ' || chars[i-1] == '"' {
			chars[i] = unicode.ToTitle(r)
		}
	}
	// Usa -> USA
	abbreviations := []string{"Usa", "Uk"}
	for _, e := range abbreviations {
		if strings.Contains(string(chars), e) {
			chars = []rune(strings.ReplaceAll(string(chars), e, strings.ToUpper(e)))
		}
	}

	return string(chars)
}
