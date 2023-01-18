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

type authenticationUserRepository struct{}

func (repository *authenticationUserRepository) Execute(userAuthData *dto.UserAuthData) (bool, error) {
	storedUser, err := repository.searchUser(userAuthData)

	if err != nil {
		// TODO: Apierror error al buscar usuario
		return false, err
	}

	return repository.checksumMatches(storedUser, userAuthData), nil
}

func (repository *authenticationUserRepository) searchUser(userData *dto.UserAuthData) (*dto.UserAuthData, error) {
	client, _ := config.GetConnection()

	ctx, cancel := context.WithTimeout(context.Background(), config.TimeOutInterval)
	defer cancel()
	defer config.CloseConnection(ctx, client)

	collection := client.Database(Database).Collection("users")

	var retrievedUser dto.UserAuthData
	err := collection.FindOne(ctx, repository.generateFilterBson(userData)).Decode(&retrievedUser)
	if err != nil {
		return nil, err
	}

	return &retrievedUser, nil
}

func (repository *authenticationUserRepository) checksumMatches(stored, incoming *dto.UserAuthData) bool {
	return incoming.GenerateChecksum() == stored.Checksum
}

func (repository *authenticationUserRepository) generateFilterBson(userData *dto.UserAuthData) bson.M {
	if userData.Email == "" {
		return bson.M{"name": userData.Name}
	}
	return bson.M{"email": userData.Email}
}
