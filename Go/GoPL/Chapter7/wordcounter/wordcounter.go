package wordcounter

import (
	"bufio"
	"bytes"
	"fmt"
	"unicode"
	"unicode/utf8"
)

type WordCounter int

func ScanWords(data []byte, atEOF bool) (advance int, token []byte, err error) {
	start := 0
	for width := 0; start < len(data); start += width {
		var r rune
		r, width = utf8.DecodeRune(data[start:])
		if !unicode.IsSpace(r) && !unicode.IsPunct(r) {
			break
		}
	}

	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])
		// Scan until space or punct
		if unicode.IsSpace(r) || unicode.IsPunct(r) {
			return i + width, data[start:i], nil
		}
	}
	// Check that it is
	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}

	return start, nil, nil
}

func (c *WordCounter) Write(p []byte) (int, error) {
	s := bufio.NewScanner(bytes.NewBuffer(p))

	s.Split(ScanWords)
	for s.Scan() {
		*c++
	}

	return len(p), s.Err()
}

func (c *WordCounter) String() string {
	return fmt.Sprintf("%d word(s)", *c)
}

type LineCounter int

func (l *LineCounter) Write(p []byte) (int, error) {
	s := bufio.NewScanner(bytes.NewBuffer(p))

	for s.Scan() {
		*l++
	}
	return len(p), s.Err()
}

func (l *LineCounter) String() string {
	return fmt.Sprintf("%d line(s)", *l)
}
