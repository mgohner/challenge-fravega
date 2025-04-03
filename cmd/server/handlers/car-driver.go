package handlers

import (
	carDriver "challenge-fravega/internal/car-driver"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CarDriverHandler struct {
	service carDriver.Service
}

func (h *CarDriverHandler) SetupRoutes(router *gin.Engine) {
	router.GET("/car-drivers/:id", h.GetCarDriver)
	router.GET("/car-drivers", h.GetCarDrivers)
}

func (h *CarDriverHandler) GetCarDriver(c *gin.Context) {
	id := c.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	driver, err := h.service.GetDriver(uuid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, driver)
}

func (h *CarDriverHandler) GetCarDrivers(c *gin.Context) {
	drivers, err := h.service.GetDrivers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, drivers)
}

// static functions

func NewCarDriverHandler(service carDriver.Service) *CarDriverHandler {
	return &CarDriverHandler{service: service}
}
