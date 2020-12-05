package utils

import (
	"bufio"
	"fmt"
	"os"
)

type ConverterFunc func(line string, rowNumber int) error

func ReadTextFile(fileName string, converterFunc ConverterFunc) error {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("failed to open file %s: %w", fileName, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rowNumber := 0
	for scanner.Scan() {
		line := scanner.Text()
		if err = converterFunc(line, rowNumber); err != nil {
			return err
		}
		rowNumber++
	}
	return nil
}
