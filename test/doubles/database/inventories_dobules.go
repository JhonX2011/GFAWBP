package database

import (
	"time"

	"github.com/JhonX2011/GFAWBP/pkg/domain/models/internal_structs/entities"
)

func GetDataInventories() *entities.Inventories {
	return &entities.Inventories{
		ID:              1,
		InventoryID:     "INV-01",
		MovableID:       "MOV-01",
		DestinationNode: "BRRC02",
		Quantity:        20,
		Version:         0,
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
	}
}
