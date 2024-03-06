package main

import (
	app "ascii-art/internal/app"
	"fmt"
	"strings"
)

func main() {
	argument, err := app.CheckArgument()
	if err != nil {
		fmt.Println(err)
		return
	}

	inputData := strings.ReplaceAll(argument, "\\n", "\n")

	if len(inputData) == 0 {
		return
	}

	if inputData == "\n" {
		fmt.Println()
		return
	}

	inputSlice, err := app.CheckEmptyString(inputData)
	if err != nil {
		return
	}

	text, err := app.ReadBannerFromFile()
	if err != nil {
		fmt.Println(err)
		return
	}

	mapOfBanner := app.BannerToMap(text)

	app.PrintChar(inputSlice, mapOfBanner)
}
