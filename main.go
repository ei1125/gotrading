package main

import (
	"log"

	"./config"
	"./utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	log.Println("test")
	// fmt.Println(config.Config.ApiKey)
	// fmt.Println(config.Config.ApiSecret)
}
