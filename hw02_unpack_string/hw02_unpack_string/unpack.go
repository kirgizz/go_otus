package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")
var systemSymbol = "n"
var slashSymbol = "\\"

func Unpack(checkString string) (string, error) {
	if len(checkString) == 0 {
		return "", nil
	}
	checkWord := []rune(checkString)
	if unicode.IsDigit(checkWord[0]) {
		return "", ErrInvalidString
	}
	builder := strings.Builder{}
	for i, r := range checkWord {
		// double int characters - error
		if len(checkWord) >= i+2 && unicode.IsDigit(checkWord[i+1]) && unicode.IsDigit(r) {
			return "", ErrInvalidString
		}
		// unpack string
		if repeatNumber, err := strconv.Atoi(string(checkWord[i])); err == nil {
			symbol := string(checkWord[i-1])
			if symbol == systemSymbol && string(checkWord[i-2]) == slashSymbol {
				builder.WriteString(strings.Repeat(slashSymbol+symbol, repeatNumber))
				continue
			}
			if repeatNumber == 0 {
				continue
			}
			builder.WriteString(strings.Repeat(symbol, repeatNumber-1))
			continue
		}
		// trim character before zero
		if len(checkWord) >= i+2 && (string(checkWord[i+1]) == "0") {
			continue
		}
		// by default add character
		builder.WriteRune(r)
	}
	return builder.String(), nil
}
