// internal/handler/auth_handler.go
package handler

import (
	"api-gateway/pkg/jwt"
	"api-gateway/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	jwtSecret string
}

func NewAuthHandler(jwtSecret string) *AuthHandler {
	return &AuthHandler{
		jwtSecret: jwtSecret,
	}
}

// Login godoc
// @Summary      User login
// @Description  Authenticate user and get JWT token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request body LoginRequest true "Login credentials"
// @Success      200 {object} LoginResponse
// @Failure      400 {object} ErrorResponse
// @Failure      401 {object} ErrorResponse
// @Router       /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request")
		return
	}

	if req.Username != "admin" || req.Password != "password" {
		response.Error(c, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	token, err := jwt.GenerateToken("user-123", req.Username, "admin", h.jwtSecret)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	response.Success(c, http.StatusOK, "Login successful", LoginData{
		Token: token,
		Type:  "Bearer",
	})
}

// Models
type LoginRequest struct {
	Username string `json:"username" example:"admin" binding:"required"`
	Password string `json:"password" example:"password" binding:"required"`
}

type LoginData struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	Type  string `json:"type" example:"Bearer"`
}

type LoginResponse struct {
	Success bool      `json:"success" example:"true"`
	Message string    `json:"message" example:"Login successful"`
	Data    LoginData `json:"data"`
}

type ErrorResponse struct {
	Success bool   `json:"success" example:"false"`
	Error   string `json:"error" example:"Error message"`
}
