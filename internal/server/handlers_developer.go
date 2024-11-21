package server

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (s *Server) createDeveloperHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	id, err := s.db.CreateDeveloper(req.Name, req.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := map[string]string{"id": id}
	jsonResp, _ := json.Marshal(resp)
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(jsonResp)
}

func (s *Server) getDeveloperHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	developer, err := s.db.GetDeveloper(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResp, _ := json.Marshal(developer)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsonResp)
}

func (s *Server) updateDeveloperHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var req struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := s.db.UpdateDeveloper(id, req.Name, req.Email); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) deleteDeveloperHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if err := s.db.DeleteDeveloper(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
