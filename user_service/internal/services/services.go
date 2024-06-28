package services

import (
	"context"
	"database/sql"
	"errors"
	"time"

	database "github.com/sweetloveinyourheart/miro-whiteboard/user_service/internal/db"
	"github.com/sweetloveinyourheart/miro-whiteboard/user_service/internal/repository"
	"github.com/sweetloveinyourheart/miro-whiteboard/user_service/internal/utils"
)

type UserServices struct {
	db         *sql.DB
	repository repository.UserRepository
}

type IUserServices interface {
	CreateNewUser(newUser NewUser) (bool, error)
	GetAuthCredentials(user UserCredential) (AuthenticationCredential, error)
}

func CreateUserService(db *sql.DB) IUserServices {
	repository := repository.NewUserRepository(db)

	return &UserServices{
		db,
		repository,
	}
}

type NewUser struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type UserCredential struct {
	Email    string
	Password string
}

type AuthenticationCredential struct {
	AccessToken  string
	RefreshToken string
}

func (sv *UserServices) CreateNewUser(newUser NewUser) (bool, error) {
	ctx := context.Background()

	_, err := sv.repository.GetUserByEmail(ctx, newUser.Email)
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
	userId, userErr := sv.repository.CreateUser(ctx, newUserRecord)
	if userErr != nil {
		return false, userErr
	}

	newUserCredential := database.CreateUserCredentialParams{
		UserID:       userId,
		PasswordHash: passwordHash,
	}
	userCredentialErr := sv.repository.CreateUserCredential(ctx, newUserCredential)
	if userCredentialErr != nil {
		return false, userCredentialErr
	}

	return true, nil
}

func (sv *UserServices) GetAuthCredentials(user UserCredential) (AuthenticationCredential, error) {
	ctx := context.Background()

	userInfo, err := sv.repository.GetUserInfoWithCredentials(ctx, user.Email)
	if err != nil {
		return AuthenticationCredential{}, errors.New("email or password is not valid")
	}

	isValidPassword := utils.CheckPasswordHash(user.Password, userInfo.Pwd)
	if !isValidPassword {
		return AuthenticationCredential{}, errors.New("email or password is not valid")
	}

	accessToken, err := utils.GenerateToken(user.Email, 15*time.Minute)
	if err != nil {
		return AuthenticationCredential{}, errors.New("failed to generate access token")
	}

	refreshToken, err := utils.GenerateToken(user.Email, 7*24*time.Hour) // Refresh token valid for 7 days
	if err != nil {
		return AuthenticationCredential{}, errors.New("failed to generate refresh token")
	}

	return AuthenticationCredential{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
