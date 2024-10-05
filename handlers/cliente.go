package handlers

import (
    "goclients/db"
    "goclients/models"
    "encoding/json"
    "log"
    "net/http"
)

func GetClients(w http.ResponseWriter, r *http.Request)  {

	database := db.Connect()
	defer database.Close()

    filas, err := database.Query("SELECT id, nombre, email, telefono FROM clientes")
	
	
	if err != nil {
		log.Fatalf("Error al obtener clientes: %v", err)
		json.NewEncoder(w).Encode(models.APIResponse{
			Code: http.StatusInternalServerError,
			Data: "Error al procesar los datos del cliente",
		})
		return
	}	

	defer filas.Close()

	var clients []models.Cliente
	for filas.Next() {
        var cliente models.Cliente
        filas.Scan(&cliente.ID, &cliente.Nombre, &cliente.Email, &cliente.Telefono)
        clients = append(clients, cliente)
    }

	json.NewEncoder(w).Encode(models.APIResponse{
        Code: http.StatusOK,
        Data: clients,
    })
}

func CreateClient(w http.ResponseWriter, r *http.Request)  {
	
	database := db.Connect()
	defer database.Close()

	var cliente models.Cliente
    json.NewDecoder(r.Body).Decode(&cliente)

    insertar, err := database.Prepare("INSERT INTO clientes(nombre, email, telefono) VALUES(?, ?, ?)")
    if err != nil {
        log.Fatalf("Error al preparar la inserción: %v", err)
    }
    insertar.Exec(cliente.Nombre, cliente.Email, cliente.Telefono)

	json.NewEncoder(w).Encode(models.APIResponse{
        Code: http.StatusOK,
        Data: cliente,
    })
}