package models

import (
	"time"
)

type Job struct {
	RunID         uint   `gorm:"primary_key"`
	Suite         string `gorm:"size:100"`
	Arch          string `gorm:"size:100"`
	Package       string `gorm:"size:100"`
	Version       string `gorm:"size:100"`
	Trigger       string
	Status        string `gorm:"size:25"`
	Requestor     string `gorm:"size:256;index"`
	PinPackages   string `gorm:"type:text"`
	Worker        string
	PackageID     uint            `gorm:"not null"`
	PackageStatus []PackageStatus `gorm:"constraint:OnDelete:CASCADE;foreignKey:JobID"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
