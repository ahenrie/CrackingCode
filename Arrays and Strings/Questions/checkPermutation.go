package main

import (
	"slices"
	"sort"
)

func arePermutations(string1 string, string2 string) bool {
	// check length
	if len(string1) != len(string2) {
		return false
	}

	slice1 := []rune{}
	slice2 := []rune{}

	//convert strings to slices
	for _, runes := range string1 {
		slice1 = append(slice1, runes)
	}
	for _, runes := range string2 {
		slice1 = append(slice2, runes)
	}

	sort.Slice(slice1, func(i, j int) bool {
		return slice1[i] < slice1[j]
	})

	sort.Slice(slice2, func(i, j int) bool {
		return slice2[i] < slice2[j]
	})

	return slices.Equal(slice1, slice2)
}
