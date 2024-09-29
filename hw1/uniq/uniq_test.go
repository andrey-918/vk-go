package main

import (
	"github.com/stretchr/testify/require"
	"testing"
	"uniq/utils"
)

func TestProcessFile(t *testing.T) {
	tests := []struct {
		name           string
		input          []string
		expected       []string
		countFlag      bool
		duplicatesFlag bool
		uniqueFlag     bool
		ignoreCase     bool
		fieldCount     int
		charCount      int
	}{
		{
			name:           "Basic unique lines",
			input:          []string{"I love music.", "I love music.", "I love music.", "I love music of Kartik.", "I love music of Kartik.", "Thanks.", "I love music of Kartik.", "I love music of Kartik."},
			expected:       []string{"I love music.", "I love music of Kartik.", "Thanks.", "I love music of Kartik."},
			countFlag:      false,
			duplicatesFlag: false,
			uniqueFlag:     false,
			ignoreCase:     false,
			fieldCount:     0,
			charCount:      0,
		},
		{
			name:           "Count occurrences -c",
			input:          []string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks.", "I love music of Kartik.", "I love music of Kartik."},
			expected:       []string{"3 I love music.", "1 ", "2 I love music of Kartik.", "1 Thanks.", "2 I love music of Kartik."},
			countFlag:      true,
			duplicatesFlag: false,
			uniqueFlag:     false,
			ignoreCase:     false,
			fieldCount:     0,
			charCount:      0,
		},
		{
			name:           "Only print duplicate lines -d",
			input:          []string{"I love music.", "I love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks.", "I love music of Kartik."},
			expected:       []string{"I love music.", "I love music of Kartik."},
			countFlag:      false,
			duplicatesFlag: true,
			uniqueFlag:     false,
			ignoreCase:     false,
			fieldCount:     0,
			charCount:      0,
		},
		{
			name:           "Only print unique lines -u",
			input:          []string{"I love music.", "I love music.", "", "I love music of Kartik.", "Thanks.", "I love music of Kartik."},
			expected:       []string{"", "I love music of Kartik.", "Thanks.", "I love music of Kartik."},
			countFlag:      false,
			duplicatesFlag: false,
			uniqueFlag:     true,
			ignoreCase:     false,
			fieldCount:     0,
			charCount:      0,
		},
		{
			name:           "test with -f 1",
			input:          []string{"We love music.", "I love music.", "They love music.", "", "I love music of Kartik.", "We love music of Kartik.", "Thanks."},
			expected:       []string{"We love music.", "", "I love music of Kartik.", "Thanks."},
			countFlag:      false,
			duplicatesFlag: false,
			uniqueFlag:     false,
			ignoreCase:     false,
			fieldCount:     1,
			charCount:      0,
		},
		{
			name:           "test with -s 1",
			input:          []string{"I love music.", "A love music.", "C love music.", "", "I love music of Kartik.", "We love music of Kartik.", "Thanks."},
			expected:       []string{"I love music.", "", "I love music of Kartik.", "We love music of Kartik.", "Thanks."},
			countFlag:      false,
			duplicatesFlag: false,
			uniqueFlag:     false,
			ignoreCase:     false,
			fieldCount:     0,
			charCount:      1,
		},
		{
			name:           "test with [-f 1 -s 1]",
			input:          []string{"I 1ove music.", "A 2ove music.", "C 3ove music.", "", "I 4ove music of Kartik.", "We 5ove music of Kartik.", "Thanks."},
			expected:       []string{"I 1ove music.", "", "I 4ove music of Kartik.", "Thanks."},
			countFlag:      false,
			duplicatesFlag: false,
			uniqueFlag:     false,
			ignoreCase:     false,
			fieldCount:     1,
			charCount:      1,
		},
		{
			name:           "test with [-f 1 -s 1 -c]",
			input:          []string{"I 1ove music.", "A 2ove music.", "C 3ove music.", "", "I 4ove music of Kartik.", "We 5ove music of Kartik.", "Thanks."},
			expected:       []string{"3 I 1ove music.", "1 ", "2 I 4ove music of Kartik.", "1 Thanks."},
			countFlag:      true,
			duplicatesFlag: false,
			uniqueFlag:     false,
			ignoreCase:     false,
			fieldCount:     1,
			charCount:      1,
		},
		{
			name:           "Ignore case -i",
			input:          []string{"I LOVE MUSIC.", "I love music.", "I LoVe MuSiC.", "", "I love MuSIC of Kartik.", "I love music of kartik.", "Thanks.", "I love music of kartik.", "I love MuSIC of Kartik."},
			expected:       []string{"I LOVE MUSIC.", "", "I love MuSIC of Kartik.", "Thanks.", "I love music of kartik."},
			countFlag:      false,
			duplicatesFlag: false,
			uniqueFlag:     false,
			ignoreCase:     true,
			fieldCount:     0,
			charCount:      0,
		},
		{
			name:           "Options -c, -d, and -u are mutually exclusive.",
			input:          []string{"I LOVE MUSIC.", "I love music.", "I LoVe MuSiC.", "", "I love MuSIC of Kartik.", "I love music of kartik.", "Thanks.", "I love music of kartik.", "I love MuSIC of Kartik."},
			expected:       []string{"Error: Options -c, -d, and -u are mutually exclusive."},
			countFlag:      true,
			duplicatesFlag: true,
			uniqueFlag:     true,
			ignoreCase:     true,
			fieldCount:     0,
			charCount:      0,
		},
		{
			name:           "Options -c, -d, and -u are mutually exclusive.",
			input:          []string{"I LOVE MUSIC.", "I love music.", "I LoVe MuSiC.", "", "I love MuSIC of Kartik.", "I love music of kartik.", "Thanks.", "I love music of kartik.", "I love MuSIC of Kartik."},
			expected:       []string{"Error: Options -c, -d, and -u are mutually exclusive."},
			countFlag:      true,
			duplicatesFlag: false,
			uniqueFlag:     true,
			ignoreCase:     true,
			fieldCount:     0,
			charCount:      0,
		},
		{
			name:           "Options -c, -d, and -u are mutually exclusive.",
			input:          []string{"I LOVE MUSIC.", "I love music.", "I LoVe MuSiC.", "", "I love MuSIC of Kartik.", "I love music of kartik.", "Thanks.", "I love music of kartik.", "I love MuSIC of Kartik."},
			expected:       []string{"Error: Options -c, -d, and -u are mutually exclusive."},
			countFlag:      true,
			duplicatesFlag: true,
			uniqueFlag:     false,
			ignoreCase:     true,
			fieldCount:     0,
			charCount:      0,
		},
		{
			name:           "Empty",
			input:          []string{""},
			expected:       []string{""},
			countFlag:      false,
			duplicatesFlag: false,
			uniqueFlag:     false,
			ignoreCase:     false,
			fieldCount:     0,
			charCount:      0,
		},

		{
			name:           "1 word",
			input:          []string{"90"},
			expected:       []string{"90"},
			countFlag:      false,
			duplicatesFlag: false,
			uniqueFlag:     false,
			ignoreCase:     false,
			fieldCount:     0,
			charCount:      0,
		},
		{
			name:           "1 word and f 10 s 10",
			input:          []string{"90"},
			expected:       []string{"90"},
			countFlag:      false,
			duplicatesFlag: false,
			uniqueFlag:     false,
			ignoreCase:     false,
			fieldCount:     10,
			charCount:      10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flags := utils.Flags{
				CountFlag:      tt.countFlag,
				DuplicatesFlag: tt.duplicatesFlag,
				UniqueFlag:     tt.uniqueFlag,
				IgnoreCase:     tt.ignoreCase,
				FieldCount:     tt.fieldCount,
				CharCount:      tt.charCount,
			}

			output := utils.ProcessFile(tt.input, flags)

			require.Equal(t, tt.expected, output)
		})
	}
}
