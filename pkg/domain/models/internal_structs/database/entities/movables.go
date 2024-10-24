package entities

import "time"

const (
	Checked   Status = "CHECKED"
	Picked    Status = "PICKED"
	Canalized Status = "CANALIZED"
	Stored    Status = "STORED"
)

type Status string

type Movables struct {
	ID           string
	OriginNode   string
	PartialityID string
	Status       Status
	Version      int64
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Inventories  []Inventories
}
