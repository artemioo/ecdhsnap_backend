package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	ecdhsnap "github.com/artemioo/ecdhsnap_backend"
)

func (h *Handler) WelcomePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test"))
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var new_user ecdhsnap.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&new_user)
	if err != nil {
		fmt.Println("JSON Decoding Error:", err)
		panic(err)
	}

	id, err := h.services.CreateUser(new_user)
	if err != nil {
		//вернуть ошибку + поменять статускод
		http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(201)
	result, err := json.Marshal(id)
	w.Write(result)
}
