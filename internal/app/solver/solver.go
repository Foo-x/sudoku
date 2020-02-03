package solver

import (
	"github.com/Foo-x/sudoku/internal/app/solver/presenter"
	"github.com/Foo-x/sudoku/internal/pkg/parser"
)

const input = "../../../assets/input.txt"

func Solve() {
	grid := "003020600900305001001806400008102900700000008006708200002609500800203009005010300"
	values, _ := parser.Parse(grid)

	presenter.Display(values)
}
