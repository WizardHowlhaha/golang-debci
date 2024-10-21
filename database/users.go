package database

import (
	"github.com/WizardHowlhaha/golang-debci/models"
	"gorm.io/gorm"
)

func MigrateUsers(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}

func CreateUserData(db *gorm.DB) {

	db.Migrator().AddColumn(&models.Key{}, "UserID")

	db.Migrator().CreateConstraint(&models.Key{}, "User")

	db.Exec(`INSERT INTO users (username)
	SELECT DISTINCT user FROM keys`)

	db.Exec(`UPDATE keys SET user_id = (
        SELECT id FROM users WHERE users.username = keys.user
    )`)

	db.Migrator().DropColumn(&models.Key{}, "User")
}
