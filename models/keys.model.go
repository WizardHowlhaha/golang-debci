package models

import (
	"crypto/sha1"
	"encoding/hex"
	"time"

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
	uuid := uuid.New()
	key.Key = uuid.String()
	key.EncryptedKey = encrypt(key.Key)
	return
}

func encrypt(key string) string {
	hasher := sha1.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func AuthenticateKey(db *gorm.DB, keyString string) (*User, error) {
	var key Key
	encryptedKey := encrypt(keyString)
	err := db.Where("encrypted_key = ?", encryptedKey).First(&key).Error
	if err != nil {
		return nil, err
	}

	var user User
	err = db.First(&user, key.UserID).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func MigrateKeys(db *gorm.DB) {
	db.AutoMigrate(&Key{})
}
