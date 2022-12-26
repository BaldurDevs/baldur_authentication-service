package rest

import (
	"authentication/internal/core/gateway/dto"
	"authentication/internal/core/usecase/user"
	"authentication/internal/entrypoints/handler/contracts"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthHandlerFactory() AuthHandler {
	return NewAuthHandler()
}

func NewAuthHandler() AuthHandler {
	return AuthHandler{
		authenticateUser: user.AuthenticateUserFactory(),
	}
}

type AuthHandler struct {
	autoMapper       automapper.AutoMapper
	authenticateUser user.AuthenticateUser
}

func (handler *AuthHandler) RegisterRouters(r *gin.Engine) {
	r.POST(basePath+"/authenticate", handler.authenticate)
}

func (handler *AuthHandler) authenticate(c *gin.Context) {

	authUserRequest, err := handler.getUserDataFromRequest(c)
	if err != nil {
		// TODO: errorHandler
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	var userData dto.UserAuthData

	// TODO: Mappear authUserRequest --> userData

	authResult, err := handler.authenticateUser.Execute(c, &userData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, authResult)
}

func (handler *AuthHandler) getUserDataFromRequest(c *gin.Context) (*contracts.AuthUserRequest, error) {
	authUserRequest := contracts.AuthUserRequest{}
	err := json.NewDecoder(c.Request.Body).Decode(&authUserRequest)
	if err != nil {
		return nil, err
	}

	return &authUserRequest, nil
}
