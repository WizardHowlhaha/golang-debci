package models

type Package struct {
	ID            uint            `gorm:"primaryKey"`
	Name          string          `gorm:"size:128;unique;not null"`
	Job           []Job           `gorm:"constraint:OnDelete:CASCADE;foreignKey:PackageID"`
	PackageStatus []PackageStatus `gorm:"constraint:OnDelete:CASCADE;foreignKey:PackageID"`
}
