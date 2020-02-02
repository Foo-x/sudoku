package strategy

import (
	s "strings"

	"github.com/Foo-x/sudoku/internal/pkg/structure"
)

func Assign(values map[string]string, square string, column string) (map[string]string, bool) {
	others := s.ReplaceAll(values[square], column, "")
	if all(func(v string) bool {
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

		if !all(func(v string) bool {
			_, ok := eliminate(values, v, values[square])
			return ok
		}, keys) {
			return nil, false
		}
	}

	for _, u := range structure.UnitMap[square] {
		dplaces := filter(func(v string) bool { return s.Contains(values[v], column) }, u)
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

func all(f func(v string) bool, slice []string) bool {
	for _, x := range slice {
		if !f(x) {
			return false
		}
	}
	return true
}

func filter(f func(v string) bool, s []string) []string {
	var filtered []string
	for _, x := range s {
		if f(x) {
			filtered = append(filtered, x)
		}
	}
	return filtered
}

func contains(slice []string, s string) bool {
	for _, x := range slice {
		if x == s {
			return true
		}
	}
	return false
}
