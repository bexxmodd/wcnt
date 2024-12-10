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
				t.Errorf("expected %d words, got %d", tc.expectedBytes, count)
			}
		})
	}
}
