package main

import (
	"first_goland_project/handler"
	gh "github.com/gorilla/handlers"
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

	cors := gh.CORS(
		gh.AllowCredentials(),
		gh.AllowedMethods([]string{"GET", "POST", "DELETE"}),
		gh.AllowedOrigins([]string{"*"}),
		gh.AllowedHeaders([]string{"Content-Type"}),
	)

	router := mux.NewRouter()
	router.HandleFunc("/api/user/{email}", h.GetUser).Methods("GET")
	router.HandleFunc("/api/user", h.CreateUser).Methods("POST")
	router.HandleFunc("/api/zone", h.GetZones).Methods("GET")
	router.HandleFunc("/api/zone/book", h.BookZone).Methods("POST")
	router.HandleFunc("/api/zone/book", h.CheckBooking).Methods("GET")
	router.HandleFunc("/api/zone/book", h.CancelBooking).Methods("DELETE")
	router.HandleFunc("/api/user/{email}/stat", h.GetStat).Methods("GET")
	router.HandleFunc("/api/invitations/{id}", h.GetInvitations).Methods("GET")
	router.HandleFunc("/api/events/{id}", h.GetEvents).Methods("GET")
	log.Fatal(http.ListenAndServe(":9777", cors(router)))
}
