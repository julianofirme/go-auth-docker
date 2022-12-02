package migrations

import (
	"github.com/jfirme-sys/go-auth-docker.git/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.User{})
}
