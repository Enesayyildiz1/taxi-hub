package handler

import (
	"driver-service/internal/dto"
	"driver-service/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DriverHandler struct {
	service service.DriverService
}

func NewDriverHandler(service service.DriverService) *DriverHandler {
	return &DriverHandler{
		service: service,
	}
}

func (h *DriverHandler) CreateDriver(c *gin.Context) {
	var req dto.CreateDriverRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request",
			"details": err.Error(),
		})
		return
	}

	response, err := h.service.CreateDriver(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create driver",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Driver created successfully",
		"data":    response,
	})
}

func (h *DriverHandler) GetDriver(c *gin.Context) {
	id := c.Param("id")

	response, err := h.service.GetDriver(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Driver not found",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": response,
	})
}

func (h *DriverHandler) GetDrivers(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("pageSize", "20")

	page, _ := strconv.Atoi(pageStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)

	response, err := h.service.GetDrivers(c.Request.Context(), page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to get drivers",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *DriverHandler) UpdateDriver(c *gin.Context) {
	id := c.Param("id")

	var req dto.UpdateDriverRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request",
			"details": err.Error(),
		})
		return
	}

	if err := h.service.UpdateDriver(c.Request.Context(), id, &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to update driver",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Driver updated successfully",
	})
}

func (h *DriverHandler) DeleteDriver(c *gin.Context) {
	id := c.Param("id")

	if err := h.service.DeleteDriver(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to delete driver",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Driver deleted successfully",
	})
}
