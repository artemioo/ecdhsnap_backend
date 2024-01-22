package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
}

func (h *Handler) InitRoutes() http.Handler {
	router := chi.NewRouter()

	router.Get("/", h.WelcomePage)

	/* 	router.Route("/user", func(r chi.Router) {
		r.Get("/name", getUserName)
		r.Get("/adress", getUserAdress)
		r.Get("/pubkey", getUserPubKey)
	}) */
	return router
}
