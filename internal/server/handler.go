package server

import (
	"bookLibrary/internal/model"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
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

func (s *Server) Authenticate() http.HandlerFunc {
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

func (s *Server) UpdatePassword() http.HandlerFunc {

	type request struct {
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(r)
		id, ok := vars["id"]
		if !ok {
			http.Error(w, errors.New("").Error(), http.StatusBadRequest)
			return
		}

		userID, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		newRequest := &request{}
		if err := json.NewDecoder(r.Body).Decode(newRequest); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user, err := s.store.User().FindByID(uint(userID))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		userAfterUpdating, err := s.store.User().UpdatePassword(user.Email, newRequest.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		s.response(w, r, http.StatusOK, userAfterUpdating)
	}
}

func (s *Server) AddBook() http.HandlerFunc {
	type Book struct {
		Title string `json:"title"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		newTitle := &Book{}
		if err := json.NewDecoder(r.Body).Decode(newTitle); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		vars := mux.Vars(r)
		userId, err := strconv.Atoi(vars["id"])
		//if userId <= 0
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		user, err := s.store.User().FindByID(uint(userId))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		user, err = s.store.User().AddBookByEmail(user.Email, newTitle.Title)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		s.response(w, r, http.StatusOK, user)
	}
}

func (s Server) response(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
