package routes

import "net/http"

func SetupRoutes(server *http.ServeMux) {

	server.HandleFunc("GET /products", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	server.HandleFunc("GET /products/{id}", func(w http.ResponseWriter, r *http.Request) {})
	server.HandleFunc("POST /products", func(w http.ResponseWriter, r *http.Request) {})
	server.HandleFunc("PUT /products/{id}", func(w http.ResponseWriter, r *http.Request) {})
	server.HandleFunc("DELETE /products/{id}", func(w http.ResponseWriter, r *http.Request) {})

}
