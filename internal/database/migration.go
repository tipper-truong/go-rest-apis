package database

import (
	"github.com/jinzhu/gorm"
	"github.com/tipper-truong/go-rest-apis/internal/database/comment"
)

// Import comment and gorm package from our rest api project
// Migrate our database and creates our comment table
func MigrateDB(db *gorm.DB) error {
	if result := db.AutoMigrate(&comment.Comment{}); result.Error != nil {
		return result.Error
	}
	return nil
}
