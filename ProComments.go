package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/FelixMH/ProComments/migration"
	"github.com/FelixMH/ProComments/routes"
	"github.com/urfave/negroni"
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

	// Inicia las rutas
	router := routes.InitRoutes()

	// Inicia los middlewares
	n := negroni.Classic()
	n.UseHandler(router)

	// Inicia el servidor
	server := &http.Server{
		Addr:    ":8080",
		Handler: n,
	}
	log.Println("Iniciado el servidor en http://localhost:8080")
	log.Println(server.ListenAndServe())
	log.Println("Finalizó la ejecución del programa")
}
