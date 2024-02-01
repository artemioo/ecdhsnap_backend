package handler

import (
	"encoding/json"
	"net/http"
	"time"

	ecdhsnap "github.com/artemioo/ecdhsnap_backend"
)

func (h *Handler) CreateMessage(w http.ResponseWriter, r *http.Request) {
	var new_message ecdhsnap.Message
	err := json.NewDecoder(r.Body).Decode(&new_message)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	new_message.Sent_at = time.Now() // or frontend?

	id, err := h.services.CreateMessage(new_message)
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
