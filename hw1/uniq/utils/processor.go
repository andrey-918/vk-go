package utils

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func ProcessFile(input io.Reader, output io.Writer, countFlag, duplicatesFlag, uniqueFlag, ignoreCase bool, fieldCount, charCount int) {
	scanner := bufio.NewScanner(input)
	if (countFlag && duplicatesFlag) || (countFlag && uniqueFlag) || (duplicatesFlag && uniqueFlag) {
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
		if ignoreCase {
			line = strings.ToLower(line)
		}

		fields := strings.Fields(line)

		if fieldCount > 0 && len(fields) > fieldCount {
			line = strings.Join(fields[fieldCount:], " ")
		} else if fieldCount > 0 {
			line = ""
		}

		if charCount > 0 && len(line) > charCount {
			line = line[charCount:]
		}

		if curIn == 0 {
			origLast = origCur
		}
		if line != lastLine && curIn != 0 {
			if countFlag {
				PrintLine(output, fmt.Sprintf("%d %s", curIn, origLast))
			} else if duplicatesFlag {
				if curIn > 1 {
					PrintLine(output, origLast)
				}
			} else if uniqueFlag {
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

	if curIn > 0 { // Обработка последней строки
		if countFlag {
			PrintLine(output, fmt.Sprintf("%d %s", curIn, origLast))
		} else if duplicatesFlag {
			if curIn > 1 {
				PrintLine(output, origLast)
			}
		} else if uniqueFlag {
			if curIn == 1 {
				PrintLine(output, origLast)
			}
		} else {
			PrintLine(output, origLast)
		}
	}

	if err := scanner.Err(); err != nil {
		PrintLine(output, "Error reading input: "+err.Error())
		return
	}
}
