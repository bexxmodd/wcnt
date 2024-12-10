package root

import (
	"fmt"
	"os"

	"github.com/bexxmodd/wcnt/blob/master/cmd/count"
	"github.com/spf13/cobra"
)

type CounterFunc func(string) (int, error)

var (
	countBytes bool
	countWords bool
	countLines bool
	countChars bool

	rootCmd = &cobra.Command{
		Use:   "wcnt",
		Short: "Count words",
		Run: func(cmd *cobra.Command, args []string) {
			var counter CounterFunc
			switch {
			case countBytes:
				counter = count.CountBytes
			case countLines:
				counter = count.CountLines
			case countWords:
				counter = count.CountWords
			case countChars:
				counter = count.CountChars
			default:
				printAll(args)
				return
			}
			printCount(args, counter)
		},
	}
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&countBytes, "bytes", "c", false, "print the byte count.")
	rootCmd.PersistentFlags().BoolVarP(&countLines, "lines", "l", false, "print the lines count.")
	rootCmd.PersistentFlags().BoolVarP(&countWords, "words", "w", false, "print the words count.")
	rootCmd.PersistentFlags().BoolVarP(&countChars, "chars", "m", false, "print the characters count.")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "An error when trying to execute wcnt %s\n", err)
	}
}

func printCount(paths []string, counter CounterFunc) {
	for _, p := range paths {
		cnt, err := counter(p)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			return
		}
		fmt.Printf("\t%d %s\n", cnt, p)
	}
}

func printAll(paths []string) {
	for _, p := range paths {
		l, e2 := count.CountLines(p)
		w, e3 := count.CountWords(p)
		b, e1 := count.CountBytes(p)
		if e1 != nil || e2 != nil || e3 != nil {
			fmt.Fprint(os.Stderr, "Error when trying to count\n")
		}
		fmt.Printf("\t%d %d %d %s\n", l, w, b, p)
	}
}
