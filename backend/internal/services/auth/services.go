package auth

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
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
	Delete(ctx context.Context, key string) error
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
		User: response.ToUserResponse(userFound),
		TokensResponse: &response.TokensResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}, nil

}

func (s *service) Logout(ctx context.Context, authHeader string) error {

	claims, _, err := jwt.ParseClaims(authHeader)
	if err != nil {
		return errors.New(http.StatusUnauthorized, errors.StatusUnauthorizedCode, []string{"Invalid token"})
	}

	key := fmt.Sprintf("user_refresh_token:%s", claims["user_code"])
	if err := s.cache.Delete(ctx, key); err != nil {
		s.logger.Error(ctx, "logout_cache_delete", "error deleting cache", log.Field("user_code", claims["user_code"]), log.Field("error", err))
	}

	return nil
}

func (s *service) RefreshToken(ctx context.Context, refreshToken string) (*response.TokensResponse, error) {
	claims, isValid, err := jwt.ParseClaims(refreshToken)
	if err != nil {
		return nil, errors.New(http.StatusUnauthorized, errors.StatusUnauthorizedCode, []string{"Invalid refresh token"})
	}

	userCode := claims["user_code"].(string)

	if claims["type"] != "refresh" || !isValid {
		if err := s.cache.Delete(ctx, fmt.Sprintf("user_refresh_token:%s", userCode)); err != nil {
			s.logger.Error(ctx, "refresh_token_cache_delete", "error deleting cache", log.Field("user_code", userCode), log.Field("error", err))
		}
		s.logger.Error(ctx, "refresh_token_invalid_type", "expected refresh token type", log.Field("type", claims["type"]))
		return nil, errors.New(http.StatusUnauthorized, errors.StatusUnauthorizedCode, []string{"Invalid refresh token type"})
	}

	user, err := uuid.Parse(userCode)
	if err != nil {
		return nil, errors.New(http.StatusUnauthorized, errors.StatusUnauthorizedCode, []string{"Invalid user code in refresh token"})
	}

	newAccessToken, err := jwt.GenerateAccessToken(user)
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, errors.StatusInternalServerErrorCode, []string{"Failed to generate new access token"})
	}

	newRefreshToken, err := jwt.GenerateRefreshToken(user)
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, errors.StatusInternalServerErrorCode, []string{"Failed to generate new refresh token"})
	}

	key := fmt.Sprintf("user_refresh_token:%s", userCode)
	if err := s.cache.Set(ctx, key, newRefreshToken, utilCache.WithTTL(7*24*time.Hour)); err != nil {
		s.logger.Error(ctx, "refresh_token_cache_set", "error setting cache", log.Field("user_code", userCode), log.Field("error", err))
	}

	return &response.TokensResponse{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
	}, nil
}
