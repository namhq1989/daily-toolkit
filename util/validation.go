package util

import uuid "github.com/satori/go.uuid"

// IsValidUUID ...
func IsValidUUID(id uuid.UUID) bool {
	return id.String() != ""
}
