# 1. Imagen base oficial de Golang (no Alpine, para tener soporte CGO)
FROM golang:1.20

# 2. Seteamos el directorio de trabajo dentro del contenedor
WORKDIR /app

# 3. Copiamos el go.mod y go.sum para descargar dependencias
COPY go.mod go.sum ./

# 4. Descargar dependencias
RUN go mod download

# 5. Copiar el código de la API en el contenedor
COPY . .

# 6. Habilitar CGO y compilar con soporte para SQLite
ENV CGO_ENABLED=1
RUN go build -o /goclients

# 7. Definir el puerto que se expondrá
EXPOSE 8080

# 8. Comando de ejecución de la API
CMD ["/goclients"]
