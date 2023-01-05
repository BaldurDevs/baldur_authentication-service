package kvsuser

import (
	"authentication/internal/core/gateway"
	"authentication/internal/core/gateway/dto"
	"authentication/internal/infra/gateway/kvs/config"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

func NewAuthenticationUserRepository() gateway.AuthenticationUserRepository {
	return &authenticationUserRepository{}
}

type authenticationUserRepository struct {
}

func (repository *authenticationUserRepository) Execute(userAuthData *dto.UserAuthData) (bool, error) {

	entity, err := repository.retrieveUser(userAuthData.Email)
	if err != nil {
		return false, err
	}

	result, err := repository.PasswordMatches(entity.Password, userAuthData.Password)
	if err != nil {
		return false, err
	}

	return result, nil
}

func (repository *authenticationUserRepository) retrieveUser(email string) (*dto.UserAuthData, error) {
	client, err := config.GetConnection()

	ctx, cancel := context.WithTimeout(context.Background(), config.TimeOutInterval)
	defer cancel()
	defer config.CloseConnection(ctx, client)

	collection := client.Database(Database).Collection("users")

	var userData dto.UserAuthData
	err = collection.FindOne(ctx, bson.M{"email": email}).Decode(&userData)
	if err != nil {
		return nil, err
	}

	return &userData, nil
}

func (repository *authenticationUserRepository) PasswordMatches(storagePassword string, receivedPassword string) (bool, error) {
	// TODO: Codificar la contraseña en base64
	return storagePassword == receivedPassword, nil
}
