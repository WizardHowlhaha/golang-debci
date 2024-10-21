package main

import (
	"github.com/WizardHowlhaha/golang-debci/initializers"
	"github.com/WizardHowlhaha/golang-debci/models"
)

func main() {

	// 迁移 keys 表
	models.MigrateKeys(initializers.DB)

	// 迁移 users 表
	models.MigrateUsers(initializers.DB)
}
