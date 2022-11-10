package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateAirplane(c *gin.Context) {
	var newAirplane airplane

	if err := c.ShouldBindJSON(&newAirplane); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newAirplane.ID = uuid.New().String()

	airplanes, err := ReadAirplanesFromFile("airplanes.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	airplanes = append(airplanes, newAirplane)

	err = WriteAirplanesToFile("airplanes.json", airplanes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newAirplane)
}

func GetAirplanes(c *gin.Context) {
	airplanes, err := ReadAirplanesFromFile("airplanes.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, airplanes)
}

func GetAirplaneById(c *gin.Context) {
	airplanes, err := ReadAirplanesFromFile("airplanes.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")

	_, airplane, found := FindAirplaneById(airplanes, id)
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Airplane not found"})
		return
	}

	c.JSON(http.StatusOK, airplane)
}

func UpdateAirplane(c *gin.Context) {
	airplanes, err := ReadAirplanesFromFile("airplanes.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")

	index, _, found := FindAirplaneById(airplanes, id)
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Airplane not found"})
		return
	}

	var updatedAirplane airplane

	if err := c.ShouldBindJSON(&updatedAirplane); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedAirplane.ID = id

	airplanes[index] = updatedAirplane

	err = WriteAirplanesToFile("airplanes.json", airplanes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedAirplane)
}

func DeleteAirplane(c *gin.Context) {
	airplanes, err := ReadAirplanesFromFile("airplanes.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")

	index, _, found := FindAirplaneById(airplanes, id)

	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Airplane not found"})
		return
	}

	airplanes = append(airplanes[:index], airplanes[index+1:]...)

	err = WriteAirplanesToFile("airplanes.json", airplanes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
