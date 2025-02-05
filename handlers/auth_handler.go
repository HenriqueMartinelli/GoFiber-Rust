package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func LoginCert(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Login ainda não implementado"})

	// var request struct {
	// 	ID int `json:"id"`
	// }
	// if err := c.BodyParser(&request); err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Dados incorretos"})
	// }

	// log.Printf("Tentativa de login com cert ID: %d", request.ID)
	// response, err := services.LoginWithCert(request.ID)
	// if err != nil {
	// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	// }
	// return c.JSON({response})
	// r := "ok"
	// return r
}
