package entities

import "time"

type Inventories struct {
	ID              uint
	InventoryID     string
	MovableID       string
	DestinationNode string
	Quantity        int
	Version         int64
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
