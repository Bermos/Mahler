package project

import (
	"github.com/Bermos/Platform/internal/service"
	"github.com/google/uuid"
)

type Project struct {
	ID       uuid.UUID          `json:"id"`
	Name     string             `json:"name"`
	Services []*service.Service `json:"services"`
}
