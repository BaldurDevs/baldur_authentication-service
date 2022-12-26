package user

import (
	"authentication/internal/core/gateway/dto"
	"github.com/gin-gonic/gin"
)

type AuthenticateUser interface {
	Execute(ctx *gin.Context, user *dto.UserAuthData) (*dto.UserAuthData, error)
}

func AuthenticateUserFactory() AuthenticateUser {
	return NewAuthenticateUser()
}

func NewAuthenticateUser() AuthenticateUser {
	return &authenticateUser{}
}

type authenticateUser struct {
}

func (useCase *authenticateUser) Execute(ctx *gin.Context, user *dto.UserAuthData) (*dto.UserAuthData, error) {
	return nil, nil
}
