package migrations

import (
	"gorm.io/gorm"
)

//Migrations Table
func Migrations(db *gorm.DB) {
	User(db)
	Buku(db)
	Peminjaman(db)
	History(db)
}
