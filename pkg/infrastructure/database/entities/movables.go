package entities

import (
	"time"

	"github.com/JhonX2011/GFAWBP/pkg/domain/models/internal_structs/database/entities"
	"gorm.io/plugin/optimisticlock"
)

type Movables struct {
	ID           string                 `gorm:"column:id;primaryKey"`
	OriginNode   string                 `gorm:"column:origin_node;not null"`
	PartialityID string                 `gorm:"column:partiality_id;not null"`
	Status       string                 `gorm:"column:status;not null"`
	CreatedAt    time.Time              `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time              `gorm:"column:updated_at;autoUpdateTime"`
	Version      optimisticlock.Version `gorm:"column:version"`

	Inventories []Inventories `gorm:"foreignKey:MovableID"`
}

func (e *Movables) ToDomain() *entities.Movables {
	var inventories []entities.Inventories
	auxInventories := e.Inventories
	for i := range auxInventories {
		inventories = append(inventories, *auxInventories[i].ToDomain())
	}

	return &entities.Movables{
		ID:           e.ID,
		OriginNode:   e.OriginNode,
		PartialityID: e.PartialityID,
		Status:       entities.Status(e.Status),
		Version:      e.Version.Int64,
		CreatedAt:    e.CreatedAt,
		UpdatedAt:    e.UpdatedAt,
		Inventories:  inventories,
	}
}

func (e *Movables) FromDomain(entity *entities.Movables) {
	e.ID = entity.ID
	e.OriginNode = entity.OriginNode
	e.PartialityID = entity.PartialityID
	e.Status = string(entity.Status)
	e.Version = optimisticlock.Version{Int64: entity.Version}
	e.CreatedAt = entity.CreatedAt
	e.UpdatedAt = entity.UpdatedAt

	for i := range entity.Inventories {
		inventory := entity.Inventories[i]
		var inv Inventories
		inv.FromDomain(&inventory)
		e.Inventories = append(e.Inventories, inv)
	}
}
