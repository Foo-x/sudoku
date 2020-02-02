package parser

import (
	s "strings"

	"github.com/Foo-x/sudoku/internal/pkg/strategy"
	"github.com/Foo-x/sudoku/internal/pkg/structure"
	"github.com/Foo-x/sudoku/pkg/collection"
)

func Parse(grid string) (map[string]string, bool) {
	values := map[string]string{}
	for _, square := range structure.Squares {
		values[square] = structure.Columns
	}

	for square, c := range GridValues(grid) {
		if s.Contains(structure.Columns, c) {
			_, ok := strategy.Assign(values, square, c)
			if !ok {
				return nil, false
			}
		}
	}
	return values, true
}

func GridValues(grid string) map[string]string {
	chars := s.Split(grid, "")
	chars = collection.FilterS(func(v string) bool { return s.Contains(structure.Columns, v) || s.Contains("0.", v) }, chars)

	if len(chars) != 81 {
		panic("Invalid characters in input.")
	}

	values := map[string]string{}
	for i, s := range structure.Squares {
		values[s] = chars[i]
	}
	return values
}
