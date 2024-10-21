/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package api

import (
	"fmt"
	"log"
	"os"

	"github.com/WizardHowlhaha/golang-debci/initializers"
	"github.com/WizardHowlhaha/golang-debci/models"
	"github.com/spf13/cobra"
)

// apiCmd represents the api command
var deleteCmd = &cobra.Command{
	Use:  "delete [username]",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]
		delete(username)
	},
}

func delete(username string) {
	var user models.User
	db := initializers.DB
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		fmt.Printf("User: `%s` not exist! \n", username)
		os.Exit(0)
	}
	if err := db.Delete(&user).Error; err != nil {
		log.Fatal("Error delete user:", username)
	}
}
