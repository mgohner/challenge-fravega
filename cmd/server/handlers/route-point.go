package handlers

import (
	routePoint "challenge-fravega/internal/route-point"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RoutePointHandler struct {
	routePointService routePoint.Service
}

func (h *RoutePointHandler) SetupRoutes(router *gin.Engine) {
	router.Group("/route-points").
		GET("/", h.GetRoutePoints).
		GET("/:id", h.GetRoutePoint).
		POST("/add-purchase-order", h.CreateRoutePoint)
}

func (h *RoutePointHandler) GetRoutePoints(c *gin.Context) {
	res, err := h.routePointService.GetRoutePoints()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *RoutePointHandler) GetRoutePoint(c *gin.Context) {
	res, err := h.routePointService.GetRoutePoint(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *RoutePointHandler) CreateRoutePoint(c *gin.Context) {
	req := &routePoint.AddPurchaseOrder{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.routePointService.CreateRoutePoint(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// static functions

func NewRoutePointHandler(routePointService routePoint.Service) *RoutePointHandler {
	return &RoutePointHandler{routePointService: routePointService}
}
