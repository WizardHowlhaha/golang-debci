package models

import (
	"time"

	"github.com/WizardHowlhaha/golang-debci/lib"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Key struct {
	ID           uint `gorm:"primaryKey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Username     string `gorm:"size:256;not null;"`
	Key          string `gorm:"-"`
	EncryptedKey string `gorm:"size:40;not null;index"`
	UserID       uint   `gorm:"not null"`
}

func (key *Key) BeforeCreate(tx *gorm.DB) (err error) {
	uid := uuid.New()
	key.Key = uid.String()
	key.EncryptedKey = lib.Encrypt(key.Key)
	return
}
