package handler

import (
	"net/http"
)

func (h *Handler) WelcomePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test"))
}
