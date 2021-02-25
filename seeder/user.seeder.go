package seeder

import (
	"helloworld/models"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func SeedUser() ([]models.User, []models.Role) {
	pass, err := bcrypt.GenerateFromPassword([]byte("1234"), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	var role = []models.Role{
		models.Role{
			ID:   1,
			Role: "Petugas",
		},
		models.Role{
			ID:   2,
			Role: "Member",
		},
	}

	var users = []models.User{
		models.User{
			Username: "sayidin",
			Password: string(pass),
			FullName: "Afif Musyayyidin",
			RoleID:   1,
		},
		models.User{
			Username: "Fifa",
			Password: string(pass),
			FullName: "Fifa Musyayyidin",
			RoleID:   2,
		},
	}
	return users, role
}
