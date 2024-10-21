package database

import (
	"github.com/WizardHowlhaha/golang-debci/models"
	"gorm.io/gorm"
	"log"
)

func MigratePackageStatus(db *gorm.DB) {
	db.AutoMigrate(&models.Package{})
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("INSERT INTO packages(name) SELECT DISTINCT package FROM jobs").Error; err != nil {
			return err
		}

		if err := tx.Exec(`
			UPDATE jobs
			SET package_id = (
				SELECT packages.id
				FROM packages
				WHERE jobs.package = packages.name
			)`).Error; err != nil {
			return err
		}

		if err := tx.Migrator().DropColumn(&models.Job{}, "Package"); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		log.Fatalf("failed to apply migration: %v", err)
	}
}
