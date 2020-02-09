package service

import (
	"strings"

	"github.com/google/uuid"
)

// NewShortUUID returns a short UUID
func NewShortUUID() string {

	UUID := uuid.New().String()

	return strings.Split(UUID, "-")[0]
}
