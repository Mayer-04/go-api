package handler

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/Mayer-04/mongo-api-go/internal/models"
	"github.com/Mayer-04/mongo-api-go/pkg/database"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var product models.Product

type response struct {
	Success bool           `json:"success"`
	Message string         `json:"message"`
	Data    models.Product `json:"data,omitempty"`
}

func GetProducts(w http.ResponseWriter, r *http.Request) {

	collection := database.GetCollection("products")

	filter := bson.D{}

	cursor, err := collection.Find(context.Background(), filter)

	if err != nil {
		http.Error(w, "failed documents", http.StatusNotFound)
		return
	}

	defer cursor.Close(context.Background())

	var products []models.Product

	for cursor.Next(context.Background()) {
		var product models.Product
		err := cursor.Decode(&product)
		if err != nil {
			log.Fatal(err)
		}

		products = append(products, product)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)

}

func GetproductById(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")

	collection := database.GetCollection("products")

	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": objectID}

	err = collection.FindOne(context.Background(), filter).Decode(&product)

	if errors.Is(err, mongo.ErrNoDocuments) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	res, err := json.Marshal(product)

	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	collection := database.GetCollection("products")

	validate := validator.New()

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := validate.Struct(product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := collection.InsertOne(context.Background(), product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	collection := database.GetCollection("products")

	validator := validator.New()

	productID, err := primitive.ObjectIDFromHex(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid user ID", http.StatusBadRequest)
		return
	}

	filter := bson.M{"_id": productID}

	var productUpdate models.Product

	err = json.NewDecoder(r.Body).Decode(&productUpdate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validator.Struct(productUpdate); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	update := bson.M{"$set": productUpdate}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if result.ModifiedCount == 0 {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(productUpdate)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {

	collection := database.GetCollection("products")

	productID := r.PathValue("id")

	objectID, err := primitive.ObjectIDFromHex(productID)

	if err != nil {
		http.Error(w, "invalid user ID", http.StatusBadRequest)
		return
	}

	filter := bson.M{"_id": objectID}

	_, err = collection.DeleteOne(context.Background(), filter)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := response{
		Success: true,
		Message: "user deleted successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}
