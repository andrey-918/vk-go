package utils

import (
	"fmt"
	"io"
	"os"
)

func OpenInputFile(filename string) (io.Reader, error) {
	inputFile, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	return inputFile, nil
}

func CreateOutputFile(filename string) (io.Writer, error) {
	outputFile, err := os.Create(filename)
	if err != nil {
		return nil, fmt.Errorf("error creating output file: %w", err)
	}
	return outputFile, nil
}

func PrintLine(output io.Writer, line string) {
	if output != nil {
		fmt.Fprintln(output, line)
	} else {
		fmt.Println(line)
	}
}
