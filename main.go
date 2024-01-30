package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/OpsInc/enroller-client/cmd"
)

func main() {
	// Remove date and time from log output
	log.SetFlags(0)

	_, err := os.Stat(".env")
	if err == nil {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Unable to load .env file with error: %v", err)
		}
	}

	cmd.Execute()
}
