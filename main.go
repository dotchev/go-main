package main

import (
	"fmt"

	"github.com/dotchev/go-main/plugin"
	echo "github.com/dotchev/go-plugin"
)

func main() {
	fmt.Println("Hello!")

	r := plugin.PluginRequest{"main"}
	echo.Call(&r)
}
