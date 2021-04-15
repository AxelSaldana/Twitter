package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/Forest24/mux"
	"github.com/rs/cors"
)

/* Manejdaroes de seteo  de mi puerto,el handler y pongo a ver el server */
func Manejadores() {
	router := mux.NewRouter()

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
