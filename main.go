package main

import (
	"first_goland_project/handler"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("unable to load .env variables")
	}
}

func main() {
	h := handler.GetService()
	defer h.CloseConn()

	router := mux.NewRouter()
	router.HandleFunc("/user/{email}", h.GetUser).Methods("GET")
	router.HandleFunc("/user", h.GetUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":9777", router))
}
