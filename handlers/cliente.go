package handlers

import (
    "goclients/models"
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
	"strconv"
    "database/sql"
    "reflect"
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
    client, err := cliente.GetClientById(database, id)

    if err != nil {

        var code = http.StatusNotFound
        var errMsg = err.Error()

        if errMsg != "client_not_found" {
            code = http.StatusInternalServerError
        } 

        response := models.APIResponse{Code: code}
        response.RespondWithError(w, errMsg)
        return
    }

    // Mapa de campos a actualizar con reflect
    json.NewDecoder(r.Body).Decode(&cliente)
    updates := make(map[string]interface{})
    v := reflect.ValueOf(cliente)
    t := v.Type()

    for i := 0; i < v.NumField(); i++ {
        field := t.Field(i)
        value := v.Field(i)

        if !value.IsZero() {
            updates[field.Tag.Get("json")] = value.Interface()
        }
    }


    if len(updates) == 0 {
        response := models.APIResponse{Code: http.StatusBadRequest}
        response.RespondWithError(w, "No se proporcionaron campos para actualizar")
        return
    }

    // SQL Dinamic
    query := "UPDATE clientes SET "
    args := []interface{}{}

    for columName, value := range updates {
        query += columName + " = ?, "
        args = append(args, value)
    }
    
    // Eliminar la última coma y espacio
    query = query[:len(query)-2]
    query += " WHERE id = ?"
    args = append(args, id)

    _, err = database.Exec(query, args...)
    if err != nil {
        response := models.APIResponse{Code: http.StatusInternalServerError,}
        response.RespondWithError(w,"Error al actualizar el usuario")
        return
    }

    client, err = cliente.GetClientById(database, id)

    if err != nil {

        var code = http.StatusNotFound
        var errMsg = err.Error()

        if errMsg != "client_not_found" {
            code = http.StatusInternalServerError
        } 

        response := models.APIResponse{Code: code}
        response.RespondWithError(w, errMsg)
        return
    }

    response := models.APIResponse{Code: http.StatusOK, Data: client}
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
    client, err := cliente.GetClientById(database, id)

    if err != nil {

        var code = http.StatusNotFound
        var errMsg = err.Error()

        if errMsg != "cliente no encontrado" {
            code = http.StatusInternalServerError
        } 

        response := models.APIResponse{Code: code}
        response.RespondWithError(w, errMsg)
        return
    }

    if err != nil {

        response := models.APIResponse{Code: http.StatusInternalServerError}
        response.RespondWithError(w, "Error al obtener el cliente")
        return
    }

    response := models.APIResponse{Code: http.StatusOK, Data: client}
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

    var cliente models.Cliente
    _, err = cliente.GetClientById(database, id)

    if err != nil {

        var code = http.StatusNotFound
        var errMsg = err.Error()

        if errMsg != "cliente no encontrado" {
            code = http.StatusInternalServerError
        } 

        response := models.APIResponse{Code: code}
        response.RespondWithError(w, errMsg)
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