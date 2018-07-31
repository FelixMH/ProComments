package main

import (
	"flag"
	"log"

	"github.com/FelixMH/ProComments/migration"
)

func main() {
	var migrate string

	flag.StringVar(&migrate, "migrate", "no", "Generates the migration to the BD")
	flag.Parse()
	if migrate == "yes" {
		log.Println("Migration started.....")
		migration.Migrate()
		log.Println("Migration ended")
	}
}
