package database

import (
	"time"

	"github.com/JhonX2011/GFAWBP/pkg/domain/models/internal_structs/entities"
)

func GetDataEvents() *entities.Events {
	return &entities.Events{
		ID:           0,
		OriginNode:   "BRRC01",
		PartialityID: "74ad5f3c-3015-4a7e-b831-d2eab4ec9439",
		MovableID:    "MOV-01",
		Type:         "",
		Rehydration:  false,
		ArrivedLate:  false,
		EventID:      "ad7e72de-bec6-489c-9dfe-dfb55ffc6e11",
		EventData:    nil,
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
	}
}
