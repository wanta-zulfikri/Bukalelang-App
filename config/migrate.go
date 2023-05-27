package config

import (
	// "BukaLelang/app/features/bids/repository"
	historys "BukaLelang/app/features/historys/repository"
	lelangs "BukaLelang/app/features/lelangs/repository"
	users "BukaLelang/app/features/users/repository"
	// "BukaLelang/app/features/transactions/repository"

	"gorm.io/gorm"
)
func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
        &users.User{},
		&lelangs.Lelang{},
		// &bids{},
		&historys.History{},
		// &transactions{}, 
	)
	return err
}