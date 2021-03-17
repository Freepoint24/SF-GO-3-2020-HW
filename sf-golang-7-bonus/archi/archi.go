package archi

import (
	"strconv"
	"strings"
)

const escapeSymbol = '\\'

// todo
func Compress(input string) string {
	var counter uint64
	var current rune
	var output strings.Builder

	var flushCounter = func() {
		if counter > 1 {
			output.WriteString(strconv.FormatUint(counter, 10))
		}
	}

	for _, cursor := range input {
		switch {
		default:
			flushCounter()
			fallthrough
		case counter == 0:
			if (cursor >= '0' && cursor <= '9') || (cursor == escapeSymbol) {
				output.WriteRune(escapeSymbol)
			}
			output.WriteRune(cursor)
			counter = 1
			current = cursor
		case cursor == current:
			counter++
		}

	}
	flushCounter()
	return output.String()
}

// todo
func Decompress(input string) string {

	var escaped, single bool
	var counter, i uint64
	var current rune
	var output strings.Builder

	var flushCurrent = func() {
		for i = 0; i < counter; i++ {
			output.WriteRune(current)
		}
	}

	for _, cursor := range input {
		switch {
		case cursor == escapeSymbol:
			escaped = true
			flushCurrent()
		case escaped:
			escaped = false
			current = cursor
			counter = 1
			single = true
		case cursor >= '0' && cursor <= '9':
			if single {
				counter = 0
				single = false
			}
			counter = counter*10 + uint64(cursor-'0')
		default:
			flushCurrent()
			counter = 1
			single = true
			current = cursor
		}
	}
	flushCurrent()
	return output.String()
}
