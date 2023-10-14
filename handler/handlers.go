package handler

import (
    "first_goland_project/crud"
    "net/http"
)

type Service struct {
    ds crud.Service
}

func GetService() Service {
    hs := Service{crud.GetService()}
    return hs
}

func GetUser(rw http.ResponseWriter, r *http.Request) {

}
