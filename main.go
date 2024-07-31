package main

import (
	"github.com/gradinsky-cloudbees/platform-rest-api/cmd"
	"log"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
