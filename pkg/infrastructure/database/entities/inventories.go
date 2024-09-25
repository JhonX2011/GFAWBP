package entities

import (
	"time"

	"github.com/JhonX2011/GFAWBP/pkg/domain/models/internal_structs/entities"
	"gorm.io/plugin/optimisticlock"
)

type Inventories struct {
	ID              uint                   `gorm:"column:id;primaryKey"`
	InventoryID     string                 `gorm:"column:inventory_id;not null"`
	MovableID       string                 `gorm:"column:movable_id;not null"`
	DestinationNode string                 `gorm:"column:destination_node"`
	Quantity        int                    `gorm:"column:quantity"`
	CreatedAt       time.Time              `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt       time.Time              `gorm:"column:updated_at;autoUpdateTime"`
	Version         optimisticlock.Version `gorm:"column:version"`
}

func (e *Inventories) ToDomain() *entities.Inventories {
	return &entities.Inventories{
		ID:              e.ID,
		InventoryID:     e.InventoryID,
		MovableID:       e.MovableID,
		DestinationNode: e.DestinationNode,
		Quantity:        e.Quantity,
		Version:         e.Version.Int64,
		CreatedAt:       e.CreatedAt,
		UpdatedAt:       e.UpdatedAt,
	}
}

func (e *Inventories) FromDomain(entity *entities.Inventories) {
	e.ID = entity.ID
	e.InventoryID = entity.InventoryID
	e.MovableID = entity.MovableID
	e.DestinationNode = entity.DestinationNode
	e.Quantity = entity.Quantity
	e.Version = optimisticlock.Version{Int64: entity.Version}
	e.CreatedAt = entity.CreatedAt
	e.UpdatedAt = entity.UpdatedAt
}
