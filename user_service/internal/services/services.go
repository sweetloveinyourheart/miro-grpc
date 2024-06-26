package services

import (
	"context"
	"database/sql"

	database "github.com/sweetloveinyourheart/miro-whiteboard/user_service/internal/db"
	"github.com/sweetloveinyourheart/miro-whiteboard/user_service/internal/utils"
)

type UserServices struct {
	db      *sql.DB
	queries *database.Queries
}

type IUserServices interface {
	CreateNewUser(newUser NewUser) (bool, error)
}

type NewUser struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func CreateUserService(db *sql.DB) IUserServices {
	queries := database.New(db)

	return &UserServices{
		db,
		queries,
	}
}

func (sv *UserServices) CreateNewUser(newUser NewUser) (bool, error) {
	ctx := context.Background()

	_, err := sv.queries.GetUserByEmail(ctx, newUser.Email)
	if err == nil {
		return false, nil
	}

	passwordHash, err := utils.HashPassword(newUser.Password)
	if err != nil {
		return false, err
	}

	newUserRecord := database.CreateUserParams{
		Email:     newUser.Email,
		FirstName: utils.ToNullString(newUser.FirstName),
		LastName:  utils.ToNullString(newUser.LastName),
	}
	userId, userErr := sv.queries.CreateUser(ctx, newUserRecord)
	if userErr != nil {
		return false, userErr
	}

	newUserCredential := database.CreateUserCredentialParams{
		UserID:       userId,
		PasswordHash: passwordHash,
	}
	userCredentialErr := sv.queries.CreateUserCredential(ctx, newUserCredential)
	if userCredentialErr != nil {
		return false, userCredentialErr
	}

	return true, nil
}
