package app

import "fmt"

func PrintChar(s []string, char map[rune][]string) {
	for _, val := range s {
		if len(val) != 0 {
			for i := 0; i < 8; i++ {
				for _, symb := range val {
					fmt.Print(char[symb][i])
				}
				fmt.Println()
			}
		} else if len(val) == 0 {
			fmt.Println()
		}
	}
}
