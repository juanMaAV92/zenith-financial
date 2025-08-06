package auth

import (
	"context"
	"net/http"

	"github.com/juanMaAV92/go-utils/errors"
	"github.com/juanMaAV92/zenith-financial/backend/internal/domain/request"
	"github.com/juanMaAV92/zenith-financial/backend/internal/domain/response"
	"github.com/juanMaAV92/zenith-financial/backend/internal/entities"
	"github.com/juanMaAV92/zenith-financial/backend/utils/crypto"
	"github.com/juanMaAV92/zenith-financial/backend/utils/jwt"
)

type userRepository interface {
	GetByEmail(ctx context.Context, email string) (*entities.User, error)
}

type service struct {
	userRepository userRepository
}

func NewService(userRepo userRepository) *service {
	return &service{userRepository: userRepo}
}

func (s *service) Login(ctx context.Context, req *request.UserLogin) (*response.UserLogin, error) {
	var userFound *entities.User
	userFound, err := s.userRepository.GetByEmail(ctx, req.Email)
	if userFound == nil || err != nil {
		return nil, errors.New(http.StatusUnauthorized, errors.StatusUnauthorizedCode, []string{"Invalid email or password"})
	}

	isValidPassword := crypto.ValidatePassword(req.Password, userFound.PasswordSalt, userFound.PasswordHash)
	if !isValidPassword {
		return nil, errors.New(http.StatusUnauthorized, errors.StatusUnauthorizedCode, []string{"Invalid email or password"})
	}

	accessToken, err := jwt.GenerateAccessToken(userFound.Code)
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, errors.StatusInternalServerErrorCode, []string{"Failed to generate access token"})
	}

	refreshToken, err := jwt.GenerateRefreshToken(userFound.Code)
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, errors.StatusInternalServerErrorCode, []string{"Failed to generate refresh token"})
	}

	return &response.UserLogin{
		User:         response.ToUserResponse(userFound),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil

}
