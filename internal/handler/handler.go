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
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Set-Cookie"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	router.Post("/login", h.login)
	router.Post("/logout", h.logout)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})

	router.Route("/user", func(r chi.Router) {
		r.Post("/create", h.CreateUser)
		r.Get("/{username}", h.GetUserPubKey)
	})

	router.Get("/users", h.GetAllUsers)

	router.Route("/pair", func(r chi.Router) {
		r.Post("/create", h.CreatePair)
		r.Get("/related", h.GetRelatedPairs)
	})
	router.Route("/message", func(r chi.Router) {
		r.Post("/create", h.CreateMessage)
		r.Get("/related/{pairId}", h.GetRelatedMessages)
	})
	return router
}
