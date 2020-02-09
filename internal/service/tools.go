package service

import "github.com/google/uuid"

func NewShortUUID() string {

	UUID := uuid.New().String()

	return UUID
}
