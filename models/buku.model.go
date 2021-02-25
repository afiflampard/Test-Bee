package models

import "gorm.io/gorm"

type Buku struct {
	gorm.Model
	OrderDetails []OrderDetail `gorm:"foreignKey:IDBuku"`
	KodeBuku     string        `gorm:"column:kode_buku; type:varchar(255); not null;unique" json:"kode_buku"`
	JudulBuku    string        `gorm:"column:judul_buku; type:varchar(255); not null" json:"judul_buku"`
	PenulisBuku  string        `gorm:"column:penulis_buku; type:varchar(255); not null" json:"penulis_buku"`
	Stok         uint          `gorm:"column:stok" json:"stok"`
	MaxStock     uint          `gorm:"column:maxStok" json:"max_stock"`
}
