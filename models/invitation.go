package models

import "time"

type Invitation struct {
	ID            int
	InviteeID     int
	InviterID     int
	InvitableID   int
	InvitableType string
	Source        string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
