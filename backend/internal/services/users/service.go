package users

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/juanMaAV92/go-utils/errors"
	"github.com/juanMaAV92/zenith-financial/backend/internal/domain/request"
	"github.com/juanMaAV92/zenith-financial/backend/internal/domain/response"
	"github.com/juanMaAV92/zenith-financial/backend/internal/entities"
	"github.com/juanMaAV92/zenith-financial/backend/utils/crypto"
)

type userRepository interface {
	Create(ctx context.Context, user *entities.User) error
	GetByEmail(ctx context.Context, email string) (*entities.User, error)
}

type service struct {
	userRepository userRepository
}

func NewService(userRepo userRepository) *service {
	return &service{userRepository: userRepo}
}

func (s *service) CreateUser(ctx context.Context, req *request.CreateUser) (*response.User, error) {
	existingUser, err := s.userRepository.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New(http.StatusConflict, "USER_EXISTS", []string{"User already exists"})
	}

	salt, err := crypto.GeneratePasswordSalt()
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, errors.StatusInternalServerErrorCode, []string{"Unable to process request"})
	}

	hashedPassword, err := crypto.HashPassword(req.Password, salt)
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, errors.StatusInternalServerErrorCode, []string{"Unable to process request"})
	}

	newUser := &entities.User{
		Code:         uuid.New(),
		Username:     req.UserName,
		Email:        req.Email,
		PasswordHash: hashedPassword,
		PasswordSalt: salt,
		Currency:     req.Currency,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	err = s.userRepository.Create(ctx, newUser)
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, "CREATE_USER_ERROR", []string{"Unable to create user"})
	}

	return response.ToUserResponse(newUser), nil
}

func (s *service) ValidateCredentials(ctx context.Context, req *request.ValidateUserCredentials) (*response.User, error) {
	var userFound *entities.User
	userFound, err := s.userRepository.GetByEmail(ctx, req.Email)
	if userFound == nil || err != nil {
		return nil, errors.New(http.StatusNotFound, errors.StatusNotFoundCode, []string{"User not found"})
	}

	isValidPassword := crypto.ValidatePassword(req.Password, userFound.PasswordSalt, userFound.PasswordHash)
	if !isValidPassword {
		return nil, errors.New(http.StatusUnauthorized, errors.StatusUnauthorizedCode, []string{"Invalid email or password"})
	}

	return response.ToUserResponse(userFound), nil
}
