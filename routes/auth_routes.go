package routes

import (
	"github.com/HenriqueMartinelli/GoFiber-Rust/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(app *fiber.App) {
	// authGroup := app.Group("/login", middlewares.AuthMiddleware)
	authGroup := app.Group("/login")

	authGroup.Post("/cert", handlers.LoginCert)
	// authGroup.Post("/cert/system", handlers.LoginCertSystem)
	// authGroup.Post("/pass", handlers.LoginPass)
}
