package handler

import (
	"encoding/json"
	"net/http"

	ecdhsnap "github.com/artemioo/ecdhsnap_backend"
)

func (h *Handler) WelcomePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test"))
}

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
	var user_id ecdhsnap.User
	err := json.NewDecoder(r.Body).Decode(&user_id)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	user, err := h.services.GetUserPubKey(user_id.Id)

	w.WriteHeader(201)
	result, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(result)
}
