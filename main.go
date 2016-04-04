package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	config := loadConfig()
	app := defaultApplication(config.logger)

	if config.realURL != "" {
		app.CheckCompatibility(config.configPath, config.realURL)
	} else {
		svr, err := app.CreateServer(config.configPath, config.monkeyConfigPath)

		if err != nil {
			log.Fatal(err)
		} else {
			config.logger.Printf("Listening on port %d", config.port)
			err = http.ListenAndServe(fmt.Sprintf(":%d", config.port), svr)
			if err != nil {
				config.logger.Fatal("There was a problem starting the mockingjay server on port %d: %s", config.port, err.Error())
			}
		}
	}
}
