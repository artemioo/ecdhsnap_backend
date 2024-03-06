package handler

import (
	"encoding/json"
	"net/http"

	ecdhsnap "github.com/artemioo/ecdhsnap_backend"
	"github.com/go-chi/chi/v5"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var new_user ecdhsnap.User
	err := json.NewDecoder(r.Body).Decode(&new_user)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	id, err := h.services.CreateUser(new_user)
	if err != nil {
		http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(201)
	result, err := json.Marshal(id)
	if err != nil {
		http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(result)
}

func (h *Handler) GetUserPubKey(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	id, PubKey, err := h.services.GetUserPubKey(username)
	userData := map[string]interface{}{
		"id":     id,
		"pubKey": PubKey,
	}

	w.WriteHeader(200)
	result, err := json.Marshal(userData)
	if err != nil {
		http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(result)
}

func (h *Handler) GetAllUsers(w http.ResponseWriter, r *http.Request) {

	users, err := h.services.GetAllUsers()
	if err != nil {
		http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	/* 	users, err = json.Marshal(users)
	   	if err != nil {
	   		http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
	   		return
	   	} */

	//w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(users))

}
