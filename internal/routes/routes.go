package routes

import (
	"net/http"

	"github.com/Mayer-04/mongo-api-go/internal/handler"
)

func SetupRoutes(server *http.ServeMux) {

	server.HandleFunc("GET /products", handler.GetProducts)
	server.HandleFunc("GET /products/{id}", handler.GetproductById)
	server.HandleFunc("POST /products", handler.CreateProduct)
	server.HandleFunc("PUT /products/{id}", handler.UpdateProduct)
	server.HandleFunc("DELETE /products/{id}", handler.DeleteProduct)

}
