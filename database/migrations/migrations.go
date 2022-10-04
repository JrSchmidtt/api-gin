package migrations

import (
	"github.com/JrSchmidtt/api-gin/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB){
	db.AutoMigrate(models.Book{})
}