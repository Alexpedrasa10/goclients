package models

type Cliente struct {
    ID       int    `json:"id"`
	Nombre	 string `json:"nombre"`
	Email    string `json:"email"`
	Telefono string `json:"telefono"` 
}