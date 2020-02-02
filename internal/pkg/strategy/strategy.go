package strategy

import (
	s "strings"

	"github.com/Foo-x/sudoku/internal/pkg/structure"
	"github.com/Foo-x/sudoku/pkg/collection"
)

func Assign(values map[string]string, square string, column string) (map[string]string, bool) {
	others := s.ReplaceAll(values[square], column, "")
	if collection.AllS(func(v string) bool {
		_, ok := eliminate(values, square, v)
		return ok
	}, s.Split(others, "")) {
		return values, true
	}
	return nil, false
}

func eliminate(values map[string]string, square string, column string) (map[string]string, bool) {
	if !s.Contains(values[square], column) {
		return values, true
	}

	values[square] = s.ReplaceAll(values[square], column, "")
	if len(values[square]) == 0 {
		return nil, false
	}
	if len(values[square]) == 1 {
		var keys []string
		for k := range structure.Peers[square] {
			keys = append(keys, k)
		}

		if !collection.AllS(func(v string) bool {
			_, ok := eliminate(values, v, values[square])
			return ok
		}, keys) {
			return nil, false
		}
	}

	for _, u := range structure.UnitMap[square] {
		dplaces := collection.FilterS(func(v string) bool { return s.Contains(values[v], column) }, u)
		if len(dplaces) == 0 {
			return nil, false
		}
		if len(dplaces) == 1 {
			_, ok := Assign(values, dplaces[0], column)
			if !ok {
				return nil, false
			}
		}
	}

	return values, true
}
