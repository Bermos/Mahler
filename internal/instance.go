package internal

import (
	"github.com/Bermos/Platform/internal/project"
	"github.com/Bermos/Platform/internal/resource"
)

type Instance struct {
	Name               string
	Projects           []*project.Project
	AvailableResources []resource.Resource
}

func (i *Instance) AddProject(p *project.Project) {
	i.Projects = append(i.Projects, p)
}
