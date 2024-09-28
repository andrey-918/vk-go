package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"uniq/utils"
)

func main() {
	countFlag, duplicatesFlag, uniqueFlag, ignoreCase, fieldCount, charCount := utils.ParseFlags()

	var output io.Writer
	if len(flag.Args()) > 1 {
		var err error
		output, err = utils.CreateOutputFile(flag.Args()[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		defer output.(io.Closer).Close()
	} else {
		output = os.Stdout // Вывод в консоль
	}

	var input io.Reader
	if len(flag.Args()) > 0 {
		var err error
		input, err = utils.OpenInputFile(flag.Args()[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		defer input.(io.Closer).Close()
	} else {
		input = os.Stdin
	}

	utils.ProcessFile(input, output, countFlag, duplicatesFlag, uniqueFlag, ignoreCase, fieldCount, charCount)
}
