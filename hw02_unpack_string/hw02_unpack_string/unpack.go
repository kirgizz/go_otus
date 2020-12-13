package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")
var systemSymbol = "n"

func Unpack(checkString string) (string, error) {
	if len(checkString) == 0 {
		return "", nil
	}

	if unicode.IsDigit(rune(checkString[0])) {
		return "", ErrInvalidString
	}
	builder := strings.Builder{}
	for i, r := range checkString {
		if unicode.IsDigit(r) {
			//double int characters - error
			if len(checkString) >= i+1 && unicode.IsDigit(rune(checkString[i+1])) {
				return "", ErrInvalidString
			}
			repeatNumber, err := strconv.Atoi(string(r))

			if err != nil || repeatNumber < 0 {
				return "", ErrInvalidString
			}
			symbol := string((checkString[i-1]))
			if symbol == systemSymbol {
				if string(checkString[i-2]) == "\\" {
					builder.WriteString(strings.Repeat("\\"+systemSymbol, repeatNumber))
					continue
				}
			}
			for j := 0; j < repeatNumber-1; j++ {
				builder.WriteString(symbol)
			}
			continue
		}
		//trim character before zero
		if len(checkString) >= i+2 && (string(checkString[i+1]) == "0") {
			continue
		}
		//by default add character
		builder.WriteRune(r)
	}
	return builder.String(), nil
}
