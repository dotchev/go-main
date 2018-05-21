package plugin

import (
	"fmt"

	m "github.com/dotchev/go-main/plugin"
)

func Call(r *m.PluginRequest) {
	fmt.Println("Plugin: ", r.Name)
}
