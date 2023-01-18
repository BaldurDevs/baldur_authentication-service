package user

import (
	"authentication/internal/core/factory"
	"authentication/internal/core/gateway"
	"authentication/internal/core/gateway/dto"
	"authentication/internal/entrypoints/handler/contracts"
	"errors"
)

type AuthenticateUser interface {
	Execute(user *contracts.AuthUserRequest) (bool, error)
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

func (useCase *authenticateUser) Execute(user *contracts.AuthUserRequest) (bool, error) {
	// TODO: Correct this map
	userData := &dto.UserAuthData{
		Email:    user.Email,
		Name:     user.Name,
		Password: user.Password,
	}

	if err := useCase.validateUserData(user); err != nil {
		return false, err
	}

	result, err := useCase.authenticationRepository.Execute(userData)
	if err != nil {
		return false, err
	}

	return result, nil
}

func (useCase *authenticateUser) validateUserData(user *contracts.AuthUserRequest) error {
	if user.Email == "" && user.Name == "" {
		return errors.New("user email and name empty")
	}
	return nil
}
