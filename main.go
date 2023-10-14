package main

import (
    "context"
    "first_goland_project/crud"
    "github.com/gorilla/mux"
    "github.com/joho/godotenv"
    "log"
)

func init() {
    if err := godotenv.Load(); err != nil {
        log.Fatal("unable to load .env variables")
    }
}

func main() {
    conn := crud.connectDb()
    defer conn.Close(context.Background())

    router := mux.NewRouter()
    router.HandleFunc("/user/{email}")
}
