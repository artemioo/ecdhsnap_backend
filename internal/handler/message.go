package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	ecdhsnap "github.com/artemioo/ecdhsnap_backend"
	"github.com/go-chi/chi/v5"
)

func (h *Handler) CreateMessage(w http.ResponseWriter, r *http.Request) {
	var new_message ecdhsnap.Message
	err := json.NewDecoder(r.Body).Decode(&new_message)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	new_message.Sent_at = time.Now()

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

func (h *Handler) GetRelatedMessages(w http.ResponseWriter, r *http.Request) {
	PairID := chi.URLParam(r, "pairId")
	PairIdInt, err := strconv.Atoi(PairID) // convert to int

	/* 	var Pair ecdhsnap.Pair
	   	err := json.NewDecoder(r.Body).Decode(&P)
	   	if err != nil {
	   		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
	   		return
	   	} */

	id, err := h.services.GetRelatedMessages(PairIdInt)
	if err != nil {
		http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(id))
}
