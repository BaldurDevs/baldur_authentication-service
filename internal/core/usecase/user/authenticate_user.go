package user

import (
	"authentication/internal/core/factory"
	"authentication/internal/core/gateway"
	"authentication/internal/core/gateway/dto"
	"errors"
)

//go:generate mockgen -destination=./authenticate_user_mock.go -source=./authenticate_user.go -package=user

type AuthenticateUser interface {
	Execute(user *dto.UserAuthData) (bool, error)
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

func (useCase *authenticateUser) Execute(user *dto.UserAuthData) (bool, error) {
	if err := useCase.validateUserData(user); err != nil {
		return false, err
	}

	result, err := useCase.authenticationRepository.Execute(user)
	if err != nil {
		return false, err
	}

	return result, nil
}

func (useCase *authenticateUser) validateUserData(user *dto.UserAuthData) error {
	if user.Email == "" && user.Name == "" {
		return errors.New("user email and name empty")
	}
	return nil
}
