package main

import (
	"github.com/WizardHowlhaha/golang-debci/database"
	"github.com/WizardHowlhaha/golang-debci/models"
	"github.com/WizardHowlhaha/golang-debci/route"
)

func main() {
	db := database.DB

	db.AutoMigrate(&models.User{}, &models.Key{}, &models.Job{})
}
