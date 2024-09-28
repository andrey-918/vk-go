package utils

import (
	"flag"
)

func ParseFlags() (bool, bool, bool, bool, int, int) {
	countFlag := flag.Bool("c", false, "count consecutive occurrences of each line")
	duplicatesFlag := flag.Bool("d", false, "only print duplicate lines")
	uniqueFlag := flag.Bool("u", false, "only print unique lines")
	ignoreCase := flag.Bool("i", false, "ignore case differences")
	fieldCount := flag.Int("f", 0, "ignore the first num_fields fields")
	charCount := flag.Int("s", 0, "ignore the first num_chars characters")

	flag.Parse()

	return *countFlag, *duplicatesFlag, *uniqueFlag, *ignoreCase, *fieldCount, *charCount
}
