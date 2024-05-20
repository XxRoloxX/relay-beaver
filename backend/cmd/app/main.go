package main

import (
	"backend/internal/database"
	"backend/internal/request"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	// fmt.Println("Hello, World!")
	router := mux.NewRouter()
	godotenv.Load()
	database := database.InitializeDatabase()

	requestRepository := request.RequestMongoRepository{Db: *database}
	requestService := request.RequestService{Repo: requestRepository}
	requestHandler := request.RequestHandler{Service: requestService}

	router.HandleFunc("/", handler).Methods("GET")
	router.HandleFunc("/requests", requestHandler.GetRequestsHandler).Methods("GET")

	http.ListenAndServe(":8080", router)
}
