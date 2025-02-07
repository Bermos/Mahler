package app

import (
	"context"
)

func NewApp() *App {
	return &App{}
}

type App struct {
}

func (a App) ListProjects(ctx context.Context, i *struct{}) (*struct{}, error) {
	return nil, nil
}
