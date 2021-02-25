package migrations

import (
	"helloworld/models"
	"helloworld/seeder"
	"log"

	"gorm.io/gorm"
)

var users, role = seeder.SeedUser()

func User(db *gorm.DB) error {
	dbRole := db.Debug().Migrator().DropTable(&models.Role{})

	if dbRole != nil {
		log.Fatal("Cannot Drop Table")
	}
	dbRole = db.Debug().AutoMigrate(&models.Role{})
	if dbRole != nil {
		log.Fatal("Cannot migrate Table")
	}
	dbUser := db.Debug().Migrator().DropTable(&models.User{})
	if dbUser != nil {
		log.Fatal("Cannot Drop table")
	}
	dbUser = db.Debug().AutoMigrate(&models.User{})
	if dbUser != nil {
		log.Fatal("Cannot migrate Table")
	}

	for _, role := range role {
		err := db.Debug().Create(&role).Error
		if err != nil {
			log.Fatalf("Failed to Create Role")
		}
	}

	for _, user := range users {
		err := db.Debug().Create(&user).Error
		if err != nil {
			log.Fatalf("failde to create user")
		}
	}

	return nil
}
