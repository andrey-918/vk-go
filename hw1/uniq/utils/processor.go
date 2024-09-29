package utils

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func ProcessFile(input io.Reader, output io.Writer, flags Flags) {
	scanner := bufio.NewScanner(input)

	if (flags.CountFlag && flags.DuplicatesFlag) || (flags.CountFlag && flags.UniqueFlag) || (flags.DuplicatesFlag && flags.UniqueFlag) {
		PrintLine(output, "Error: Options -c, -d, and -u are mutually exclusive.")
		return
	}

	lastLine := ""
	curIn := 0
	origLast := ""
	origCur := ""

	for scanner.Scan() {
		line := scanner.Text()
		origCur = line
		if flags.IgnoreCase {
			line = strings.ToLower(line)
		}

		fields := strings.Fields(line)

		if flags.FieldCount > 0 && len(fields) > flags.FieldCount {
			line = strings.Join(fields[flags.FieldCount:], " ")
		} else if flags.FieldCount > 0 {
			line = ""
		}

		if flags.CharCount > 0 && len(line) > flags.CharCount {
			line = line[flags.CharCount:]
		}

		if curIn == 0 {
			origLast = origCur
		}
		if line != lastLine && curIn != 0 {
			if flags.CountFlag {
				PrintLine(output, fmt.Sprintf("%d %s", curIn, origLast))
			} else if flags.DuplicatesFlag {
				if curIn > 1 {
					PrintLine(output, origLast)
				}
			} else if flags.UniqueFlag {
				if curIn == 1 {
					PrintLine(output, origLast)
				}
			} else {
				PrintLine(output, origLast)
			}
			curIn = 0
			origLast = origCur
		}

		curIn++
		lastLine = line
	}

	if err := scanner.Err(); err != nil {
		PrintLine(output, "Error reading input: "+err.Error())
		return
	}

	if curIn > 0 { // Обработка последней строки
		if flags.CountFlag {
			PrintLine(output, fmt.Sprintf("%d %s", curIn, origLast))
		} else if flags.DuplicatesFlag {
			if curIn > 1 {
				PrintLine(output, origLast)
			}
		} else if flags.UniqueFlag {
			if curIn == 1 {
				PrintLine(output, origLast)
			}
		} else {
			PrintLine(output, origLast)
		}
	}

}
