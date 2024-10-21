/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package api

import (
	"fmt"
	"log"

	"github.com/WizardHowlhaha/golang-debci/database"
	"github.com/WizardHowlhaha/golang-debci/models"
	"github.com/spf13/cobra"
)

// apiCmd represents the api command
var setkeyCmd = &cobra.Command{
	Use:  "setkey [username] [keyname]",
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]
		keyname := args[1]
		if keyname == "" {
			keyname = "default"
		}
		setKey(username, keyname)
	},
}

func setKey(username, keyname string) {
	var user models.User
	db := database.DB
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		user = models.User{Username: username}
		if err := db.Create(&user).Error; err != nil {
			log.Fatal("Error creating user: ", username)
		}
	}

	key := models.Key{UserID: user.ID, Username: user.Username}
	if err := db.Create(&key).Error; err != nil {
		log.Fatal("Error creating key:", keyname)
	}

	fmt.Println(key.Key)
}
