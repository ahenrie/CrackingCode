package questions

import (
	"fmt"
	"slices"
	"sort"
)

func ArePermutations(string1 string, string2 string) (bool, string) {
	// check length
	if len(string1) != len(string2) {
		return false, "The strings were different lengths"
	}

	slice1 := []rune{}
	slice2 := []rune{}

	//convert strings to slices
	for _, runes := range string1 {
		slice1 = append(slice1, runes)
	}
	for _, runes := range string2 {
		slice2 = append(slice2, runes)
	}

	fmt.Println(slice1)
	fmt.Println(slice2)

	// sort
	sort.Slice(slice1, func(i, j int) bool {
		return slice1[i] < slice1[j]
	})
	sort.Slice(slice2, func(i, j int) bool {
		return slice2[i] < slice2[j]
	})

	fmt.Println(slice1)
	fmt.Println(slice2)

	return slices.Equal(slice1, slice2), "\nHopefully they are equal"
}
