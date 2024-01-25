package main

import (
	"log"

	"githab.com/command-assistant/internal/assistant"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	assist := assistant.Init()

	if err := assist.Start(); err != nil {
		log.Fatal(err)
	}
}
