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

func (user *User) AfterDelete(db *gorm.DB) (err error) {
	db.Clauses(clause.Returning{}).Where("ID = ?", user.ID).Delete(&Key{})
	return
}
