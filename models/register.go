package models

import "gorm.io/gorm"

// 所有要 migrate 的 model 一律註冊在這裡
func MigrateAll(db *gorm.DB) error {
	return db.AutoMigrate(
		&User{},
		&Card{},
	)
}
