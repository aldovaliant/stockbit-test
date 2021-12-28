package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(findFirstStringInBracket("a(aldo)t(y)u"))
}

func findFirstStringInBracket(str string) string {
	if len(str) >= 2 { // at least have length >= 2 for the bracket
		indexFirstBracketFound := strings.Index(str, "(")
		if indexFirstBracketFound >= 0 {
			indexClosingBracketFound := strings.Index(str, ")")
			if indexClosingBracketFound >= 0 {
				runes := []rune(str)
				return string(runes[indexFirstBracketFound+1 : indexClosingBracketFound])
			}
		}
	}
	return ""
}
