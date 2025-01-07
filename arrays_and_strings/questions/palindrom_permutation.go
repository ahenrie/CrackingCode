package questions

import "strings"

func PalPerm(inString string) (bool, []string) {

	// Remove space and lowercase
	inString = strings.ReplaceAll(inString, " ", "")
	inString = strings.ToLower(inString)

}
