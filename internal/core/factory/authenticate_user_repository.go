package factory

import (
	"authentication/internal/core/gateway"
	"authentication/internal/infra/gateway/kvs/kvsuser"
)

func AuthenticateUserRepositoryFactory() gateway.AuthenticationUserRepository {
	return kvsuser.NewAuthenticationUserRepository()
}
