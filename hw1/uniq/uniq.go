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

	var output func([]string) error
	if len(flag.Args()) > 1 {
		var err error
		output, err = utils.CreateOutputFile(flag.Args()[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		defer func() {
			if err := output(nil); err != nil {
				fmt.Println(err)
			}
		}()
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

	if output != nil {
		if err := output(outputLines); err != nil {
			fmt.Println(err)
		}
	} else {
		for _, line := range outputLines {
			fmt.Println(line)
		}
	}
}
