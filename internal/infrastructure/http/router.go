package http

import (
	"api-golang-codebase/internal/infrastructure/http/handler"
	"api-golang-codebase/internal/infrastructure/presistence/postgres"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	userRepo := postgres.NewUserRepository(db)
	userHandler := handler.NewUserHandler(userRepo)
	app.Get("/users/:id", userHandler.GetUser)
}
