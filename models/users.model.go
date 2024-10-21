package models

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"size:256;not null;unique"`
	Admin    bool   `gorm:"default:false"`
	Key      []Key  `gorm:"constraint:OnDelete:CASCADE;foreignKey:UserID"`
}

func MigrateUsers(db *gorm.DB) {
	db.AutoMigrate(&User{})
}

func CreateUserData(db *gorm.DB) {

	db.Migrator().AddColumn(&Key{}, "UserID")

	db.Migrator().CreateConstraint(&Key{}, "User")

	db.Exec(`INSERT INTO users (username)
	SELECT DISTINCT user FROM keys`)

	db.Exec(`UPDATE keys SET user_id = (
        SELECT id FROM users WHERE users.username = keys.user
    )`)

	db.Migrator().DropColumn(&Key{}, "User")
}

func (user *User) AfterDelete(db *gorm.DB) (err error) {
	db.Clauses(clause.Returning{}).Where("ID = ?", user.ID).Delete(&Key{})
	return
}
