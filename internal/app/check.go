package app

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func CheckArgument() (string, error) {
	numOfArgument := 1

	arg := os.Args[numOfArgument:]
	if len(arg) != numOfArgument {
		return "", errors.New("argument not equal 1")
	}
	argument := arg[0]

	for _, val := range argument {
		if !(val >= 0 && val <= '~') {
			return "", errors.New("not ASCII characters")
		}
	}

	return argument, nil
}

func CheckEmptyString(inputData string) ([]string, error) {
	inputSlice := strings.Split(inputData, "\n")

	countOfEmptyString := 0

	for _, val := range inputSlice {
		if val == "" {
			countOfEmptyString++
		}
	}
	if countOfEmptyString == len(inputSlice) {
		for i := 0; i < len(inputSlice)-1; i++ {
			fmt.Println()
		}
		return nil, errors.New("empty string")
	}
	return inputSlice, nil
}
