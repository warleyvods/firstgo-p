package migrations

import (
	"firstgo-p/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	err := db.AutoMigrate(models.Book{})
	if err != nil {
		return
	}
}
