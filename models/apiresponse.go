package models

import (
    "encoding/json"
    "net/http"
)

type APIResponse struct {
    Code int         `json:"code"` 
    Data interface{} `json:"data"`
}

// Método para enviar una respuesta JSON
func (resp *APIResponse) RespondWithJSON(w http.ResponseWriter) {
    response, _ := json.Marshal(resp)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(resp.Code)
    w.Write(response)
}

// Método para enviar un error
func (resp *APIResponse) RespondWithError(w http.ResponseWriter, message string) {
    resp.Data = map[string]string{"error": message}
    resp.RespondWithJSON(w)
}
