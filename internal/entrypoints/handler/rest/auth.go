package rest

import (
	"authentication/internal/core/gateway/dto"
	"encoding/json"
	"net/http"

	"authentication/internal/core/usecase/user"
	"authentication/internal/entrypoints/handler/contracts"
	"github.com/BaldurDevs/baldur_go-library/pkg/goutils/mapper"
	"github.com/gin-gonic/gin"
)

func AuthHandlerFactory() *AuthHandler {
	return NewAuthHandler(
		user.AuthenticateUserFactory(),
		mapper.MapperFactory(),
	)
}

func NewAuthHandler(authUser user.AuthenticateUser, maps mapper.Mapper) *AuthHandler {
	return &AuthHandler{
		authenticateUser: authUser,
		mapper:           maps,
	}
}

type AuthHandler struct {
	authenticateUser user.AuthenticateUser
	mapper           mapper.Mapper
}

func (handler *AuthHandler) RegisterRoutes(r *gin.Engine) {
	r.POST(basePath+"/authenticate", handler.authenticate)
}

func (handler *AuthHandler) authenticate(c *gin.Context) {
	authRequest, err := handler.getUserDataFromRequest(c)
	if err != nil {
		// TODO: errorHandler
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	var authUserRequest dto.UserAuthData
	err = handler.mapper.Map(authRequest, &authUserRequest)
	if err != nil {
		// TODO: errorHandler
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	result, err := handler.authenticateUser.Execute(&authUserRequest)
	if err != nil || !result {
		c.JSON(http.StatusUnauthorized, map[string]any{"status": "error", "message": "wrong user data"})
		return
	}

	c.JSON(http.StatusOK, map[string]any{"status": "ok", "message": "user authenticated successfully"})
}

func (handler *AuthHandler) getUserDataFromRequest(c *gin.Context) (*contracts.AuthUserRequest, error) {
	authUserRequest := contracts.AuthUserRequest{}
	err := json.NewDecoder(c.Request.Body).Decode(&authUserRequest)
	if err != nil {
		return nil, err
	}

	return &authUserRequest, nil
}
