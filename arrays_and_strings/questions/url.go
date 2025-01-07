package questions

import "strings"

func UrlMaker(inString string) string {
	inString = strings.TrimRight(inString, " ")
	inString = strings.ReplaceAll(inString, " ", "%20")
	return inString
}
