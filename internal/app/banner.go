package app

import (
	"fmt"
	"os"
	"strings"
)

func ReadBannerFromFile() ([]byte, error) {
	text, err := os.ReadFile("banner/standard.txt")
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}
	return text, nil
}

func BannerToMap(s []byte) map[rune][]string {
	mapOfBanner := make(map[rune][]string)
	temp := []string{}
	firstSymbol := ' '
	text := strings.Split(string(s), "\n")

	for _, val := range text {
		if len(val) != 0 {
			temp = append(temp, val)
		} else if len(val) == 0 && len(temp) != 0 {
			mapOfBanner[firstSymbol] = temp
			firstSymbol++
			temp = nil
		}
	}
	return mapOfBanner
}
