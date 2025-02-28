package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"uniq/utils"
)

func main() {
	flags := utils.ParseFlags()

	var outputFile *os.File
	if len(flag.Args()) > 1 {
		var err error
		outputFile, err = os.Create(flag.Args()[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		defer outputFile.Close()
	}

	var input []string
	if len(flag.Args()) > 0 {
		var err error
		input, err = utils.OpenInputFile(flag.Args()[0])
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			input = append(input, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading from stdin:", err)
			return
		}
	}

	outputLines := utils.ProcessFile(input, flags)

	if outputFile != nil {
		for _, line := range outputLines {
			if _, err := fmt.Fprintln(outputFile, line); err != nil {
				fmt.Println(err)
			}
		}
	} else {
		for _, line := range outputLines {
			fmt.Println(line)
		}
	}
}
