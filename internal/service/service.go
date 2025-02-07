package service

import (
	"github.com/Bermos/Platform/internal/resource"
	"github.com/google/uuid"
)

type Service struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Resource resource.Resource
}
