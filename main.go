package main

import (
	"fmt"

	"github.com/CrackingCode/arrays_and_strings/questions"
)

func main() {
	string1 := "abc"
	string2 := "cba"
	string3 := "aabc"
	stringURL := "Mr John Smith     "
	PalPerm := "Tact Coa"

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

}
