package gateway

import (
	"authentication/internal/core/gateway/dto"
)

//go:generate mockgen -destination=./mock/authentication_user_repository_mock.go -source=./authentication_user_repository.go -package=mock

type AuthenticationUserRepository interface {
	Execute(entity *dto.UserAuthData) (bool, error)
}
