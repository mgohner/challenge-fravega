package handlers

import (
	"challenge-fravega/internal/route"

	"net/http"

	"github.com/gin-gonic/gin"
)

type RouteHandler struct {
	service route.Service
}

func (h *RouteHandler) SetupRoutes(router *gin.Engine) {
	router.Group("/routes").
		GET("/", h.GetRoutes).
		GET("/:id", h.GetRoute).
		POST("/", h.NewRoute)
}

func (h *RouteHandler) GetRoutes(c *gin.Context) {
	res, err := h.service.GetRoutes()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *RouteHandler) GetRoute(c *gin.Context) {
	res, err := h.service.GetRoute(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *RouteHandler) NewRoute(c *gin.Context) {
	req := &route.CreateRoute{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.service.CreateRoute(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

// static functions

func NewRouteHandler(routeService route.Service) *RouteHandler {
	return &RouteHandler{
		service: routeService,
	}
}
