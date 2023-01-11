package sfolib

import (
	"bufio"
	"os"
	"strings"
)

// read N lines from file
func ReadLines(s string, n int) ([]string, error) {
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

// read first line from file
func ReadFirstLine(s string) (string, error) {
	f, err := os.Open(s)
	if err != nil {
		return "", err
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	sc.Scan()
	return sc.Text(), nil
}

// read N line from file
func ReadLine(s string, n int) (string, error) {
	f, err := os.Open(s)
	if err != nil {
		return "", err
	}
	defer f.Close()

	var line string
	sc := bufio.NewScanner(f)
	for i := 0; i < n && sc.Scan(); i++ {
		if i == n-1 {
			line = sc.Text()
		}
	}

	return line, nil
}

// return all lines of file as slice of strings
func LoadFile(s string) ([]string, error) {
	f, err := os.ReadFile(s)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(f), "\n")
	return lines, nil
}

// return true if S file exists
func Exists(s string) bool {
	_, err := os.Stat(s)
	return !os.IsNotExist(err)
}
