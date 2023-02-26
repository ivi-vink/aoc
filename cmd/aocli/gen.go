package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate a new key pair",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
	},
}
