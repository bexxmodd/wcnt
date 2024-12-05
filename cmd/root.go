package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "os"
)

var rootCmd = &cobra.Command{
    Use:   "wcnt",
    Short: "wORD cOUnt",
    Run: func(cmd *cobra.Command, args []string) {
    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintf(os.Stderr, "An error when trying to execute wct %s\n", err)
    } else {
        fmt.Println("I AM COBRA!")
    }
}