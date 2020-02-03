package presenter

import (
	"fmt"
	s "strings"

	"github.com/Foo-x/sudoku/internal/pkg/structure"
	"github.com/Foo-x/sudoku/pkg/collection"
)

func Display(values map[string]string) {
	var lengths []int
	for _, square := range structure.Squares {
		lengths = append(lengths, len(values[square]))
	}
	width := 1 + collection.MaxI(lengths)

	element := s.Repeat("-", width*3)
	line := s.Join([]string{element, element, element}, "+")
	for _, r := range structure.Raws {
		for _, c := range structure.Columns {
			var sep string
			if c == '3' || c == '6' {
				sep = "|"
			} else {
				sep = ""
			}
			fmt.Print(center(values[string([]rune{r, c})], width) + sep)
		}
		fmt.Println()
		if r == 'C' || r == 'F' {
			fmt.Println(line)
		}
	}
}

func center(str string, length int) string {
	l := len(str)
	if l >= length {
		return str
	}

	left := (length - l) / 2
	right := (length - l) - left

	return s.Repeat(" ", left) + str + s.Repeat(" ", right)
}
