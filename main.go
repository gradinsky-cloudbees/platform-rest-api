package main

import (
	"gradinsky-cloudbees/platform-rest-api/cmd"
	"log"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
