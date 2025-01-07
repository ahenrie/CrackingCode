package questions

func AllUniqueChars(inString string) bool {
	seen := make(map[rune]bool)

	for _, char := range inString {
		if seen[char] {
			return false
		}
		seen[char] = true
	}
	return true
}
