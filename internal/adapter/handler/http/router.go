package http

import (
	http "github.com/MM16z/mm16-webboard-be-golang/internal/adapter/handler/http/auth"
	"github.com/MM16z/mm16-webboard-be-golang/internal/adapter/storage/postgres"
	"github.com/MM16z/mm16-webboard-be-golang/internal/adapter/storage/postgres/repository"
	service "github.com/MM16z/mm16-webboard-be-golang/internal/core/service/auth"
	"github.com/gofiber/fiber/v2"
)

func InitRouter(app *fiber.App, db *postgres.DB) {
	authRepository := repository.NewUserRepository(db)
	authService := service.NewAuthService(authRepository)
	authHandler := http.NewAuthHandler(authService)

	api := app.Group("/api")
	auth := api.Group("/auth")
	auth.Post("/register", authHandler.Register)
}
