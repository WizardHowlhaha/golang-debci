package models

import (
	"time"

	"gorm.io/gorm"
)

type Job struct {
	RunID       uint   `gorm:"primary_key"`
	Suite       string `gorm:"size:100"`
	Arch        string `gorm:"size:100"`
	Package     string `gorm:"size:100"`
	Version     string `gorm:"size:100"`
	Trigger     string
	Status      string `gorm:"size:25"`
	Requestor   string `gorm:"size:256;index"`
	PinPackages string `gorm:"type:text"`
	Worker      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (k *Job) Migrate(db *gorm.DB) {
	db.AutoMigrate(&Job{})
}
