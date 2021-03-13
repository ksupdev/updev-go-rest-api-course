package database

import (
	"github.com/jinzhu/gorm"
	"github.com/ksupdev/updev-go-rest-api-course/internal/comment"
)

// MigrateDB - migration our Database and creates our comment table
func MigrateDB(db *gorm.DB) error {
	if result := db.AutoMigrate(&comment.Comment{}); result.Error != nil {
		return result.Error
	}
	return nil
}
