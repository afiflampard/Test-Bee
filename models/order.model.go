package models

import (
	"time"

	"gorm.io/gorm"
)

// Order ...
type Order struct {
	gorm.Model
	TanggalPeminjaman time.Time  `gorm:"column:tanggal_peminjaman" json:"tanggal_peminjaman"`
	TanggalKembali    time.Time  `gorm:"column:tanggal_kembali" json:"tanggal_kembali"`
	IDPetugas         uint       `gorm:"column:petugas_id" json:"PetugasId"`
	Petugas           User       `gorm:"foreignKey:IDPetugas"`
	IDUser            uint       `gorm:"column:user_id" json:"userId"`
	User              User       `gorm:"foreignKey:IDUser"`
	NoState           uint       `gorm:"column:stateNo"`
	OrderState        OrderState `gorm:"foreignKey:NoState"`
}

// OrderDetail ...
type OrderDetail struct {
	gorm.Model
	IDOrder uint `gorm:"column:order_id"`
	IDBuku  uint `gorm:"column:buku_id"`
	//Bukus []Buku `gorm:"foreignKey:IDBuku"`
	Order Order `gorm:"foreignKey:IDOrder"`
}

// OrderState ...
type OrderState struct {
	ID   uint   `gorm:"column:id; primary_key; AUTO_INCREMENT" json:"id"`
	No   uint   `gorm:"column:state_no"`
	Name string `gorm:"column:state_name"`
}

type History struct {
	gorm.Model
	IDOrder    uint       `gorm:"column:orderId" json:"order_id"`
	Order      Order      `gorm:"foreignKey:IDOrder"`
	IDBuku     uint       `gorm:"column:bukuId" json:"buku_id"`
	Buku       Buku       `gorm:"foreignKey:IDBuku"`
	NoState    uint       `gorm:"column:stateNo" json:"no_state"`
	OrderState OrderState `gorm:"foreignKey:NoState"`
}
