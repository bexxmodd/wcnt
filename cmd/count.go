package count 

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"
)

var ReadFileFunc = ReadFile

func CountBytes(p string) (int, error) {
	t, err := ReadFileFunc(p)
	if err != nil {
		return -1, err
	}

	return len(t), nil
}

func CountLines(p string) (int, error) {
	t, err := ReadFileFunc(p)
	if err != nil {
		return -1, err
	}

	// -1 at the end to remove the empty string generated from final split.
	return len(strings.Split(t, "\n")) - 1, nil
}

func CountWords(p string) (int, error) {
	t, err := ReadFileFunc(p)
	if err != nil {
		return -1, err
	}

	return len(strings.Fields(string(t))), nil
}

func CountChars(p string) (int, error) {
	t, err := ReadFileFunc(p)
	if err != nil {
		return -1, err
	}

	return utf8.RuneCountInString(string(t)), nil
}

func ReadFile(p string) (string, error) {
	fp, err := fullPath(p)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't determine full path: %s\n", fp)
		return "", err
	}
	content, err := os.ReadFile(fp)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't read file: %s\n", p)
		return "", err
	}

	return string(content), nil
}

func fullPath(rp string) (string, error) {
	if filepath.IsAbs(rp) {
		return rp, nil
	}

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Can't get working directory.")
		return "", err
	}

	fp := filepath.Join(pwd, rp)
	fp = filepath.Clean(fp)
	return fp, nil
}
