package structure

import (
	s "strings"

	"github.com/Foo-x/sudoku/pkg/collection"
)

const Raws = "ABCDEFGHI"
const Columns = "123456789"
const blockSize = 3

var Squares []string
var Units [][]string
var UnitMap map[string][][]string
var Peers map[string]map[string]struct{}

func init() {
	Squares = cross(s.Split(Raws, ""), s.Split(Columns, ""))

	for _, r := range Raws {
		unit := collection.FilterS(func(v string) bool { return s.HasPrefix(v, string([]rune{r})) }, Squares)
		Units = append(Units, unit)
	}
	for _, c := range Columns {
		unit := collection.FilterS(func(v string) bool { return s.HasSuffix(v, string([]rune{c})) }, Squares)
		Units = append(Units, unit)
	}
	for i := 0; i < len(Raws); i += blockSize {
		for j := 0; j < len(Columns); j += blockSize {
			Units = append(Units, cross(s.Split(Raws, "")[i:i+blockSize], s.Split(Columns, "")[j:j+blockSize]))
		}
	}

	UnitMap = map[string][][]string{}
	for _, s := range Squares {
		for _, unit := range Units {
			for _, label := range unit {
				if label == s {
					UnitMap[label] = append(UnitMap[label], unit)
				}
			}
		}
	}

	Peers = map[string]map[string]struct{}{}
	for _, s := range Squares {
		peer := map[string]struct{}{}
		for _, units4Key := range UnitMap[s] {
			for _, label := range units4Key {
				peer[label] = struct{}{}
			}
		}
		delete(peer, s)
		Peers[s] = peer
	}
}

func cross(sliceA, sliceB []string) []string {
	var crossed []string
	for _, a := range sliceA {
		for _, b := range sliceB {
			crossed = append(crossed, a+b)
		}
	}
	return crossed
}
