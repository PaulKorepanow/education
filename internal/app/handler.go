package app

import (
	"bookLibrary/internal/model"
	"encoding/json"
	"net/http"
)

func (s *Server) handleUsersCreate() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
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
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		s.response(w, r, http.StatusOK, newUser)
	}
}

func (s Server) response(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
