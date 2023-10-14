package handler

import (
	"encoding/json"
	"first_goland_project/crud"
	"first_goland_project/model"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func writeAnswer(rw http.ResponseWriter, ans any, status int) {
	m, _ := json.Marshal(ans)
	rw.WriteHeader(status)
	_, _ = rw.Write(m)
}

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
		writeAnswer(rw, model.ErrorMessage{Message: "user is not found"}, http.StatusBadRequest)
		return
	}
	writeAnswer(rw, u, http.StatusOK)
}

func (s *Service) CreateUser(rw http.ResponseWriter, r *http.Request) {
	d := json.NewDecoder(r.Body)
	var u model.CreateUser
	if err := d.Decode(&u); err != nil {
		writeAnswer(rw, err.Error(), http.StatusBadRequest)
	}
	b, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	u.Password = string(b)
	res, err := s.ds.CreateUser(u)
	if err != nil {
		writeAnswer(rw, model.ErrorMessage{Message: "user already exists"}, http.StatusBadRequest)
		return
	}
	writeAnswer(rw, res, http.StatusOK)
}
