package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "aocli [command]",
	Short: "AOCLI is a CLI for running Advent of Code solvers and generating boilerplate code",
}

func init() {
	rootCmd.AddCommand(genCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
