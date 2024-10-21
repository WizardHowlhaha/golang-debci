package main

import (
	"fmt"
	"log"
	"os"

	"github.com/WizardHowlhaha/golang-debci/initializers"
	"github.com/WizardHowlhaha/golang-debci/models"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: apicli <command> [args...]")
		fmt.Println("Commands:")
		fmt.Println("  setkey <username> [keyname]  - Set an API key for the given user")
		fmt.Println("  auth <key>                   - Authenticate the given key")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "setkey":
		// Ensure at least 2 arguments are provided for the setkey command
		if len(os.Args) < 3 {
			fmt.Println("Usage: debci-api setkey <username> [keyname]")
			os.Exit(1)
		}
		username := os.Args[2]
		keyname := "default"
		if len(os.Args) > 3 {
			keyname = os.Args[3]
		}
		setKey(username, keyname)

	case "auth":
		if len(os.Args) < 3 {
			fmt.Println("Usage: debci-api auth <key>")
			os.Exit(1)
		}
		key := os.Args[2]
		authKey(key)

	default:
		fmt.Println("Unknown command:", command)
		fmt.Println("Usage: debci-api <command> [args...]")
		fmt.Println("Commands:")
		fmt.Println("  setkey <username> [keyname]  - Set an API key for the given user")
		fmt.Println("  auth <key>                   - Authenticate the given key")
		os.Exit(1)
	}
}

func setKey(username, keyname string) {
	var user models.User
	db := initializers.DB
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

func authKey(keyStr string) {
	// Authenticate the key
	user, err := models.AuthenticateKey(db, keyStr)
	if err != nil {
		fmt.Fprintln(os.Stderr, "E: invalid key")
		os.Exit(1)
	}

	fmt.Printf("I: Valid key for user `%s`\n", user.Username)
}
