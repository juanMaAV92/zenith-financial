package auth

import (
	"context"
	libErrors "errors"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/juanMaAV92/go-utils/errors"
	jwtUtils "github.com/juanMaAV92/go-utils/jwt"
	"github.com/juanMaAV92/zenith-financial/backend/internal/domain/request"
	"github.com/juanMaAV92/zenith-financial/backend/internal/domain/response"
	"github.com/juanMaAV92/zenith-financial/backend/internal/entities"
	"github.com/juanMaAV92/zenith-financial/backend/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var jwtConfig = jwtUtils.JwtConfig{
	SecretKey:       "test_secret",
	Issuer:          "zenith-financial",
	AccessTokenTTL:  15 * time.Minute,
	RefreshTokenTTL: 24 * time.Hour,
	SigningMethod:   jwt.SigningMethodHS256,
}

func Test_Login(t *testing.T) {
	ctx := context.Background()
	jwtUtils.InitJWTConfig(&jwtConfig)
	var nilUser *entities.User
	testCases := []struct {
		name             string
		req              *request.UserLogin
		expectedResponse *response.UserLogin
		expectedError    *errors.ErrorResponse
		mockFunc         func(*mocks.UserRepository, *mocks.Cache, *mocks.Logger)
	}{
		{
			name: "valid login",
			req:  &request.UserLogin{Email: "test@example.com", Password: "12345677"},
			expectedResponse: &response.UserLogin{
				User: &response.User{
					Code:      uuid.MustParse("123e4567-e89b-12d3-a456-426614174000"),
					Email:     "test@example.com",
					UserName:  "testuser",
					Currency:  "USD",
					CreatedAt: time.Now(),
				},
				TokensResponse: &response.TokensResponse{
					AccessToken:  "access_token",
					RefreshToken: "refresh_token",
				},
			},
			mockFunc: func(repo *mocks.UserRepository, cache *mocks.Cache, logger *mocks.Logger) {
				repo.On("GetByEmail", mock.Anything, "test@example.com").Return(
					&entities.User{
						ID:           1,
						Code:         uuid.MustParse("123e4567-e89b-12d3-a456-426614174000"),
						Email:        "test@example.com",
						Username:     "testuser",
						PasswordHash: "$2a$12$mXOO6awNuioYxS2DLxmIZeQVadom64q3xP0MBiCHTljiKAwDLYLTO",
						PasswordSalt: "7da8aa7388bbe6e878064f084ac736a4",
						Currency:     "USD",
						CreatedAt:    time.Now(),
						UpdatedAt:    time.Now(),
					}, nil)
				cache.On("Set", mock.Anything, "user_refresh_token:123e4567-e89b-12d3-a456-426614174000", mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "Email not found",
			req:  &request.UserLogin{Email: "test1@example.com", Password: "12345677"},
			expectedError: &errors.ErrorResponse{
				HttpCode: http.StatusUnauthorized,
				Code:     errors.StatusUnauthorizedCode,
				Messages: []string{"Invalid email or password"},
			},
			mockFunc: func(repo *mocks.UserRepository, cache *mocks.Cache, logger *mocks.Logger) {
				repo.On("GetByEmail", mock.Anything, "test1@example.com").Return(nilUser, nil)
			},
		},
		{
			name: "Invalid password",
			req:  &request.UserLogin{Email: "test2@example.com", Password: "badPassword"},
			expectedError: &errors.ErrorResponse{
				HttpCode: http.StatusUnauthorized,
				Code:     errors.StatusUnauthorizedCode,
				Messages: []string{"Invalid email or password"},
			},
			mockFunc: func(repo *mocks.UserRepository, cache *mocks.Cache, logger *mocks.Logger) {
				repo.On("GetByEmail", mock.Anything, "test2@example.com").Return(
					&entities.User{
						ID:           1,
						Code:         uuid.MustParse("123e4567-e89b-12d3-a456-426614174000"),
						Email:        "test2@example.com",
						Username:     "testuser",
						PasswordHash: "$2a$12$mXOO6awNuioYxS2DLxmIZeQVadom64q3xP0MBiCHTljiKAwDLYLTO",
						PasswordSalt: "7da8aa7388bbe6e878064f084ac736a4",
						Currency:     "USD",
						CreatedAt:    time.Now(),
						UpdatedAt:    time.Now(),
					}, nil)
			},
		},
		{
			name: "valid login - cache error",
			req:  &request.UserLogin{Email: "test@example.com", Password: "12345677"},
			expectedResponse: &response.UserLogin{
				User: &response.User{
					Code:      uuid.MustParse("123e4567-e89b-12d3-a456-426614174000"),
					Email:     "test@example.com",
					UserName:  "testuser",
					Currency:  "USD",
					CreatedAt: time.Now(),
				},
				TokensResponse: &response.TokensResponse{
					AccessToken:  "access_token",
					RefreshToken: "refresh_token",
				},
			},
			mockFunc: func(repo *mocks.UserRepository, cache *mocks.Cache, logger *mocks.Logger) {
				repo.On("GetByEmail", mock.Anything, "test@example.com").Return(
					&entities.User{
						ID:           1,
						Code:         uuid.MustParse("123e4567-e89b-12d3-a456-426614174000"),
						Email:        "test@example.com",
						Username:     "testuser",
						PasswordHash: "$2a$12$mXOO6awNuioYxS2DLxmIZeQVadom64q3xP0MBiCHTljiKAwDLYLTO",
						PasswordSalt: "7da8aa7388bbe6e878064f084ac736a4",
						Currency:     "USD",
						CreatedAt:    time.Now(),
						UpdatedAt:    time.Now(),
					}, nil)
				cache.On("Set", mock.Anything, "user_refresh_token:123e4567-e89b-12d3-a456-426614174000", mock.Anything, mock.Anything).Return(libErrors.New("cache error"))
				logger.On("Error", mock.Anything, "login_cache_set", "error setting cache", mock.Anything, mock.Anything).Return()
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			userRepository := new(mocks.UserRepository)
			cache := new(mocks.Cache)
			logger := new(mocks.Logger)
			service := NewService(userRepository, cache, logger)

			tc.mockFunc(userRepository, cache, logger)

			resp, err := service.Login(ctx, tc.req)
			if tc.expectedError != nil {
				assert.Error(t, err)
				errorResponse, ok := err.(*errors.ErrorResponse)
				assert.True(t, ok)
				assert.Equal(t, tc.expectedError.HttpCode, errorResponse.ErrorHTTPCode())
				assert.Equal(t, tc.expectedError.Code, errorResponse.ErrorCode())
			} else {
				assert.Equal(t, tc.expectedResponse.Code, resp.Code)
				assert.Equal(t, tc.expectedResponse.User.Code, resp.User.Code)
				assert.Equal(t, tc.expectedResponse.User.Email, resp.User.Email)
				assert.Equal(t, tc.expectedResponse.User.UserName, resp.User.UserName)
				assert.Equal(t, tc.expectedResponse.User.Currency, resp.User.Currency)
				assert.NotEqual(t, "", resp.TokensResponse.AccessToken)
				assert.NotEqual(t, "", resp.TokensResponse.RefreshToken)
			}
			cache.AssertExpectations(t)
			userRepository.AssertExpectations(t)
			logger.AssertExpectations(t)

		})
	}

}

func Test_Logout(t *testing.T) {
	ctx := context.Background()
	userCode := uuid.MustParse("123e4567-e89b-12d3-a456-426614174000")

	jwtUtils.InitJWTConfig(&jwtConfig)
	token, err := jwtUtils.GenerateAccessToken(userCode)
	if err != nil {
		t.Fatalf("failed to generate token: %v", err)
	}

	testCases := []struct {
		name          string
		token         string
		mockFunc      func(*mocks.UserRepository, *mocks.Cache, *mocks.Logger)
		expectedError *errors.ErrorResponse
	}{
		{
			name:  "valid logout",
			token: token,
			mockFunc: func(repo *mocks.UserRepository, cache *mocks.Cache, logger *mocks.Logger) {
				cache.On("Delete", mock.Anything, "user_refresh_token:"+userCode.String()).Return(nil)
			},
		},
		{
			name:  "invalid token",
			token: "invalid_token",
			expectedError: &errors.ErrorResponse{
				HttpCode: http.StatusUnauthorized,
				Code:     errors.StatusUnauthorizedCode,
				Messages: []string{"Invalid token"},
			},
			mockFunc: func(repo *mocks.UserRepository, cache *mocks.Cache, logger *mocks.Logger) {
			},
		},
		{
			name:  "failed delete from cache",
			token: token,
			mockFunc: func(repo *mocks.UserRepository, cache *mocks.Cache, logger *mocks.Logger) {
				cache.On("Delete", mock.Anything, "user_refresh_token:"+userCode.String()).Return(libErrors.New("cache delete error"))
				logger.On("Error", mock.Anything, "logout_cache_delete", "error deleting cache", mock.Anything, mock.Anything).Return()
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			userRepository := new(mocks.UserRepository)
			cache := new(mocks.Cache)
			logger := new(mocks.Logger)
			service := NewService(userRepository, cache, logger)

			tc.mockFunc(userRepository, cache, logger)
			err = service.Logout(ctx, tc.token)
			if tc.expectedError != nil {
				assert.Error(t, err)
				errorResponse, ok := err.(*errors.ErrorResponse)
				assert.True(t, ok)
				assert.Equal(t, tc.expectedError.HttpCode, errorResponse.ErrorHTTPCode())
				assert.Equal(t, tc.expectedError.Code, errorResponse.ErrorCode())
			}
			cache.AssertExpectations(t)
			userRepository.AssertExpectations(t)
			logger.AssertExpectations(t)
		})
	}

}

func Test_RefreshToken(t *testing.T) {
	ctx := context.Background()
	userCode := uuid.MustParse("123e4567-e89b-12d3-a456-426614174000")

	jwtUtils.InitJWTConfig(&jwtConfig)
	accessToken, err := jwtUtils.GenerateAccessToken(userCode)
	if err != nil {
		t.Fatalf("failed to generate token: %v", err)
	}

	refreshToken, err := jwtUtils.GenerateRefreshToken(userCode)
	if err != nil {
		t.Fatalf("failed to generate token: %v", err)
	}

	testCases := []struct {
		name             string
		token            string
		expectedResponse *response.TokensResponse
		expectedError    *errors.ErrorResponse
		mockFunc         func(*mocks.UserRepository, *mocks.Cache, *mocks.Logger)
	}{
		{
			name:  "error parsing token",
			token: "invalid_token",
			expectedError: &errors.ErrorResponse{
				HttpCode: http.StatusUnauthorized,
				Code:     errors.StatusUnauthorizedCode,
				Messages: []string{"Invalid refreshtoken"},
			},
			mockFunc: func(repo *mocks.UserRepository, cache *mocks.Cache, logger *mocks.Logger) {
			},
		},
		{
			name:  "receive an access token",
			token: accessToken,
			expectedError: &errors.ErrorResponse{
				HttpCode: http.StatusUnauthorized,
				Code:     errors.StatusUnauthorizedCode,
				Messages: []string{"Invalid refresh token type"},
			},
			mockFunc: func(repo *mocks.UserRepository, cache *mocks.Cache, logger *mocks.Logger) {
				cache.On("Delete", mock.Anything, fmt.Sprintf("user_refresh_token:%s", userCode.String())).Return(nil)
				logger.On("Error", mock.Anything, "refresh_token_invalid_type", mock.Anything, mock.Anything, mock.Anything).Return()
			},
		},
		{
			name:  "success refresh token",
			token: refreshToken,
			expectedResponse: &response.TokensResponse{
				AccessToken:  accessToken,
				RefreshToken: refreshToken,
			},
			mockFunc: func(repo *mocks.UserRepository, cache *mocks.Cache, logger *mocks.Logger) {
				cache.On("Set", mock.Anything, fmt.Sprintf("user_refresh_token:%s", userCode.String()), mock.Anything, mock.Anything).Return(nil)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			userRepository := new(mocks.UserRepository)
			cache := new(mocks.Cache)
			logger := new(mocks.Logger)
			service := NewService(userRepository, cache, logger)

			tc.mockFunc(userRepository, cache, logger)
			response, err := service.RefreshToken(ctx, tc.token)
			if tc.expectedError != nil {
				assert.Error(t, err)
				errorResponse, ok := err.(*errors.ErrorResponse)
				assert.True(t, ok)
				assert.Equal(t, tc.expectedError.HttpCode, errorResponse.ErrorHTTPCode())
				assert.Equal(t, tc.expectedError.Code, errorResponse.ErrorCode())
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, response)
				assert.NotEqual(t, "", response.AccessToken)
				assert.NotEqual(t, "", response.RefreshToken)
			}
			cache.AssertExpectations(t)
			userRepository.AssertExpectations(t)
			logger.AssertExpectations(t)
		})
	}
}
