package pkg

import (
	"fmt"
	"os"
)

// read text file and write to []byte
func ReadTextFile(fileName string) ([]byte, error) {
	text, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}
	return text, nil
}
