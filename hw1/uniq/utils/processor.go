package utils

import (
	"fmt"
	"strings"
)

func ProcessFile(input []string, flags Flags) []string {
	var output []string

	if (flags.CountFlag && flags.DuplicatesFlag) || (flags.CountFlag && flags.UniqueFlag) || (flags.DuplicatesFlag && flags.UniqueFlag) {
		output = append(output, "Error: Options -c, -d, and -u are mutually exclusive.")
		return output
	}

	lastLine := ""
	curIn := 0
	origLast := ""
	origCur := ""

	for _, line := range input {
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
				output = append(output, fmt.Sprintf("%d %s", curIn, origLast))
			} else if flags.DuplicatesFlag {
				if curIn > 1 {
					output = append(output, origLast)
				}
			} else if flags.UniqueFlag {
				if curIn == 1 {
					output = append(output, origLast)
				}
			} else {
				output = append(output, origLast)
			}
			curIn = 0
			origLast = origCur
		}

		curIn++
		lastLine = line
	}

	if curIn > 0 { // Обработка последней строки
		if flags.CountFlag {
			output = append(output, fmt.Sprintf("%d %s", curIn, origLast))
		} else if flags.DuplicatesFlag {
			if curIn > 1 {
				output = append(output, origLast)
			}
		} else if flags.UniqueFlag {
			if curIn == 1 {
				output = append(output, origLast)
			}
		} else {
			output = append(output, origLast)
		}
	}

	return output
}
