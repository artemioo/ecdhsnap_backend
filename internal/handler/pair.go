package handler

import (
	"encoding/json"
	"net/http"

	ecdhsnap "github.com/artemioo/ecdhsnap_backend"
)

type SessionRequest struct {
	SessionID string `json:"session_id"`
}

func (h *Handler) CreatePair(w http.ResponseWriter, r *http.Request) {
	session, err_s := store.Get(r, "session-id")
	if err_s != nil {
		http.Error(w, "Internal Server Error: "+err_s.Error(), http.StatusInternalServerError)
		return
	}

	UserInitiatorId := session.Values["UserId"].(int)

	var new_pair ecdhsnap.Pair
	err := json.NewDecoder(r.Body).Decode(&new_pair)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	new_pair.Id_user_initiator = UserInitiatorId

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

func (h *Handler) GetRelatedPairs(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session-id")
	UserId := session.Values["UserId"].(int)
	pairs, err := h.services.GetRelatedPairs(UserId)
	if err != nil {
		http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(pairs))
}
