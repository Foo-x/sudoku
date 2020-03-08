package solver

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Foo-x/sudoku/internal/app/solver/presenter"
	"github.com/Foo-x/sudoku/internal/pkg/io"
	"github.com/Foo-x/sudoku/internal/pkg/parser"
)

const input = "../../../assets/input.txt"

func Solve() {
	if len(os.Args) < 2 {
		fmt.Println("No arguments. See README.")
		return
	}

	input := os.Args[1]
	wd, _ := os.Getwd()
	lines, err := io.ReadFile(filepath.Join(wd, input))
	if err != nil {
		fmt.Println("Invalid file path.")
		return
	}
	if len(lines) == 0 {
		fmt.Println("No lines in the file.")
		return
	}

	values, _ := parser.Parse(lines[0])

	presenter.Display(values)
}
