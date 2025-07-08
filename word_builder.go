package main

import (
	"strings"
)

type wordBuilder struct {
	strings.Builder
}

func (buf *wordBuilder) writeWord(word string) {
	if buf.Len() != 0 {
		buf.WriteRune(' ')
	}
	buf.WriteString(word)
}
