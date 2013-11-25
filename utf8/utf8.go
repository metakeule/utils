package utf8

import (
	"strings"
	"unicode/utf8"
)

func pad(s string, no int, padWith string) (pads string) {
	l := utf8.RuneCountInString(s)
	if l >= no {
		return
	}
	return strings.Repeat(padWith, no-l)
}

func RPad(s string, no int) string { return s + pad(s, no, " ") }
func LPad(s string, no int) string { return pad(s, no, " ") + s }

func RPadWith(s string, no int, padWith string) string { return s + pad(s, no, padWith) }
func LPadWith(s string, no int, padWith string) string { return pad(s, no, padWith) + s }
