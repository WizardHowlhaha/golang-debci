/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package api

import (
	"fmt"
	"os"

	"github.com/WizardHowlhaha/golang-debci/initializers"
	"github.com/WizardHowlhaha/golang-debci/models"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var authCmd = &cobra.Command{
	Use: "auth [keystring]",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		keystr := args[0]
		authKey(keystr)
	},
}

func authKey(keyStr string) {
	// Authenticate the key
	db := initializers.DB
	user, err := models.AuthenticateKey(db, keyStr)
	if err != nil {
		fmt.Fprintln(os.Stderr, "E: invalid key")
		os.Exit(1)
	}

	fmt.Printf("I: Valid key for user `%s`\n", user.Username)
}
