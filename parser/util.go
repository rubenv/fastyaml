package parser

import "strings"

func ReadUntil(in string, offset int, c rune) string {
	pos := strings.IndexRune(in[offset:], c)
	if pos < 0 {
		return ""
	}
	return in[offset : offset+pos]
}
