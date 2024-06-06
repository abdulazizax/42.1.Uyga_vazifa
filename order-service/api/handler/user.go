package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"user/models"
	"user/storage"
)

type Handler struct {
	Storage *storage.User
}

func NewHandeler(s *storage.User) *Handler {
	return &Handler{Storage: s}
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var body models.UserRequest

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := h.Storage.CreateUser(&body)
	if err != nil {
		http.Error(w, "failed creating new user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) UpdateUserById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var body models.UserRequest
	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := h.Storage.UpdateUserById(intId, &body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := h.Storage.DeleteUserByID(intId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
