package main

import (
	"fmt"

	"github.com/CrackingCode/arrays_and_strings/questions"
)

func main() {
	string1 := "abc"
	string2 := "cba"
	string3 := "aabc"
	string4 := "bbc"
	stringURL := "Mr John Smith     "
	PalPerm := "Tact Coa"
	stringCompression := "aabcccddddeffaaaaaa       "

	perm, answer := questions.ArePermutations(string1, string2)
	println(perm, answer)

	perm, answer = questions.ArePermutations(string1, string3)
	println(perm, answer)

	unique := questions.AllUniqueChars(string3)
	println(unique)

	unique = questions.AllUniqueChars(string1)
	println(unique)

	println(questions.UrlMaker(stringURL))

	isPalPerm := questions.PalPerm(PalPerm)
	fmt.Printf("The string: %v yields: %v", PalPerm, isPalPerm)
	println("")
	println(questions.OneWay(string1, string4))

	println(questions.StringCompression(stringCompression))

	questions.FizzBuzz(20)
}
