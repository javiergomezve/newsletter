package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"newsletter-back/services"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (c *AuthController) RouteSetup(rg *gin.RouterGroup) {
	rg.POST("/login", c.login)
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (c *AuthController) login(ctx *gin.Context) {
	req := LoginRequest{}
	if errs := validateRequest(ctx, &req); errs != nil {
		abortWithValidationError(ctx, errs)

		return
	}

	token, err := c.authService.Login(req.Email, req.Password)
	if err != nil {
		abortWithValidationError(ctx, []InputError{{
			Field: "email", Message: "Invalid email/password combination",
		}})

		return
	}

	successResponse(ctx, SuccessResponseOptions{
		Message: "auth success",
		Code:    http.StatusOK,
		Data: gin.H{
			"token": token,
		},
	})
}
