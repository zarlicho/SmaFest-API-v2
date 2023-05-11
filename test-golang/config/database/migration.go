package database

import (
	"fmt"
	"test-golang/config/mongdb"
	"test-golang/models"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.Product{},
		&models.User{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println(("Migration Success"))
}
