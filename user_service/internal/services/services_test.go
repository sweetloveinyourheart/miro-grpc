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

// MockQueries struct to mock database.Queries
type MockQueries struct {
	mock.Mock
}

func (m *MockQueries) GetUserByEmail(ctx context.Context, email string) (db.User, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(db.User), args.Error(1)
}

func (m *MockQueries) CreateUser(ctx context.Context, params db.CreateUserParams) (int32, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(int32), args.Error(1)
}

func (m *MockQueries) GetUser(ctx context.Context, userId int32) (db.User, error) {
	args := m.Called(ctx, userId)
	return args.Get(0).(db.User), args.Error(1)
}

func (m *MockQueries) CreateUserCredential(ctx context.Context, params db.CreateUserCredentialParams) error {
	args := m.Called(ctx, params)
	return args.Error(0)
}

func (m *MockQueries) SetupTest(t *testing.T) {
	m.ExpectedCalls = nil
	m.Calls = nil
}

func TestCreateNewUser(t *testing.T) {
	mockQueries := new(MockQueries)
	sv := &UserServices{
		db:      &sql.DB{}, // Dummy DB, can be nil
		queries: mockQueries,
	}

	t.Run("User already exists", func(t *testing.T) {
		mockQueries.SetupTest(t)

		mockQueries.On("GetUserByEmail", mock.Anything, "existing@example.com").Return(db.User{}, nil)

		result, err := sv.CreateNewUser(NewUser{Email: "existing@example.com", Password: "password"})

		assert.Nil(t, err)
		assert.False(t, result)
		mockQueries.AssertExpectations(t)
	})

	t.Run("Create user fails", func(t *testing.T) {
		mockQueries.SetupTest(t)

		mockQueries.On("GetUserByEmail", mock.Anything, "new@example.com").Return(db.User{}, errors.New("user not found"))
		mockQueries.On("CreateUser", mock.Anything, mock.Anything).Return(int32(0), errors.New("create user error"))

		result, err := sv.CreateNewUser(NewUser{Email: "new@example.com", Password: "password"})

		assert.NotNil(t, err)
		assert.False(t, result)
		mockQueries.AssertExpectations(t)
	})

	t.Run("Create user credential fails", func(t *testing.T) {
		mockQueries.SetupTest(t)

		mockQueries.On("GetUserByEmail", mock.Anything, "new@example.com").Return(db.User{}, errors.New("user not found"))
		mockQueries.On("CreateUser", mock.Anything, mock.Anything).Return(int32(1), nil)
		mockQueries.On("CreateUserCredential", mock.Anything, mock.Anything).Return(errors.New("create user credential error"))

		result, err := sv.CreateNewUser(NewUser{Email: "new@example.com", Password: "password"})

		assert.NotNil(t, err)
		assert.False(t, result)
		mockQueries.AssertExpectations(t)
	})

	t.Run("Successful user creation", func(t *testing.T) {
		mockQueries.SetupTest(t)

		mockQueries.On("GetUserByEmail", mock.Anything, "new@example.com").Return(db.User{}, errors.New("user not found"))
		mockQueries.On("CreateUser", mock.Anything, mock.Anything).Return(int32(1), nil)
		mockQueries.On("CreateUserCredential", mock.Anything, mock.Anything).Return(nil)

		result, err := sv.CreateNewUser(NewUser{Email: "new@example.com", Password: "password"})

		assert.Nil(t, err)
		assert.True(t, result)
		mockQueries.AssertExpectations(t)
	})
}
