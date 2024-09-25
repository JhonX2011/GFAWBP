package entities

import (
	"time"

	"github.com/JhonX2011/GFAWBP/pkg/domain/models/internal_structs/entities"
)

type Events struct {
	ID           uint      `gorm:"column:id;primaryKey"`
	OriginNode   string    `gorm:"column:origin_node;not null"`
	PartialityID string    `gorm:"column:partiality_id;not null"`
	MovableID    string    `gorm:"column:movable_id;not null"`
	Type         string    `gorm:"column:type;not null"`
	Rehydration  bool      `gorm:"column:rehydration;not null"`
	ArrivedLate  bool      `gorm:"column:arrived_late;not null"`
	EventID      string    `gorm:"column:event_id;not null"`
	EventData    []byte    `gorm:"column:event_data;not null"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (e *Events) ToDomain() *entities.Events {
	return &entities.Events{
		ID:           e.ID,
		OriginNode:   e.OriginNode,
		PartialityID: e.PartialityID,
		MovableID:    e.MovableID,
		Type:         entities.Type(e.Type),
		Rehydration:  e.Rehydration,
		ArrivedLate:  e.ArrivedLate,
		EventID:      e.EventID,
		EventData:    e.EventData,
		CreatedAt:    e.CreatedAt,
		UpdatedAt:    e.UpdatedAt,
	}
}

func (e *Events) FromDomain(entity *entities.Events) {
	e.ID = entity.ID
	e.OriginNode = entity.OriginNode
	e.PartialityID = entity.PartialityID
	e.MovableID = entity.MovableID
	e.Type = string(entity.Type)
	e.Rehydration = entity.Rehydration
	e.ArrivedLate = entity.ArrivedLate
	e.EventID = entity.EventID
	e.EventData = entity.EventData
	e.CreatedAt = entity.CreatedAt
	e.UpdatedAt = entity.UpdatedAt
}
