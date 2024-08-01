package app

import (
	"api-golang-codebase/config"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Server struct {
	app *fiber.App
	db  *gorm.DB
}

func NewServer(cfg *config.Config, db *gorm.DB) *Server {
	app := fiber.New()
	// http.
}
