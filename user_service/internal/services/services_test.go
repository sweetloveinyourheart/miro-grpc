package services

import (
	"context"
	"database/sql"

	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/sweetloveinyourheart/miro-whiteboard/user_service/internal/db"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) GetUserByEmail(ctx context.Context, email string) (db.User, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(db.User), args.Error(1)
}

func (m *MockRepository) CreateUser(ctx context.Context, params db.CreateUserParams) (int32, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(int32), args.Error(1)
}

func (m *MockRepository) GetUser(ctx context.Context, userId int32) (db.User, error) {
	args := m.Called(ctx, userId)
	return args.Get(0).(db.User), args.Error(1)
}

func (m *MockRepository) CreateUserCredential(ctx context.Context, params db.CreateUserCredentialParams) error {
	args := m.Called(ctx, params)
	return args.Error(0)
}

func (m *MockRepository) GetUserInfoWithCredentials(ctx context.Context, email string) (db.GetUserInfoWithCredentialsRow, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(db.GetUserInfoWithCredentialsRow), args.Error(1)
}

func (m *MockRepository) SetupTest(t *testing.T) {
	m.ExpectedCalls = nil
	m.Calls = nil
}

func TestCreateNewUser(t *testing.T) {
	mockRepository := new(MockRepository)
	sv := &UserServices{
		db:         &sql.DB{}, // Dummy DB, can be nil
		repository: mockRepository,
	}

	t.Run("User already exists", func(t *testing.T) {
		mockRepository.SetupTest(t)

		mockRepository.On("GetUserByEmail", mock.Anything, "existing@example.com").Return(db.User{}, nil)

		result, err := sv.CreateNewUser(User{Email: "existing@example.com", Password: "password"})

		assert.Nil(t, err)
		assert.False(t, result)
		mockRepository.AssertExpectations(t)
	})

	t.Run("Create user fails", func(t *testing.T) {
		mockRepository.SetupTest(t)

		mockRepository.On("GetUserByEmail", mock.Anything, "new@example.com").Return(db.User{}, errors.New("user not found"))
		mockRepository.On("CreateUser", mock.Anything, mock.Anything).Return(int32(0), errors.New("create user error"))

		result, err := sv.CreateNewUser(User{Email: "new@example.com", Password: "password"})

		assert.NotNil(t, err)
		assert.False(t, result)
		mockRepository.AssertExpectations(t)
	})

	t.Run("Create user credential fails", func(t *testing.T) {
		mockRepository.SetupTest(t)

		mockRepository.On("GetUserByEmail", mock.Anything, "new@example.com").Return(db.User{}, errors.New("user not found"))
		mockRepository.On("CreateUser", mock.Anything, mock.Anything).Return(int32(1), nil)
		mockRepository.On("CreateUserCredential", mock.Anything, mock.Anything).Return(errors.New("create user credential error"))

		result, err := sv.CreateNewUser(User{Email: "new@example.com", Password: "password"})

		assert.NotNil(t, err)
		assert.False(t, result)
		mockRepository.AssertExpectations(t)
	})

	t.Run("Successful user creation", func(t *testing.T) {
		mockRepository.SetupTest(t)

		mockRepository.On("GetUserByEmail", mock.Anything, "new@example.com").Return(db.User{}, errors.New("user not found"))
		mockRepository.On("CreateUser", mock.Anything, mock.Anything).Return(int32(1), nil)
		mockRepository.On("CreateUserCredential", mock.Anything, mock.Anything).Return(nil)

		result, err := sv.CreateNewUser(User{Email: "new@example.com", Password: "password"})

		assert.Nil(t, err)
		assert.True(t, result)
		mockRepository.AssertExpectations(t)
	})
}
