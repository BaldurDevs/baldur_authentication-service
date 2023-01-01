package gateway

import (
	"authentication/internal/core/gateway/dto"
	"github.com/gin-gonic/gin"
)

//go:generate mockgen -destination=./mock/authentication_user_repository_mock.go -source=./authentication_user_repository.go -package=mock

type AuthenticationUserRepository interface {
	Execute(ctx *gin.Context, entity *dto.UserAuthData) error
}
