package questions

import (
	"fmt"
	"strings"
)

type RuneInfo struct {
	Positions []int
	Count     int
}

func OneWay(word1, word2 string) bool {
	word1 = strings.ToLower(word1)
	word2 = strings.ToLower(word2)

	// Check lengths of strings
	if (len(word1) - len(word2)) > 1 {
		return false
	}

	// Which is shorter?
	// Word 1 should be shorter for organization
	if len(word1) > len(word2) {
		word1, word2 = word2, word1
	}

	foundDifferences := false
	i, j := 0, 0

	for i < len(word1) && j < len(word2) {
		if word1[i] != word2[j] {
			if foundDifferences {
				return false
			}
			foundDifferences = true

			if len(word1) == len(word2) {
				i++
			}
		} else {
			i++
		}
		j++
	}
	return true
}

func PrintMap(runeMap map[rune]int) {
	for char, count := range runeMap {
		fmt.Printf("Key: %c --> %d\n", char, count)
	}
}
