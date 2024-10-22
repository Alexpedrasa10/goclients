# Clientes API

Esta API permite la creación, lectura, actualización y eliminación (CRUD) de clientes. Está construida en Golang utilizando `net/http` y SQL (en este caso, con SQLite).

## Requisitos

- Go 1.16 o superior
- SQLite3

## Instalación

1. Clona este repositorio:

    ```bash
    git clone https://github.com/usuario/proyecto-clientes.git
    cd proyecto-clientes
    ```

2. Instala las dependencias necesarias:

    ```bash
    go mod tidy
    ```

3. Configura la base de datos:

   Asegúrate de tener una base de datos SQLite lista. Puedes crear una nueva usando:

    ```bash
    sqlite3 clientes.db < schema.sql
    ```

4. Ejecuta la aplicación:

    ```bash
    go run main.go
    ```

## Endpoints

### Obtener todos los clientes

- **Ruta:** `GET /clientes`
- **Descripción:** Retorna una lista de todos los clientes.
- **Ejemplo de respuesta:**

    ```json
    {
        "code": 200,
        "data": [
            {
                "id": 1,
                "nombre": "Cristal",
                "email": "cristal@example.com",
                "telefono": "3515655698"
            },
            {
                "id": 2,
                "nombre": "Alex",
                "email": "alex@example.com",
                "telefono": "3515655698"
            },
            {
                "id": 3,
                "nombre": "Ruben Botta",
                "email": "bottis@example.com",
                "telefono": "3515655698"
            }
        ]
    }
    ```

### Obtener un cliente

- **Ruta:** `GET /cliente/{id}`
- **Descripción:** Retorna los datos de un cliente específico por ID.
- **Parámetro de ruta:**
  - `id`: ID del cliente.
- **Ejemplo de respuesta:**

    ```json
    {
        "id": 3,
        "nombre": "Ruben Botta",
        "email": "bottis@example.com",
        "telefono": "3515655698"
    }
    ```

### Crear un cliente

- **Ruta:** `POST /clientes`
- **Descripción:** Crea un nuevo cliente.
- **Ejemplo de solicitud:**

    ```json
    {
      "nombre": "Ana Gómez",
      "email": "ana@example.com",
      "telefono": "987654321"
    }
    ```

- **Ejemplo de respuesta:**

    ```json
    {
        "code": 200,
        "data": {
            "id": 4,
            "nombre": "Ana Gómez",
            "email": "ana@example.com",
            "telefono": "987654321"
        }
    }
    ```

### Actualizar un cliente

- **Ruta:** `PUT /cliente/{id}`
- **Descripción:** Actualiza los datos de un cliente existente. Solo pasar los datos que desea actualizar.
- **Parámetro de ruta:**
  - `id`: ID del cliente.
- **Ejemplo de solicitud:**

    ```json
    {
      "email": "ana_updated@example.com",
    }
    ```

- **Ejemplo de respuesta:**

    ```json
    {
        "code": 200,
        "data": {
            "id": 4,
            "nombre": "Ana Gómez",
            "email": "ana_updated@example.com",
            "telefono": "3515655698"
        }
    }
    ```

### Eliminar un cliente

- **Ruta:** `DELETE /cliente/{id}`
- **Descripción:** Elimina un cliente existente por ID.
- **Parámetro de ruta:**
  - `id`: ID del cliente.
- **Ejemplo de respuesta:**

    ```json
    {
        "code": 200,
        "data": "Cliente eliminado con éxito"
    }
    ```

