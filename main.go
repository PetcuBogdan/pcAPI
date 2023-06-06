package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)


func initializeRouter(){
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(cors.Handler(cors.Options{
    	AllowedOrigins: []string{"https://*", "http://*"},
    	AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    	AllowedHeaders: []string{"Accepted", "Content-Type", "Authorization"},
    	ExposedHeaders: []string{"Link"},
    	AllowCredentials: true,
    	MaxAge: 300,
	}))

	mux.Get("/cards", GetCards)
	mux.Post("/cardAdd", AddCard)
	mux.Put("/cardEdit", EditCard)
	mux.Delete("/cardDelete", DeleteCard)
	mux.Get("/profiles", GetShipments)
	mux.Post("/profileAdd", AddShipment)
	mux.Put("/profileEdit", EditShipment)
	mux.Delete("/profileDelete", DeleteShipment)
	mux.Get("/tasks", GetTasks)
	mux.Post("/taskAdd", AddTask)
	mux.Put("/taskEdit", EditTask)
	mux.Delete("/taskDelete", DeleteTask)
	mux.Post("/startTask", StartTask)
	mux.Post("/addUser", AddUser)
	mux.Post("/addLicense", AddLicense)
	mux.Delete("/deleteLicense", DeleteLicense)
	log.Fatal(http.ListenAndServe(":8080", mux))

}



func main(){
	initializeRouter()
}