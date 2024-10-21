package api

import (
	"fmt"
	"github.com/WizardHowlhaha/golang-debci/database"
	"github.com/WizardHowlhaha/golang-debci/lib"
	"github.com/WizardHowlhaha/golang-debci/models"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
	"os"
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

func AuthenticateKey(db *gorm.DB, encryptedKey string) (*models.User, error) {
	var key models.Key
	err := db.Where("encrypted_key = ?", encryptedKey).First(&key).Error
	if err != nil {
		return nil, err
	}

	var user models.User
	err = db.First(&user, key.UserID).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func authKey(keyStr string) {
	// Authenticate the key
	db := database.DB
	encryptedKey := lib.Encrypt(keyStr)

	var key models.Key
	err := db.Where("encrypted_key = ?", encryptedKey).First(&key).Error
	if err != nil {
		fmt.Fprintln(os.Stderr, "E: invalid key - no found encrypted record for", keyStr)
		os.Exit(1)
	}

	var user models.User
	err = db.First(&user, key.UserID).Error
	if err != nil {
		fmt.Fprintln(os.Stderr, "E: invalid key - no found users for", keyStr)
		os.Exit(2)
	}

	fmt.Printf("I: Valid key for user `%s`\n", user.Username)
}
