package handler

import (
	"net/http"

	"github.com/artemioo/ecdhsnap_backend/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
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

	// Basic CORS
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.Route("/user", func(r chi.Router) {
		r.Post("/create", h.CreateUser)
		r.Get("/", h.GetUserPubKey)
	})
	router.Route("/pair", func(r chi.Router) {
		r.Post("/create", h.CreatePair)
		r.Get("/related/{userId}", h.GetRelatedPairs)
	})
	router.Route("/message", func(r chi.Router) {
		r.Post("/create", h.CreateMessage)
		r.Get("/related/{pairId}", h.GetRelatedMessages)
	})
	return router
}
