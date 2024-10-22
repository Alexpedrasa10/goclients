package models

import (
    "database/sql"
    "errors"
)

type Cliente struct {
    ID       int    `json:"id"`
	Nombre	 string `json:"nombre"`
	Email    string `json:"email"`
	Telefono string `json:"telefono"` 
}

func (c *Cliente) GetClientById(db *sql.DB, id int) (*Cliente, error) {
    var cliente Cliente

    // Consulta para buscar el cliente por ID
    err := db.QueryRow("SELECT id, nombre, email, telefono FROM clientes WHERE id = ?", id).
        Scan(&cliente.ID, &cliente.Nombre, &cliente.Email, &cliente.Telefono)

    if err != nil {
        if err == sql.ErrNoRows {
            return nil, errors.New("client_not_found")
        }

		return nil, err
    }

    return &cliente, nil
}
