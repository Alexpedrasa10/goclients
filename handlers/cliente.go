package handlers

import (
    "goclients/db"
    "goclients/models"
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
	"strconv"
)

func GetClients(w http.ResponseWriter, r *http.Request)  {

	database := db.Connect()
	defer database.Close()

    filas, err := database.Query("SELECT id, nombre, email, telefono FROM clientes")
	
	if err != nil {
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

    result, err := insertar.Exec(cliente.Nombre, cliente.Email, cliente.Telefono)
    if err != nil {
        json.NewEncoder(w).Encode(models.APIResponse{
            Code: http.StatusInternalServerError,
            Data: "Error al insertar cliente.",
        })
        return
    }

    lastInsertId, err := result.LastInsertId()
    if err != nil {
        json.NewEncoder(w).Encode(models.APIResponse{
            Code: http.StatusInternalServerError,
            Data: "Error al obtener el ID del cliente",
        })
        return
    }

    // Asignar el ID obtenido al cliente
    cliente.ID = int(lastInsertId)

	json.NewEncoder(w).Encode(models.APIResponse{
        Code: http.StatusOK,
        Data: cliente,
    })
}


func UpdateCliente(w http.ResponseWriter, r *http.Request)  {

    params := mux.Vars(r)
    id, err := strconv.Atoi(params["id"])

    if err != nil {
		json.NewEncoder(w).Encode(models.APIResponse{
			Code: http.StatusBadRequest,
			Data: "ID Invalido",
		})
        return
    }

    database := db.Connect()
	defer database.Close()

    var cliente models.Cliente
    json.NewDecoder(r.Body).Decode(&cliente)

	_, err = database.Exec("UPDATE clientes SET nombre = ?, email = ?, Telefono = ? WHERE id = ?", cliente.Nombre, cliente.Email, cliente.Telefono, id)

    json.NewEncoder(w).Encode(models.APIResponse{
        Code: http.StatusOK,
        Data: cliente,
    })
}

func GetClient(w http.ResponseWriter, r *http.Request)  {

    params := mux.Vars(r)
    id, err := strconv.Atoi(params["id"])

    if err != nil {
		json.NewEncoder(w).Encode(models.APIResponse{
			Code: http.StatusBadRequest,
			Data: "ID Invalido",
		})
        return
    }

	database := db.Connect()
	defer database.Close()

    var cliente models.Cliente
    
    err = database.QueryRow(
        "SELECT id, nombre, email, telefono FROM clientes WHERE id = ?", id).Scan(
            &cliente.ID, &cliente.Nombre, &cliente.Email, &cliente.Telefono)

    if err != nil {
        
        json.NewEncoder(w).Encode(models.APIResponse{
            Code: http.StatusInternalServerError,
            Data: "Error al obtener el cliente",
        })
        return
    }

	json.NewEncoder(w).Encode(models.APIResponse{
        Code: http.StatusOK,
        Data: cliente,
    })
}

func DeleteClient(w http.ResponseWriter, r *http.Request)  {

    params := mux.Vars(r)
    id, err := strconv.Atoi(params["id"])

    if err != nil {
		json.NewEncoder(w).Encode(models.APIResponse{
			Code: http.StatusBadRequest,
			Data: "ID Invalido",
		})
        return
    }

    database := db.Connect()
	defer database.Close()

	_, err = database.Exec("DELETE from clientes WHERE id = ?", id)

    if err != nil {
        json.NewEncoder(w).Encode(models.APIResponse{
            Code: http.StatusInternalServerError,
            Data: "Error al borrar el usuario",
        })
        return
    }

    json.NewEncoder(w).Encode(models.APIResponse{
        Code: http.StatusOK,
        Data: "Usuario borrado con exito",
    })
}