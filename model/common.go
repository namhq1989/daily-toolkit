package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type (
	// UUID ...
	UUID = uuid.UUID

	// CommonQuery ...
	CommonQuery struct {
		Offset    int
		Limit     int
		Timestamp time.Time
	}
)

// FormatStringToUUID ...
func FormatStringToUUID(s string) (id UUID) {
	id = uuid.FromStringOrNil(s)
	return id
}
