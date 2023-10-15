package handler

import (
	"encoding/json"
	"first_goland_project/crud"
	"first_goland_project/model"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
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
		writeAnswer(rw, model.Message{Message: "user is not found"}, http.StatusBadRequest)
		return
	}
	writeAnswer(rw, u, http.StatusOK)
}

func (s *Service) CreateUser(rw http.ResponseWriter, r *http.Request) {
	d := json.NewDecoder(r.Body)
	var u model.CreateUser
	if err := d.Decode(&u); err != nil {
		writeAnswer(rw, model.Message{Message: err.Error()}, http.StatusBadRequest)
	}
	b, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	u.Password = string(b)
	res, err := s.ds.CreateUser(u)
	if err != nil {
		writeAnswer(rw, model.Message{Message: "user already exists"}, http.StatusBadRequest)
		return
	}
	writeAnswer(rw, res, http.StatusOK)
}

func (s *Service) GetZones(w http.ResponseWriter, r *http.Request) {
	res := s.ds.GetZones()
	writeAnswer(w, res, http.StatusOK)
}

func (s *Service) BookZone(w http.ResponseWriter, r *http.Request) {
	d := json.NewDecoder(r.Body)
	var bz model.BookZone
	if err := d.Decode(&bz); err != nil {
		writeAnswer(w, model.Message{Message: err.Error()}, http.StatusBadRequest)
		return
	}
	if err := s.ds.BookZone(bz.UserID, bz.ZoneID); err != nil {
		writeAnswer(w, model.Message{Message: "zone can't be booked"}, http.StatusBadRequest)
		return
	}
	writeAnswer(w, model.Message{Message: "zone is booked"}, http.StatusOK)
}

func (s *Service) CancelBooking(w http.ResponseWriter, r *http.Request) {
	d := json.NewDecoder(r.Body)
	var bz model.BookZone
	if err := d.Decode(&bz); err != nil {
		writeAnswer(w, model.Message{Message: err.Error()}, http.StatusBadRequest)
		return
	}
	err := s.ds.CancelBooking(bz.UserID, bz.ZoneID)
	if err != nil {
		writeAnswer(w, model.Message{Message: "there is no such booking"}, http.StatusBadRequest)
		return
	}
	writeAnswer(w, model.Message{Message: "booking is canceled"}, http.StatusOK)
}

func (s *Service) CheckBooking(rw http.ResponseWriter, r *http.Request) {
	uid, _ := strconv.Atoi(r.URL.Query().Get("userID"))
	zid, _ := strconv.Atoi(r.URL.Query().Get("zoneID"))
	if uid == 0 || zid == 0 {
		writeAnswer(rw, model.Message{Message: "wrong query parameters"}, http.StatusBadRequest)
		return
	}
	isBooked, err := s.ds.CheckBooking(uid, zid)
	if err != nil {
		writeAnswer(rw, model.Message{Message: err.Error()}, http.StatusBadRequest)
		return
	}
	writeAnswer(rw, model.Message{Message: strconv.FormatBool(isBooked)}, http.StatusOK)
}

func (s *Service) GetStat(w http.ResponseWriter, r *http.Request) {
	email := mux.Vars(r)["email"]
	us, err := s.ds.GetStat(email)
	if err != nil {
		writeAnswer(w, model.Message{Message: "can't find user's stat"}, http.StatusBadRequest)
		return
	}
	writeAnswer(w, us, http.StatusOK)
}

func (s *Service) GetInvitations(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	if id == 0 {
		writeAnswer(w, model.Message{Message: "wrong query parameters"}, http.StatusBadRequest)
		return
	}
	inv := s.ds.GetInvitations(id)
	writeAnswer(w, inv, http.StatusOK)
}

func (s *Service) GetEvents(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	if id == 0 {
		writeAnswer(w, model.Message{Message: "wrong query parameters"}, http.StatusBadRequest)
		return
	}
	evs := s.ds.GetEvents(id)
	writeAnswer(w, evs, http.StatusOK)
}
