package db

import (
    "database/sql"
    "log"

    _ "github.com/mattn/go-sqlite3" // Import necesario para SQLite
)

func Connect() *sql.DB {
	
	db, err := sql.Open("sqlite3", "./clientes.db")
    if err != nil {
        log.Fatalf("Error al abrir la base de datos: %v", err)
    }

    // Crear la tabla de clientes si no existe
    crearTabla := `CREATE TABLE IF NOT EXISTS clientes (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        nombre TEXT,
        email TEXT,
        telefono TEXT
    );`

    if _, err := db.Exec(crearTabla); err != nil {
        log.Fatalf("Error al crear la tabla: %v", err)
    }

    return db
}