package main

import (
	"log"

	"github.com/rubenv/fastyaml"
	"github.com/rubenv/fastyaml/fixtures"
)

func main() {
	err := do()
	if err != nil {
		log.Fatal(err)
	}
}

func do() error {
	err := fastyaml.GenerateFile("fixtures", fixtures.Basic{}, "fixtures/parse_basic.go")
	if err != nil {
		return err
	}
	return nil
}
