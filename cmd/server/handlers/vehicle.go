package handlers

import (
	"challenge-fravega/internal/vehicle"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type VehicleHandler struct {
	service vehicle.Service
}

func (h *VehicleHandler) SetupRoutes(router *gin.Engine) {
	router.GET("/vehicles/:id", h.GetVehicle)
	router.GET("/vehicles", h.GetVehicles)
}

func (h *VehicleHandler) GetVehicle(c *gin.Context) {
	id := c.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	vehicle, err := h.service.GetVehicle(uuid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, vehicle)
}

func (h *VehicleHandler) GetVehicles(c *gin.Context) {
	vehicles, err := h.service.GetVehicles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, vehicles)
}

// static functions

func NewVehicleHandler(service vehicle.Service) *VehicleHandler {
	return &VehicleHandler{service: service}
}
