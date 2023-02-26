package usecases

import (
	"context"
	"fmt"
	"os"
)

var token = os.Getenv("AUTH_TOKEN")

type AuthenticationMiddlewareUsecase struct{}

func NewAuthenticationMiddlewareUsecase() *AuthenticationMiddlewareUsecase {
	return &AuthenticationMiddlewareUsecase{}
}

type AuthenticateRequest struct {
	Token string `header:"Token" binding:"required"`
}

func (u *AuthenticationMiddlewareUsecase) Authenticate(ctx context.Context, req *AuthenticateRequest) error {
	if req.Token != token {
		return fmt.Errorf("invalid token")
	}

	return nil
}
