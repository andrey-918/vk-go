package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func OpenInputFile(filename string) ([]string, error) {
	inputFile, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer inputFile.Close()

	var lines []string
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading input: %w", err)
	}

	return lines, nil
}

func CreateOutputFile(filename string) (func([]string) error, error) {
	outputFile, err := os.Create(filename)
	if err != nil {
		return nil, fmt.Errorf("error creating output file: %w", err)
	}

	return func(lines []string) error {
		for _, line := range lines {
			if _, err := fmt.Fprintln(outputFile, line); err != nil {
				return fmt.Errorf("error writing to output file: %w", err)
			}
		}
		return outputFile.Close()
	}, nil
}

func PrintLine(output io.Writer, line string) {
	if output != nil {
		fmt.Fprintln(output, line)
	} else {
		fmt.Println(line)
	}
}
