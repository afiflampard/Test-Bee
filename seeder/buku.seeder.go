package seeder

import "helloworld/models"

func SeedBuku() []models.Buku {
	var buku = []models.Buku{
		models.Buku{
			KodeBuku:    "1234",
			JudulBuku:   "Aku adalah",
			PenulisBuku: "Fifa",
			Stok:        2,
		},
		models.Buku{
			KodeBuku:    "345",
			JudulBuku:   "Jiii",
			PenulisBuku: "Fif",
			Stok:        3,
		},
	}
	return buku
}
