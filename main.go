package main

import (
    "goclients/handlers"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()

    // Rutas del CRUD
    r.HandleFunc("/clientes", handlers.GetClients).Methods("GET")
    r.HandleFunc("/clientes", handlers.CreateClient).Methods("POST")

    log.Println("Servidor escuchando en http://localhost:8000")
    log.Fatal(http.ListenAndServe(":8000", r))
}
