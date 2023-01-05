package user

import (
	"authentication/internal/core/factory"
	"authentication/internal/core/gateway"
	"authentication/internal/core/gateway/dto"
	"authentication/internal/entrypoints/handler/contracts"
	"github.com/gin-gonic/gin"
)

type AuthenticateUser interface {
	Execute(ctx *gin.Context, user *contracts.AuthUserRequest) (bool, error)
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

func (useCase *authenticateUser) Execute(ctx *gin.Context, user *contracts.AuthUserRequest) (bool, error) {
	// TODO: Correct this map
	userData := &dto.UserAuthData{
		Email:    user.Email,
		Password: user.Password,
	}

	result, err := useCase.authenticationRepository.Execute(userData)
	if err != nil {
		return false, err
	}

	return result, nil
}
