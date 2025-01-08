package questions

import "strings"

func PalPerm(inString string) bool {

	// Remove space and lowercase
	inString = strings.ReplaceAll(inString, " ", "")
	inString = strings.ToLower(inString)

	charCount := make(map[rune]int)

	for _, char := range inString {
		charCount[char]++
	}

	oddCount := 0
	for _, count := range charCount {
		if count%2 != 0 {
			oddCount++
			if oddCount > 1 {
				return false
			}
		}
	}

	return true
}
