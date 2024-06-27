package repository

import (
	"context"
	"database/sql"

	database "github.com/sweetloveinyourheart/miro-whiteboard/user_service/internal/db"
)

type UserRepository interface {
	CreateUser(ctx context.Context, arg database.CreateUserParams) (int32, error)
	CreateUserCredential(ctx context.Context, arg database.CreateUserCredentialParams) error
	GetUser(ctx context.Context, userID int32) (database.User, error)
	GetUserByEmail(ctx context.Context, email string) (database.User, error)
	GetUserInfoWithCredentials(ctx context.Context, email string) (database.GetUserInfoWithCredentialsRow, error)
}

func NewUserRepository(db *sql.DB) UserRepository {
	repository := database.New(db)
	return repository
}
