package models

import (
	"time"
)

type RegistrationToken struct {
	Value string
	ValidUntil *time.Time
}
