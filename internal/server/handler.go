package server

import (
	"bookLibrary/internal/model"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func (s *Server) handleUsersCreate() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		newUser := model.User{
			Email:    req.Email,
			Password: req.Password,
		}
		if err := s.store.User().Create(&newUser); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		s.response(w, r, http.StatusCreated, newUser)
	}
}

func (s *Server) handleSession() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		userRep := s.store.User()
		user, err := userRep.FindByEmail(req.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte(req.Password)); err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
		}

		s.response(w, r, http.StatusOK, nil)

	}
}

func (s Server) response(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
