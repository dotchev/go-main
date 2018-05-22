package plugin

import (
	uuid "github.com/satori/go.uuid"
)

type PluginRequest struct {
	Name string
	Size int
}

func UUID() string {
	id := uuid.NewV4()
	return id.String()
}
