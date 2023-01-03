package sfolib

import (
	"bufio"
	"os"
)

// read N lined from pointer to file
func ReadLines(f *os.File, n int) []string {
	lines := make([]string, 0, n)
	sc := bufio.NewScanner(f)
	for i := 0; i < n && sc.Scan(); i++ {
		lines = append(lines, sc.Text())
	}

	return lines
}

// read N lines from file
func ReadLinesFromFile(s string, n int) ([]string, error) {
	f, err := os.Open(s)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	lines := make([]string, 0, n)
	sc := bufio.NewScanner(f)
	for i := 0; i < n && sc.Scan(); i++ {
		lines = append(lines, sc.Text())
	}

	return lines, nil
}

func ReadFirstLine(s string, n int) (string, error) {
	f, err := os.Open(s)
	if err != nil {
		return "", err
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	sc.Scan()
	return sc.Text(), nil
}
