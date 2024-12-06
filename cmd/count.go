package count 

import (
	"fmt"
	"os"
	"unicode/utf8"
	"strings"
)

func CountBytes(p string) (int, error) {
	t, err := readFile(p)
	if err != nil {
		return -1, err
	}

	return len(t), nil
}

func CountLines(p string) (int, error) {
	t, err := readFile(p)
	if err != nil {
		return -1, err
	}

	// -1 at the end to remove the empty string generated from final split.
	return len(strings.Split(t, "\n")) - 1, nil
}

func CountWords(p string) (int, error) {
	t, err := readFile(p)
	if err != nil {
		return -1, err
	}

    return len(strings.Fields(string(t))), nil
}

func CountChars(p string) (int, error) {
	t, err := readFile(p)
	if err != nil {
		return -1, err
	}

	return utf8.RuneCountInString(string(t)), nil
}

func readFile(p string) (string, error) {
	content, err := os.ReadFile(p)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't read file: %s\n", p)
		return "", err
	}

	return string(content), nil
}
