package main

import (
	"fmt"

	"./config"
)

func main() {
	fmt.Println(config.Config.ApiKey)
	fmt.Println(config.Config.ApiSecret)
}
