package questions

import (
	"fmt"
	"strings"
)

func StringCompression(word string) string {
	word = strings.ReplaceAll(word, " ", "")

	if len(word) == 0 {
		return word
	}

	// Variable to hold compressed string
	var compressed string

	count := 1

	for i := 1; i < len(word); i++ {
		if word[i] == word[i-1] {
			count++
		} else {
			compressed += string(word[i-1]) + fmt.Sprintf("%d", count)
			count = 1
		}
	}
	compressed += string(word[len(word)-1]) + fmt.Sprintf("%d", count)

	if len(compressed) < len(word) {
		return compressed
	}
	return word
}
