package migrations

import (
	"helloworld/models"
	"log"

	"gorm.io/gorm"
)

//History is ...
func History(db *gorm.DB) error {
	dbHistory := db.Debug().Migrator().DropTable(&models.History{})
	if dbHistory != nil {
		log.Fatal("Cannot Drop Table")
	}
	dbHistory = db.Debug().AutoMigrate(&models.History{})
	if dbHistory != nil {
		log.Fatal("Cannot migrate Table")
	}
	return nil
}
