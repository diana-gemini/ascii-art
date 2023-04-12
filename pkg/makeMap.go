package pkg

import (
	"strings"
)

func MakeMap(s []byte) map[rune][]string {
	char := make(map[rune][]string)
	tempArr := []string{}
	firstSymbol := ' '
	// split string with \n""
	text := strings.Split(string(s), "\n")

	for _, val := range text {
		if len(val) != 0 {
			// append to array
			tempArr = append(tempArr, val)
		} else if len(val) == 0 && len(tempArr) != 0 {
			// write to map with key
			char[firstSymbol] = tempArr
			firstSymbol++
			tempArr = nil
		}
	}
	return char
}
