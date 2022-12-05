package main

import (
	"hl7/internal/cli"
	"log"
)

func main() {
	if err := cli.Execute(); err != nil {
		log.Fatal(err)
	}
}
