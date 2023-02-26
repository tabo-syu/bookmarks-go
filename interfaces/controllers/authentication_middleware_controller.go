package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tabo-syu/bookmarks/usecases"
)

type AuthenticationMiddlewareController struct {
	authentication *usecases.AuthenticationMiddlewareUsecase
}

func NewAuthenticationMiddlewareController(authentication *usecases.AuthenticationMiddlewareUsecase) *AuthenticationMiddlewareController {
	return &AuthenticationMiddlewareController{authentication}
}

func (c *AuthenticationMiddlewareController) Authenticate(g *gin.Context) {
	var req usecases.AuthenticateRequest
	if err := g.ShouldBindHeader(&req); err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	if err := c.authentication.Authenticate(g, &req); err != nil {
		g.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})

		return
	}
}
