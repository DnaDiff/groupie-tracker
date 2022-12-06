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
	if strings.Contains(string(chars), "Usa") {
		chars = []rune(strings.ReplaceAll(string(chars), "Usa", "USA"))
	}

	return string(chars)
}
