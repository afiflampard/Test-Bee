package migrations

import (
	"helloworld/models"
	"helloworld/seeder"
	"log"

	"gorm.io/gorm"
)

var orderState = seeder.SeedOrder()

func Peminjaman(db *gorm.DB) error {
	dbOrderState := db.Debug().Migrator().DropTable(&models.OrderState{})
	dbPeminjaman := db.Debug().Migrator().DropTable(&models.Order{})
	dbOrderDetail := db.Debug().Migrator().DropTable(&models.OrderDetail{})

	if dbPeminjaman != nil || dbOrderState != nil || dbOrderDetail != nil {
		log.Fatal("Cannot Drop Table")
	}

	dbOrderState = db.Debug().AutoMigrate(&models.OrderState{})
	dbPeminjaman = db.Debug().AutoMigrate(&models.Order{})
	dbOrderDetail = db.Debug().AutoMigrate(&models.OrderDetail{})
	if dbPeminjaman != nil || dbOrderState != nil || dbOrderDetail != nil {
		log.Fatal("Cannot migrate Table OrderState")
	}

	for _, orderState := range orderState {
		err := db.Debug().Create(&orderState).Error
		if err != nil {
			log.Fatalf("Failed to Create orderState")
		}
	}
	// err := db.Debug().Create(&order).Error
	// if err != nil {
	// 	log.Fatalf("Failed to Create Order")
	// }
	// for _, orderDetail := range orderDetail {
	// 	err := db.Debug().Create(&orderDetail).Error
	// 	if err != nil {
	// 		log.Fatalf("Failed to Create orderState")
	// 	}
	// }

	return nil
}
