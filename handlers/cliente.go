package handlers

import (
    "goclients/models"
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
	"strconv"
    "database/sql"
)

var database *sql.DB 

func InitDB(db *sql.DB) {
    database = db
}

func GetClients(w http.ResponseWriter, r *http.Request)  {

    filas, err := database.Query("SELECT id, nombre, email, telefono FROM clientes")
	
	if err != nil {

        response := models.APIResponse{Code: http.StatusOK}
        response.RespondWithError(w, "Cliente eliminado con éxito")
		return
	}	

	defer filas.Close()

	var clients []models.Cliente
	for filas.Next() {
        var cliente models.Cliente
        filas.Scan(&cliente.ID, &cliente.Nombre, &cliente.Email, &cliente.Telefono)
        clients = append(clients, cliente)
    }

    response := models.APIResponse{Code: http.StatusOK, Data: clients}
    response.RespondWithJSON(w)
}


func CreateClient(w http.ResponseWriter, r *http.Request)  {
	
	var cliente models.Cliente
    json.NewDecoder(r.Body).Decode(&cliente)

    insertar, err := database.Prepare("INSERT INTO clientes(nombre, email, telefono) VALUES(?, ?, ?)")
    if err != nil {
        log.Fatalf("Error al preparar la inserción: %v", err)
    }

    result, err := insertar.Exec(cliente.Nombre, cliente.Email, cliente.Telefono)
    if err != nil {

        response := models.APIResponse{Code: http.StatusInternalServerError}
        response.RespondWithError(w, "Error al insertar cliente.")
        return
    }

    lastInsertId, err := result.LastInsertId()

    if err != nil {
        response := models.APIResponse{Code: http.StatusInternalServerError}
        response.RespondWithError(w, "Error al obtener el ID del cliente")
        return
    }

    // Asignar el ID obtenido al cliente
    cliente.ID = int(lastInsertId)

    response := models.APIResponse{Code: http.StatusOK, Data: cliente}
    response.RespondWithJSON(w)
}


func UpdateCliente(w http.ResponseWriter, r *http.Request)  {

    params := mux.Vars(r)
    id, err := strconv.Atoi(params["id"])

    if err != nil {
        response := models.APIResponse{Code: http.StatusBadRequest}
        response.RespondWithError(w, "ID Invalido")
        return
    }

    var cliente models.Cliente
    json.NewDecoder(r.Body).Decode(&cliente)

	_, err = database.Exec("UPDATE clientes SET nombre = ?, email = ?, Telefono = ? WHERE id = ?", cliente.Nombre, cliente.Email, cliente.Telefono, id)

    response := models.APIResponse{Code: http.StatusOK, Data: cliente}
    response.RespondWithJSON(w)
}

func GetClient(w http.ResponseWriter, r *http.Request)  {

    params := mux.Vars(r)
    id, err := strconv.Atoi(params["id"])

    if err != nil {
        response := models.APIResponse{Code: http.StatusBadRequest}
        response.RespondWithError(w, "ID Invalido")
        return
    }

    var cliente models.Cliente
    
    err = database.QueryRow(
        "SELECT id, nombre, email, telefono FROM clientes WHERE id = ?", id).Scan(
            &cliente.ID, &cliente.Nombre, &cliente.Email, &cliente.Telefono)

    if err != nil {

        response := models.APIResponse{Code: http.StatusInternalServerError}
        response.RespondWithError(w, "Error al obtener el cliente")
        return
    }

    response := models.APIResponse{Code: http.StatusOK, Data: cliente}
    response.RespondWithJSON(w)
}

func DeleteClient(w http.ResponseWriter, r *http.Request)  {

    params := mux.Vars(r)
    id, err := strconv.Atoi(params["id"])

    if err != nil {
        response := models.APIResponse{Code: http.StatusBadRequest,}
        response.RespondWithError(w,"ID Invalido")
        return
    }

	_, err = database.Exec("DELETE from clientes WHERE id = ?", id)

    if err != nil {
        response := models.APIResponse{Code: http.StatusInternalServerError,}
        response.RespondWithError(w,"Error al borrar el usuario")
        return
    }

    response := models.APIResponse{
        Code: http.StatusOK,
        Data: "Cliente eliminado con éxito",
    }
    response.RespondWithJSON(w)
}