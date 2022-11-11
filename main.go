package main

import "github.com/gofiber/fiber/v2"

type airplane struct {
	ID           string `json:"id"`
	Manufacturer string `json:"manufacturer"`
	Model        string `json:"model"`
	Seats        int    `json:"seats"`
}

// Main
func main() {
	app := fiber.New()

	app.Post("/airplanes", CreateAirplane)
	app.Get("/airplanes", GetAirplanes)
	app.Get("/airplanes/:id", GetAirplaneById)
	app.Put("/airplanes/:id", UpdateAirplane)
	app.Delete("/airplanes/:id", DeleteAirplane)

	app.Listen(":3300")
}
