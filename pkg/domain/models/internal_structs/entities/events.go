package entities

import "time"

const (
	PickStarted     Type = "PICK_STARTED"
	PickCreated     Type = "PICK_CREATED"
	PickFinished    Type = "PICK_FINISHED"
	PutAwayCreated  Type = "PUT_AWAY_CREATED"
	PutAwayFinished Type = "PUT_AWAY_FINISHED"
	HUStarted       Type = "HU_STARTED"
	PickRefused     Type = "PICK_REFUSED"
	PutAwayRefused  Type = "PUT_AWAY_REFUSED"
)

type Type string

type Events struct {
	ID           uint
	OriginNode   string
	PartialityID string
	MovableID    string
	Type         Type
	Rehydration  bool
	ArrivedLate  bool
	HadSorting   bool
	EventID      string
	EventData    []byte
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
