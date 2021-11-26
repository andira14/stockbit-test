package main

import (
	"strings"
)

func findFirstStringInBracket(str string) string {
	if len(str) > 0 {

		indexFirstBracketFound := strings.Index(str, "(")
		indexClosingBracketFound := strings.Index(str, ")")

		// Will skip if opening or closing brackets not found
		if indexFirstBracketFound >= 0 && indexClosingBracketFound >= 1 {
			runes := []rune(str)
			return string(runes[indexFirstBracketFound+1 : indexClosingBracketFound])
		}
	}
	return ""
}
