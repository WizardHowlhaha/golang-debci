/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"os"

	"github.com/WizardHowlhaha/golang-debci/cmd/api"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var debciCmd = &cobra.Command{
	Use:       "debci [COMMAND] [ARGS]",
	ValidArgs: []string{"api"},
	Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("debci called")
	},
}

func init() {
	debciCmd.PersistentFlags().BoolP("help", "h", false, "Print usage")
	debciCmd.PersistentFlags().Lookup("help").Hidden = true
	debciCmd.CompletionOptions.DisableDefaultCmd = true
	debciCmd.AddCommand(api.ApiCmd)
}

func main() {
	err := debciCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
