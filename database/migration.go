package database

import (
	"fmt"
	"project/models"
	"project/pkg/mysql"
)

// Jika aplikasi berjalan maka auto migration akan berjalan
func RunMigration() {
	// koneksi database akan melakukan auto migrasi struct user ke database
	err := mysql.DB.AutoMigrate(
		&models.User{},
		&models.Trip{},
		&models.Country{},
		&models.Transaction{},
		// &models.Profile{},
	)
	// jika tidak ada error
	if err != nil {
		fmt.Println(err)
		panic("Migration failed")
	}

	fmt.Println("Migration success")
}
