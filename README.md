# Clientes API

Esta API permite la creación, lectura, actualización y eliminación (CRUD) de clientes. Está construida en Golang utilizando `net/http` y SQL (SQLite).

## Requisitos

- Docker
- Docker Compose

## Instalación con Docker

Sigue estos pasos para ejecutar la API utilizando Docker.

### 1. Clona este repositorio:

```bash
git clone https://github.com/Alexpedrasa10/goclients.git
cd goclients
```

### 2. Construir y ejecutar la aplicación con Docker

Construir y ejecutar la aplicación con Docker

```bash
sudo docker-compose up --build
```
Este comando:

- Construirá la imagen Docker con la API de clientes.
- Mapeará el puerto 8080 del contenedor al puerto 8080 de tu máquina local.
- Montará el archivo de la base de datos clients.db para que los datos sean persistentes.

### 3. Acceder a la API

Una vez construida la imagen puedes acceder a la API con la siguiente URL:

```bash
http://localhost:8080/clientes
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
        "code": 200,
        "data": {
            "id": 3,
            "nombre": "Ruben Botta",
            "email": "bottis@example.com",
            "telefono": "3515655698"
        }
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

