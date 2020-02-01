package solver

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/Foo-x/sudoku/internal/pkg/io"
)

const input = "../../../assets/input.txt"

func Solve() {
	_, entry, _, _ := runtime.Caller(1)
	lines, _ := io.ReadFile(filepath.Join(entry, input))

	fmt.Println(lines)
}
