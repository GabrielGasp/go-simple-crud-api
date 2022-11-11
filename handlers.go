package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateAirplane(c *fiber.Ctx) error {
	var newAirplane airplane

	if err := c.BodyParser(&newAirplane); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	newAirplane.ID = uuid.New().String()

	airplanes, err := ReadAirplanesFromFile("airplanes.json")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	airplanes = append(airplanes, newAirplane)

	err = WriteAirplanesToFile("airplanes.json", airplanes)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(newAirplane)
}

func GetAirplanes(c *fiber.Ctx) error {
	airplanes, err := ReadAirplanesFromFile("airplanes.json")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(airplanes)
}

func GetAirplaneById(c *fiber.Ctx) error {
	airplanes, err := ReadAirplanesFromFile("airplanes.json")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	id := c.Params("id")

	_, airplane, found := FindAirplaneById(airplanes, id)
	if !found {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Airplane not found"})
	}

	return c.Status(http.StatusOK).JSON(airplane)
}

func UpdateAirplane(c *fiber.Ctx) error {
	airplanes, err := ReadAirplanesFromFile("airplanes.json")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	id := c.Params("id")

	index, _, found := FindAirplaneById(airplanes, id)
	if !found {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Airplane not found"})
	}

	var updatedAirplane airplane

	if err := c.BodyParser(&updatedAirplane); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	updatedAirplane.ID = id

	airplanes[index] = updatedAirplane

	err = WriteAirplanesToFile("airplanes.json", airplanes)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(updatedAirplane)
}

func DeleteAirplane(c *fiber.Ctx) error {
	airplanes, err := ReadAirplanesFromFile("airplanes.json")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	id := c.Params("id")

	index, _, found := FindAirplaneById(airplanes, id)

	if !found {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Airplane not found"})
	}

	airplanes = append(airplanes[:index], airplanes[index+1:]...)

	err = WriteAirplanesToFile("airplanes.json", airplanes)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Airplane deleted"})
}
