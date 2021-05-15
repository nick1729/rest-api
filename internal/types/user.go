package types

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	Firstname string
	Lastname  string
	Email     string
	Age       uint64
	Created   time.Time
}
