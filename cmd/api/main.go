package main

import (
	"log"
	"net/http"

	"miniApi_BRM/internal/db"
	httpHandler "miniApi_BRM/internal/http"
	"miniApi_BRM/internal/repository"
	"miniApi_BRM/internal/service"

	"github.com/gorilla/mux"
)

func main() {

	dbConfig := db.Config{
		Host:     "usersapi-miniapi01.j.aivencloud.com",
		Port:     "20927",
		User:     "avnadmin",
		Password: "******",
		Database: "defaultdb",
		SSLMode:  "REQUIRED",
	}

	database, err := db.NewConnection(dbConfig)
	if err != nil {
		log.Fatalf("Error conectando a la base de datos: %v", err)
	}
	defer database.Close()

	userRepo := repository.NewMySQLUserRepository(database)
	userService := service.NewUserService(userRepo)
	userHandler := httpHandler.NewUserHandler(userService)

	router := mux.NewRouter()

	router.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	router.HandleFunc("/users", userHandler.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", userHandler.GetUserByID).Methods("GET")
	router.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods("GET")

	log.Println(" Servidor iniciado en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
