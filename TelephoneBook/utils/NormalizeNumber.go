package utils

import (
	"strings"
	"unicode"
)

func NormalizeNumber(rawNumber string) (string, bool) {
	clearNumber := ""
	for _, char := range rawNumber {
		if char == '+' {
			if clearNumber != "" {
				return "", false
			}
			clearNumber += string(char)
			continue
		}
		if char == '*' || char == '#' {
			if clearNumber != "" && !(strings.Contains(clearNumber, "*") || strings.Contains(clearNumber, "#")) {
				return "", false
			}
			clearNumber += string(char)
			continue
		}
		if char == '(' || char == ')' || char == '-' || char == ' ' {
			continue
		}
		if !unicode.IsDigit(char) {
			return "", false
		}
		clearNumber += string(char)
	}
	if (clearNumber[0] == '#' || clearNumber[0] == '*') && !(clearNumber[len(clearNumber)-1] == '#' || clearNumber[len(clearNumber)-1] == '*') {
		return "", false
	}
	return clearNumber, true
}
