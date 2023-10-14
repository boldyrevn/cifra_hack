package handler

import (
	"encoding/json"
	"first_goland_project/crud"
	"first_goland_project/model"
	"github.com/gorilla/mux"
	"net/http"
)

type Service struct {
	ds crud.Service
}

func GetService() Service {
	hs := Service{crud.GetService()}
	return hs
}

func (s *Service) CloseConn() {
	s.ds.CloseConn()
}

func (s *Service) GetUser(rw http.ResponseWriter, r *http.Request) {
	email := mux.Vars(r)["email"]
	u, err := s.ds.GetUser(email)
	if err != nil {
		m, _ := json.Marshal(model.ErrorMessage{Message: "user is not found"})
		rw.WriteHeader(http.StatusBadRequest)
		_, _ = rw.Write(m)
		return
	}
	rw.WriteHeader(http.StatusOK)
	m, _ := json.Marshal(u)
	_, _ = rw.Write(m)
}
