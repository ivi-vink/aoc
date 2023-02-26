package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var templateStr string = `
`

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate a new boilerplate entrypoint for a day of Advent of Code",
	RunE: func(cmd *cobra.Command, args []string) error {
		year, day := "", ""
		if len(args) == 1 {
			parts := strings.SplitN(args[0], "/", 2)
			if len(parts) == 2 {
				year = parts[0]
				day = parts[1]
			}
		} else {
			return fmt.Errorf("Expected one argument in format %q", "<year>/<day>")
		}

		if year == "" || day == "" {
			return fmt.Errorf("Expected one argument in format %q", "<year>/<day>")
		}

		_, err := strconv.Atoi(year)
		if err != nil {
			return fmt.Errorf("Couldn't parse year integer: %q", year)
		}

		dayInt, err := strconv.Atoi(day)
		if err != nil {
			return fmt.Errorf("Couldn't parse day integer: %q", day)
		}

		if dayInt < 1 || dayInt > 25 {
			return fmt.Errorf("Day must be between 1 and 25, got %d", dayInt)
		}

		dayStr := fmt.Sprintf("%02d", dayInt)
		fmt.Println(dayStr)
		return nil
	},
}
