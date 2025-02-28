package utils

import (
	"flag"
)

type Flags struct {
	CountFlag      bool
	DuplicatesFlag bool
	UniqueFlag     bool
	IgnoreCase     bool
	FieldCount     int
	CharCount      int
}

func ParseFlags() Flags {
	countFlag := flag.Bool("c", false, "count consecutive occurrences of each line")
	duplicatesFlag := flag.Bool("d", false, "only print duplicate lines")
	uniqueFlag := flag.Bool("u", false, "only print unique lines")
	ignoreCase := flag.Bool("i", false, "ignore case differences")
	fieldCount := flag.Int("f", 0, "ignore the first num_fields fields")
	charCount := flag.Int("s", 0, "ignore the first num_chars characters")

	flag.Parse()

	return Flags{
		CountFlag:      *countFlag,
		DuplicatesFlag: *duplicatesFlag,
		UniqueFlag:     *uniqueFlag,
		IgnoreCase:     *ignoreCase,
		FieldCount:     *fieldCount,
		CharCount:      *charCount,
	}
}
