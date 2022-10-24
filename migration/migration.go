package migration

import (
	"fmt"
	"task/config"
	"task/models"
)

// Automatic Migration if Running App
func RunMigration() {
	err := config.DB.AutoMigrate(
		&models.Asset{},
		&models.Wallet{},
		&models.Transaction{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
