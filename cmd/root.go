package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "advent2021",
	Short: "Advent of code 2021 in Go",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
