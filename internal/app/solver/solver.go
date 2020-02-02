package solver

import (
	"fmt"

	"github.com/Foo-x/sudoku/internal/pkg/parser"
)

const input = "../../../assets/input.txt"

func Solve() {
	grid := "003020600900305001001806400008102900700000008006708200002609500800203009005010300"
	values, _ := parser.Parse(grid)

	fmt.Println(values)
}
