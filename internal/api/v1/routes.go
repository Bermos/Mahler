package v1

import (
	"github.com/Bermos/Platform/internal/app"
	"github.com/danielgtaylor/huma/v2"
	"net/http"
)

func Register(api huma.API, app *app.App) {
	huma.Register(api, huma.Operation{
		OperationID: "ListProjects",
		Description: "List all projects",
		Method:      http.MethodGet,
		Path:        "/api/v1/projects",
		Tags:        []string{"projects"},
	}, app.ListProjects)

}
