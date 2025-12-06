package handler

import (
	"api-gateway/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct{}

func NewAdminHandler() *AdminHandler {
	return &AdminHandler{}
}

// GetStats godoc
// @Summary      Get system statistics
// @Description  Get detailed statistics about API usage (Admin only)
// @Tags         Admin
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Security     ApiKeyAuth
// @Success      200 {object} AdminStatsResponse
// @Failure      401 {object} ErrorResponse
// @Failure      403 {object} ErrorResponse
// @Router       /api/v1/admin/stats [get]
func (h *AdminHandler) GetStats(c *gin.Context) {
	response.Success(c, http.StatusOK, "Admin stats", AdminStatsData{
		TotalRequests: 1000,
		ActiveUsers:   50,
	})
}

// Models
type AdminStatsData struct {
	TotalRequests int `json:"total_requests" example:"1000"`
	ActiveUsers   int `json:"active_users" example:"50"`
}

type AdminStatsResponse struct {
	Success bool           `json:"success" example:"true"`
	Message string         `json:"message" example:"Admin stats"`
	Data    AdminStatsData `json:"data"`
}
