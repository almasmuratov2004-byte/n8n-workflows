package main

import (
	"net/http"
	"myapp/internal/auth"
	"myapp/internal/di"
	"myapp/internal/handlers"
	"myapp/pkg/logger"
	"github.com/rs/zerolog/log"
)

func main() {
	logger.Init()
	log.Info().Msg("server starting on :8080")
	container, err := di.NewContainer("postgres://postgres:Mirasena2010@localhost:5432/postgres")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to init DI")
	}

	userHandler := handlers.NewUserHandler(container.UserService)

	http.HandleFunc("/auth/register", handlers.RegisterHandler)
	http.HandleFunc("/users", auth.AuthRequired(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			userHandler.GetAll(w, r)
		case "POST":
			userHandler.Create(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	}))

	log.Info().Msg("server started on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
