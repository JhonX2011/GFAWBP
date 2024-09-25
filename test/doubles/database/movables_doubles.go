package database

import (
	"time"

	"github.com/JhonX2011/GFAWBP/pkg/domain/models/internal_structs/entities"
)

func GetDataDomainMovables() *entities.Movables {
	return &entities.Movables{
		ID:           "MO-01",
		OriginNode:   "BRRC01",
		PartialityID: "abc123",
		Status:       "CHECKED",
		Version:      0,
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		Inventories:  make([]entities.Inventories, 2),
	}
}
