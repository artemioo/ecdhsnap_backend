package handler

import (
	"net/http"

	"github.com/artemioo/ecdhsnap_backend/internal/service"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	services *service.Service
}

// конструктор для хендлеров
func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
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
