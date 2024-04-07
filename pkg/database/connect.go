package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client

func ConnectToMongo(ctx context.Context) (*mongo.Client, error) {

	mongoUri := os.Getenv("MONGODB_URI")

	// Configura las opciones del cliente MongoDB
	clientOptions := options.Client().ApplyURI(mongoUri)

	var err error
	// Conectar al servidor de MongoDB
	client, err = mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("connect MongoDB: %w", err)
	}

	// Verificar la conexión con el servidor de MongoDB
	if err := pingMongo(client); err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("ping MongoDB: %w", err)
	}

	log.Println("connection to MongoDB successfully established")

	return client, nil

}

func pingMongo(client *mongo.Client) error {

	// Configurar un contexto con un tiempo de espera de 2 segundos
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	defer cancel()

	err := client.Ping(ctx, readpref.Primary())

	return err

}

func GetCollection(col string) *mongo.Collection {
	database := client.Database("go-crud-mongo")

	collection := database.Collection(col)

	return collection

}
