package kvsuser

import (
	"authentication/internal/core/gateway"
	"authentication/internal/core/gateway/dto"
	"github.com/gin-gonic/gin"
	"log"
)

func NewAuthenticationUserRepository() gateway.AuthenticationUserRepository {
	return &authenticationUserRepository{}
}

type authenticationUserRepository struct {
}

func (repository *authenticationUserRepository) Execute(ctx *gin.Context, entity *dto.UserAuthData) error {
	log.Printf("User data: \n Email: %s \n Password: %s \n", entity.Email, entity.Password)
	return nil
}
