package rest

import (
	"authentication/internal/core/usecase/user"
	"authentication/internal/entrypoints/handler/contracts"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthHandlerFactory() *AuthHandler {
	return NewAuthHandler(
		user.AuthenticateUserFactory(),
	)
}

func NewAuthHandler(authUser user.AuthenticateUser) *AuthHandler {
	return &AuthHandler{
		authenticateUser: authUser,
	}
}

type AuthHandler struct {
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

	err = handler.authenticateUser.Execute(c, authUserRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (handler *AuthHandler) getUserDataFromRequest(c *gin.Context) (*contracts.AuthUserRequest, error) {
	authUserRequest := contracts.AuthUserRequest{}
	err := json.NewDecoder(c.Request.Body).Decode(&authUserRequest)
	if err != nil {
		return nil, err
	}

	return &authUserRequest, nil
}
