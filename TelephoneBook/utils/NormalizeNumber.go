package utils

import (
	"strings"
	"unicode"
)

func NormalizeNumber(rawNumber []rune) (string, bool) {
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
			if clearNumber != "" && (!strings.Contains(clearNumber, "*") || !strings.Contains(clearNumber, "#")) {
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
		// TODO: Нужно нормализовать до читемого варианта. С этим поможет "://github.com" функция phonenumbers.Format
		clearNumber += string(char)
	}
	return clearNumber, true
}
