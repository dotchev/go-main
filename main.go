package main

import (
	"fmt"

	"github.com/dotchev/go-main/plugin"
	echo "github.com/dotchev/go-plugin"
	"github.com/spf13/viper"
)

func main() {
	viper.ReadInConfig()
	fmt.Println("Hello!")

	r := plugin.PluginRequest{"main", 42}
	echo.Call(&r)
}
