package seeder

import (
	"helloworld/models"
)

func SeedOrder() []models.OrderState {
	var orderState = []models.OrderState{
		{
			No:   1,
			Name: "Dipinjam",
		},
		{
			No:   2,
			Name: "Dikembalikan",
		},
	}

	return orderState
}
