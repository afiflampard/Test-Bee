package migrations

import (
	"helloworld/models"
	"helloworld/seeder"
	"log"

	"gorm.io/gorm"
)

var buku = seeder.SeedBuku()

func Buku(db *gorm.DB) error {
	dbBuku := db.Debug().Migrator().DropTable(&models.Buku{})
	if dbBuku != nil {
		log.Fatal("Cannot Drop Table")
	}
	dbBuku = db.Debug().AutoMigrate(&models.Buku{})
	if dbBuku != nil {
		log.Fatal("Cannot migrate Table")
	}

	// for _, buku := range buku {
	// 	err := db.Debug().Create(&buku).Error
	// 	if err != nil {
	// 		log.Fatalf("Failed to Create Buku")
	// 	}
	// }
	return nil
}
