package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/Forest24/Twitter/middlew"
	"github.com/Forest24/Twitter/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/* Manejdaroes de seteo  de mi puerto,el handler y pongo a ver el server */
func Manejadores() {
	router := mux.NewRouter()
	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
