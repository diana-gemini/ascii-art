package main

import (
	"ascii-art/pkg"
	"fmt"
	"os"
	"strings"
)

func main() {
	numOfArgument := 1
	// argument from terminal
	arg := os.Args[numOfArgument:]
	// if len of arg not equal 1, return
	if len(arg) != numOfArgument {
		fmt.Println("Argument not equal 1")
		return
	}
	// input data from arg, input data is string
	inputData := arg[0]

	// if inputData contains not Ascii characters
	for _, val := range inputData {
		if !(val >= 0 && val <= '~') {
			fmt.Println("Not ASCII characters")
			return
		}
	}

	// replace \\n to \n, output result is string
	newInputData := strings.ReplaceAll(inputData, "\\n", "\n")

	// if input data is "", return
	if len(newInputData) == 0 {
		return
	}

	// if input data is new line, print new line then return
	if newInputData == "\n" {
		fmt.Println()
		return
	}

	// split string with new line (\n), output result is slice
	inputSlice := strings.Split(newInputData, "\n")

	count := 0

	// if len of values is 0, then sum count
	for _, val := range inputSlice {
		if val == "" {
			count++
		}
	}
	// if count equal len of inputSlice, then print \n
	if count == len(inputSlice) {
		for i := 0; i < len(inputSlice)-1; i++ {
			fmt.Println()
		}
		return
	}

	// call ReadTextFile function, return text from *.txt in []byte
	text, err := pkg.ReadTextFile("banner/standard.txt")
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}

	// call MakeMap function, output result is map[rune][]string
	charMap := pkg.MakeMap(text)

	// call PrintChar function, print symbol
	pkg.PrintChar(inputSlice, charMap)
}
