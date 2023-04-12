package pkg

import "fmt"

// function to print symbol
func PrintChar(s []string, char map[rune][]string) {
	for _, val := range s {
		// if val != ""
		if len(val) != 0 {
			// 8 line
			for i := 0; i < 8; i++ {
				// print symbol from map with key
				for _, symb := range val {
					fmt.Print(char[symb][i])
				}
				fmt.Println()
			}
			// if two or more new line
		} else if len(val) == 0 {
			fmt.Println()
		}
	}
}
