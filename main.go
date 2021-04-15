package main

import (
	"log"

	"github.com/Forest24/Twitter/bd"
	"github.com/Forest24/Twitter/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()

}
