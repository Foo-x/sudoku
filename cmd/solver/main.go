package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Foo-x/sudoku/internal/pkg/io"
)

const input = "../../assets/input.txt"

func main() {
	currentPath, _ := os.Getwd()
	lines, _ := io.ReadFile(filepath.Join(currentPath, input))

	fmt.Println(lines)
}
