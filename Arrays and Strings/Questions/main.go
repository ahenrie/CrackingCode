package main

import "fmt"

func main() {
	testString1 := "aabc"
	testString2 := "abc"
	testString3 := "bca"

	fmt.Println(allUniqueChars(testString1))
	fmt.Println(allUniqueChars(testString2))
	fmt.Println(arePermutations(testString2, testString3))
	fmt.Println(arePermutations(testString2, testString1))
}
