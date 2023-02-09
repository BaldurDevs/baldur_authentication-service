package user

import (
	"authentication/internal/core/gateway/dto"
	"authentication/internal/core/gateway/mock"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type authenticateUserTestSuite struct {
	suite.Suite

	mockCtrl                       *gomock.Controller
	mockAuthenticateUserRepository *mock.MockAuthenticationUserRepository
}

func TestApplicationAuthenticateUserTestSuite(t *testing.T) {
	suite.Run(t, new(authenticateUserTestSuite))
}

func (st *authenticateUserTestSuite) SetupSuite() {
	mockCtrl := gomock.NewController(st.T())

	st.mockCtrl = mockCtrl
	st.mockAuthenticateUserRepository = mock.NewMockAuthenticationUserRepository(mockCtrl)
}

func (st *authenticateUserTestSuite) TearDownSuite() {
	st.mockCtrl.Finish()
}

func (st *authenticateUserTestSuite) initStandardUseCase() AuthenticateUser {
	st.T().Helper()

	return NewAuthenticateUser(st.mockAuthenticateUserRepository)
}

func (st *authenticateUserTestSuite) Test_AuthenticateUser_Factory() {
	useCase := AuthenticateUserFactory()
	st.NotNil(useCase)
}

func (st *authenticateUserTestSuite) Test_AuthenticateUser_Execute_OK() {
	st.Run("Execute OK, authenticate a user ok", func() {
		mockAuthUserData := dto.MockAuthUserData()

		st.mockAuthenticateUserRepository.EXPECT().Execute(mockAuthUserData).Return(true, nil).Times(1)

		useCase := st.initStandardUseCase()
		result, err := useCase.Execute(mockAuthUserData)

		st.Nil(err)
		st.True(result)
	})

	st.Run("Execute OK, authenticate a user fail", func() {
		mockAuthUserData := dto.MockAuthUserData()

		st.mockAuthenticateUserRepository.EXPECT().Execute(mockAuthUserData).Return(false, nil).Times(1)

		useCase := st.initStandardUseCase()
		result, err := useCase.Execute(mockAuthUserData)

		st.Nil(err)
		st.False(result)
	})
}

func (st *authenticateUserTestSuite) Test_AuthenticateUser_Execute_Error() {
	st.Run("authenticate a user repository returns an error", func() {
		mockAuthUserData := dto.MockAuthUserData()
		expectedError := errors.New("error searching user")

		st.mockAuthenticateUserRepository.EXPECT().Execute(mockAuthUserData).Return(false, expectedError).Times(1)

		useCase := st.initStandardUseCase()
		result, err := useCase.Execute(mockAuthUserData)

		st.False(result)
		st.Equal(expectedError, err)
	})
}
