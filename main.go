package main

import (
    "goclients/handlers"
    "log"
    "net/http"
    "goclients/db"
    "github.com/gorilla/mux"
)

func main() {
    
    r := mux.NewRouter()

    // Init DB
    database := db.Connect()
    defer database.Close()

    handlers.InitDB(database)

    // Rutas del CRUD
    r.HandleFunc("/clientes", handlers.GetClients).Methods("GET")
    r.HandleFunc("/clientes", handlers.CreateClient).Methods("POST")
    
    r.HandleFunc("/cliente/{id}", handlers.UpdateCliente).Methods("PUT")
    r.HandleFunc("/cliente/{id}", handlers.GetClient).Methods("GET")
    r.HandleFunc("/cliente/{id}", handlers.DeleteClient).Methods("DELETE")
    

    log.Println("Servidor escuchando en http://localhost:8000")
    log.Fatal(http.ListenAndServe(":8000", r))
}
