package io

import (
	"bufio"
	"os"
)

func ReadFile(path string) ([]string, error) {
	data, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer data.Close()

	sc := bufio.NewScanner(data)

	var lines []string
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	if err := sc.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}
