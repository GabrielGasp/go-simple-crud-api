package main

import (
	"github.com/gin-gonic/gin"
)

type airplane struct {
	ID           string `json:"id"`
	Manufacturer string `json:"manufacturer"`
	Model        string `json:"model"`
	Seats        int    `json:"seats"`
}

// Main
func main() {
	router := gin.Default()

	router.POST("/airplanes", CreateAirplane)
	router.GET("/airplanes", GetAirplanes)
	router.GET("/airplanes/:id", GetAirplaneById)
	router.PUT("/airplanes/:id", UpdateAirplane)
	router.DELETE("/airplanes/:id", DeleteAirplane)

	router.Run(":3300")
}
