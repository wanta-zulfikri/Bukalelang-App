package config

import (
	// "BukaLelang/app/features/bids/repository"
	// "BukaLelang/app/features/historys/repository"
	// "BukaLelang/app/features/lelangs/repository"
	// "BukaLelang/app/features/transactions/repository"
	users"BukaLelang/app/features/users/repository"

	"gorm.io/gorm"
)
func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
        &users.User{},
		// &lelangs{},
		// &bids{},
		// &historys{},
		// &transactions{}, 
	)
	return err
}