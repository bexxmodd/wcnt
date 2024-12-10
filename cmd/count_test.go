package count

import (
	"fmt"
	"testing"

	"github.com/bexxmodd/wcnt/blob/master/cmd/count"
)

func TestCountBytes(t *testing.T) {
	tests := []struct {
		name          string
		content       string
		err           error
		expectedBytes int
	}{
		{
			content:       "hello world",
			err:           nil,
			expectedBytes: 11,
		},
		{
			content:       "",
			err:           fmt.Errorf("mock error"),
			expectedBytes: -1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			count.ReadFileFunc = func(string) (string, error) {
				return tc.content, tc.err
			}
			count, err := count.CountBytes("")

			if tc.err != nil {
				if err == nil {
					t.Error("expected error but got none")
				}
			}

			if count != tc.expectedBytes {
				t.Errorf("expected %d bytes, got %d", tc.expectedBytes, count)
			}
		})
	}
}

func TestCountLines(t *testing.T) {
	tests := []struct {
		name          string
		content       string
		err           error
		expectedLines int
	}{
		{
			content:       "hello world\nbye world\n",
			err:           nil,
			expectedLines: 2,
		},
		{
			content:       "1\n2\n3\n4\n",
			err:           nil,
			expectedLines: 4,
		},
		{
			content:       "",
			err:           fmt.Errorf("mock error"),
			expectedLines: -1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			count.ReadFileFunc = func(string) (string, error) {
				return tc.content, tc.err
			}
			count, err := count.CountLines("")

			if tc.err != nil {
				if err == nil {
					t.Error("expected error but got none")
				}
			}

			if count != tc.expectedLines {
				t.Errorf("expected %d lines, got %d", tc.expectedLines, count)
			}
		})
	}
}

func TestCountWords(t *testing.T) {
	tests := []struct {
		name          string
		content       string
		err           error
		expectedWords int
	}{
		{
			content:       "hello, world",
			err:           nil,
			expectedWords: 2,
		},
		{
			content:       "1\n2\n3\n4\n",
			err:           nil,
			expectedWords: 4,
		},
		{
			content:       "",
			err:           fmt.Errorf("mock error"),
			expectedWords: -1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			count.ReadFileFunc = func(string) (string, error) {
				return tc.content, tc.err
			}
			count, err := count.CountWords("")

			if tc.err != nil {
				if err == nil {
					t.Error("expected error but got none")
				}
			}

			if count != tc.expectedWords {
				t.Errorf("expected %d words, got %d", tc.expectedWords, count)
			}
		})
	}
}