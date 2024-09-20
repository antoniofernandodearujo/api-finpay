package main

import (
	"log"
	"net/http"

	"api-finpay/src/routes"

	"github.com/gorilla/mux"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

func main() {
	r := mux.NewRouter()
	routes.SetupRoutes(r)

	http.Handle("/", r)

	log.Println("Servidor rodando na porta 8080")

	log.Fatal(http.ListenAndServe(":8080", r))
}
