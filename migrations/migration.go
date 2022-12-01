package migrations

import (
	"github.com/jfirme-sys/go-auth-docker.git/models"
	"gorm.io/gorm"
)

func RunMigrastions(db *gorm.DB) {
	db.AutoMigrate(models.User{})
}
