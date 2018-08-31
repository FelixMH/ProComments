package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/FelixMH/ProComments/commons"
	"github.com/FelixMH/ProComments/migration"
	"github.com/FelixMH/ProComments/routes"
	"github.com/urfave/negroni"
)

func main() {
	var migrate string

	flag.StringVar(&migrate, "migrate", "no", "Generates the migration to the BD")
	flag.IntVar(&commons.Port, "port", 8080, "Port to web server")
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
		Addr:    fmt.Sprintf(":%d", commons.Port),
		Handler: n,
	}
	log.Printf("Iniciado el servidor en http://localhost:%d", commons.Port)
	log.Println(server.ListenAndServe())
	log.Println("Finalizó la ejecución del programa")
}
