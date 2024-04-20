package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Mayer-04/mongo-api-go/internal/routes"
	"github.com/Mayer-04/mongo-api-go/pkg/database"
	"github.com/joho/godotenv"
)

func main() {

	// Cargar las variables de entorno desde el archivo `.env`.
	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load .env file: %v", err)
	}

	// Crear un contexto con un tiempo de espera de 10 segundos
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Conectarse a MongoDB
	client, err := database.ConnectToMongo(ctx)
	if err != nil {
		panic(err)
	}

	// Configurar una función `defer` para desconectar de MongoDB al finalizar la ejecución de `main()`
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatalf("failed disconnect MongoDB: %v", err)
		}
		log.Println("connection to MongoDB closed successfully")
	}()

	// Configurar el servidor HTTP y definir las rutas
	server := http.NewServeMux()
	routes.SetupRoutes(server)

	// Configurar la ruta base "/v1/" y servir las rutas
	server.Handle("/v1/", http.StripPrefix("/v1", server))

	// Iniciar el servidor HTTP
	log.Fatal(http.ListenAndServe(":8080", server))
}
