package main

import (
	"log"

	config "github.com/snykk/fiber-mongo-crud/config"
)

func init() {
	// Setup Config Env
	if err := config.InitializeAppConfig(); err != nil {
		log.Fatal(err)
	}
}

func main() {

}
