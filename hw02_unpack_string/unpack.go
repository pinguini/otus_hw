package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var prev rune
	var b strings.Builder

	for i, r := range s {
		if unicode.IsDigit(r) {
			// got digit write previous letter int(r-'0') times, if previous is Letter
			if !unicode.IsLetter(prev) {
				return "", ErrInvalidString
			}
			b.WriteString(strings.Repeat(string(prev), int(r-'0')))
			prev = r
		} else {
			// got letter write previous if exists to string
			if unicode.IsLetter(prev) && i > 0 {
				b.WriteString(string(prev))
			}
			prev = r
		}
	}
	// so we have last rune in "prev" need to write to string
	if unicode.IsLetter(prev) {
		b.WriteString(string(prev))
	}
	return b.String(), nil
}