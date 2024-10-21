package models

type PackageStatus struct {
	ID        uint   `gorm:"primaryKey"`
	PackageID uint   `gorm:"index;not null"`
	JobID     uint   `gorm:"index;not null"`
	Arch      string `gorm:"size:255;not null"`
	Suite     string `gorm:"size:255;not null"`
}
