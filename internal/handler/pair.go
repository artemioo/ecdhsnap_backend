package handler

import (
	"encoding/json"
	"net/http"

	ecdhsnap "github.com/artemioo/ecdhsnap_backend"
)

func (h *Handler) CreatePair(w http.ResponseWriter, r *http.Request) {
	var new_pair ecdhsnap.Pair
	err := json.NewDecoder(r.Body).Decode(&new_pair)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}
	id, err := h.services.CreatePair(new_pair)
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
