package auth

import (
	"context"
	"fmt"
	"net/http"
	"time"

	utilCache "github.com/juanMaAV92/go-utils/cache"
	"github.com/juanMaAV92/go-utils/errors"
	"github.com/juanMaAV92/go-utils/jwt"
	"github.com/juanMaAV92/go-utils/log"
	"github.com/juanMaAV92/zenith-financial/backend/internal/domain/request"
	"github.com/juanMaAV92/zenith-financial/backend/internal/domain/response"
	"github.com/juanMaAV92/zenith-financial/backend/internal/entities"
	"github.com/juanMaAV92/zenith-financial/backend/utils/crypto"
)

type userRepository interface {
	GetByEmail(ctx context.Context, email string) (*entities.User, error)
}

type cache interface {
	Set(ctx context.Context, key string, value interface{}, opts ...utilCache.SetOption) error
}

type service struct {
	userRepository userRepository
	cache          cache
	logger         log.Logger
}

func NewService(userRepo userRepository, cache cache, logger log.Logger) *service {
	return &service{
		userRepository: userRepo,
		cache:          cache,
		logger:         logger,
	}
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

	key := fmt.Sprintf("user_refresh_token:%s", userFound.Code)
	if err := s.cache.Set(ctx, key, refreshToken, utilCache.WithTTL(7*24*time.Hour)); err != nil {
		s.logger.Error(ctx, "login_cache_set", "error setting cache", log.Field("user_code", userFound.Code), log.Field("error", err))
	}

	return &response.UserLogin{
		User:         response.ToUserResponse(userFound),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil

}
