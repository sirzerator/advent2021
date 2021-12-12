package cmd

import (
	"github.com/sirzerator/advent2021/days/day1"
	"github.com/sirzerator/advent2021/days/day2"
	"github.com/sirzerator/advent2021/days/day3"
	"github.com/sirzerator/advent2021/days/day4"
	"github.com/sirzerator/advent2021/days/day5"
	"github.com/sirzerator/advent2021/days/day6"
	"github.com/sirzerator/advent2021/days/day7"
	"github.com/sirzerator/advent2021/days/day8"
	"github.com/sirzerator/advent2021/days/day9"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:       "run [day]",
	Short:     "Run a specific days challenge",
	Long:      ``,
	Args:      cobra.ExactValidArgs(1),
	ValidArgs: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"},
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "1":
			day1.Run(Verbose)
		case "2":
			day2.Run(Verbose)
		case "3":
			day3.Run(Verbose)
		case "4":
			day4.Run(Verbose)
		case "5":
			day5.Run(Verbose)
		case "6":
			day6.Run(Verbose)
		case "7":
			day7.Run(Verbose)
		case "8":
			day8.Run(Verbose)
		case "9":
			day9.Run(Verbose)
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
