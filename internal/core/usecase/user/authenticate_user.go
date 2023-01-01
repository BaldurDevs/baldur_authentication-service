package user

import (
	"authentication/internal/core/factory"
	"authentication/internal/core/gateway"
	"authentication/internal/core/gateway/dto"
	"authentication/internal/entrypoints/handler/contracts"
	"github.com/gin-gonic/gin"
	"log"
)

type AuthenticateUser interface {
	Execute(ctx *gin.Context, user *contracts.AuthUserRequest) error
}

func AuthenticateUserFactory() AuthenticateUser {
	return NewAuthenticateUser(
		factory.AuthenticateUserRepositoryFactory(),
	)
}

func NewAuthenticateUser(
	authRepository gateway.AuthenticationUserRepository,
) AuthenticateUser {
	return &authenticateUser{
		authenticationRepository: authRepository,
	}
}

type authenticateUser struct {
	authenticationRepository gateway.AuthenticationUserRepository
}

func (useCase *authenticateUser) Execute(ctx *gin.Context, user *contracts.AuthUserRequest) error {
	// TODO: Correct this map
	userData := &dto.UserAuthData{
		Email:    user.Email,
		Password: user.Password,
	}

	err := useCase.authenticationRepository.Execute(ctx, userData)
	if err != nil {
		return err
	}

	log.Println("Authentication successfully")
	return nil
}
