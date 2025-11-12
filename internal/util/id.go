package util

import "github.com/google/uuid"

// GenerateID returns a new unique string ID.
func GenerateID() string {
    return uuid.NewString()
}
