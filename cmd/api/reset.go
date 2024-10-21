/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package api

import (
	"fmt"
	"log"
	"os"

	"github.com/WizardHowlhaha/golang-debci/database"
	"github.com/WizardHowlhaha/golang-debci/models"
	"github.com/spf13/cobra"
)

// apiCmd represents the api command
var resetCmd = &cobra.Command{
	Use:  "reset [username]",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]
		reset(username, "default")
	},
}

func reset(username, keyname string) {
	var user models.User
	db := database.DB
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		fmt.Printf("User: `%s` not exist! \n", username)
		os.Exit(0)
	}

	err := db.Where("user_id = ?", user.ID).Delete(&models.Key{}).Error
	if err != nil {
		log.Fatal("Error delete old keys for user:", username)
	}

	key := models.Key{UserID: user.ID, Username: user.Username}
	if err := db.Create(&key).Error; err != nil {
		log.Fatal("Error creating key:", keyname)
	}

	fmt.Println(key.Key)
}
