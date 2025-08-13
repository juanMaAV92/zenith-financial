package tests

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	utilCache "github.com/juanMaAV92/go-utils/cache"
	"github.com/juanMaAV92/go-utils/errors"
	jwtUtils "github.com/juanMaAV92/go-utils/jwt"
	"github.com/juanMaAV92/go-utils/testhelpers"
	"github.com/juanMaAV92/zenith-financial/backend/cmd/handlers/auth"
	"github.com/juanMaAV92/zenith-financial/backend/internal/domain/response"
	"github.com/juanMaAV92/zenith-financial/backend/internal/entities"
	"github.com/juanMaAV92/zenith-financial/backend/internal/repositories"
	authService "github.com/juanMaAV92/zenith-financial/backend/internal/services/auth"
	"github.com/juanMaAV92/zenith-financial/backend/tests/helpers"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockStore struct {
	mock.Mock
}

type MockCache struct {
	mock.Mock
}

func (m *MockCache) Set(ctx context.Context, key string, value interface{}, opts ...utilCache.SetOption) error {
	args := m.Called(ctx, key, value, opts)
	return args.Error(0)
}

func (m *MockStore) FindOne(ctx context.Context, destination interface{}, conditions interface{}) (bool, error) {
	args := m.Called(ctx, destination, conditions)
	return args.Get(0).(bool), args.Error(1)
}

func Test_login(t *testing.T) {
	path := "/users/login"
	cases := []testhelpers.HttpTestCase{
		{
			TestName: "Bad Request - Invalid JSON",
			Request: testhelpers.TestRequest{
				Method: "POST",
				Url:    path,
			},
			RequestBody: map[string]interface{}{
				"Email": true,
			},
			ExpectError: &errors.ErrorResponse{
				HttpCode: http.StatusBadRequest,
				Code:     errors.StatusBadRequestCode,
				Messages: []string{"Invalid request body"},
			},
		},
		{
			TestName: "success - Valid Credentials",
			Request: testhelpers.TestRequest{
				Method: "POST",
				Url:    path,
			},
			RequestBody: map[string]interface{}{
				"email":    "test@test.com",
				"password": "12345678",
			},
			ExpectError: nil,
			Response: testhelpers.ExpectedResponse{
				Status: http.StatusOK,
				Body: testhelpers.ToJSONString(&response.User{
					UserName: "juan",
					Email:    "test@test.com",
					Code:     uuid.MustParse("123e4567-e89b-12d3-a456-426614174000"),
				}),
			},
			MockFunc: func(s *echo.Echo, c echo.Context) {
				mockStore := c.Get("mockStore").(*MockStore)
				mockStore.On("FindOne",
					mock.Anything,
					mock.Anything,
					map[string]interface{}{"email": "test@test.com"}).Return(true, nil).Run(func(args mock.Arguments) {
					user := args.Get(1).(*entities.User)
					user.Code = uuid.MustParse("123e4567-e89b-12d3-a456-426614174000")
					user.Username = "juan"
					user.Email = "test@test.com"
					user.PasswordSalt = "bfd7d9e1a94ac31e"
					user.PasswordHash = "$2a$12$WjCNyUfAhbxYO.PRBGaGc.CHIx/1OVvqh7JkPf9CWWgppKW1cgxv2"
				})
				mockCache := c.Get("mockCache").(*MockCache)
				mockCache.On("Set",
					mock.Anything,
					"user_refresh_token:123e4567-e89b-12d3-a456-426614174000",
					mock.Anything,
					mock.AnythingOfType("[]cache.SetOption")).Return(nil)
			},
		},
		{
			TestName: "Unauthorized - Invalid Credentials",
			Request: testhelpers.TestRequest{
				Method: "POST",
				Url:    path,
			},
			RequestBody: map[string]interface{}{
				"email":    "test@test.com",
				"password": "abcdef",
			},
			ExpectError: &errors.ErrorResponse{
				HttpCode: http.StatusUnauthorized,
				Code:     "UNAUTHORIZED",
				Messages: []string{"Invalid email or password"},
			},
			MockFunc: func(s *echo.Echo, c echo.Context) {
				mockStore := c.Get("mockStore").(*MockStore)
				mockStore.On("FindOne",
					mock.Anything,
					mock.Anything,
					map[string]interface{}{"email": "test@test.com"}).Return(true, nil).Run(func(args mock.Arguments) {
					user := args.Get(1).(*entities.User)
					user.Code = uuid.MustParse("123e4567-e89b-12d3-a456-426614174000")
					user.Username = "juan"
					user.Email = "test@test.com"
					user.PasswordSalt = "bfd7d9e1a94ac31e"
					user.PasswordHash = "$2a$12$WjCNyUfAhbxYO.PRBGaGc.CHIx/1OVvqh7JkPf9CWWgppKW1cgxv2"
				})
				mockCache := c.Get("mockCache").(*MockCache)
				mockCache.On("Set",
					mock.Anything,
					"user_refresh_token:123e4567-e89b-12d3-a456-426614174000",
					mock.Anything,
					mock.AnythingOfType("[]cache.SetOption")).Return(nil)
			},
		},
	}

	app := helpers.NewTestServer()
	jwtConfig := jwtUtils.JwtConfig{
		SecretKey:       "test_secret",
		Issuer:          "zenith-financial",
		AccessTokenTTL:  15 * time.Minute,
		RefreshTokenTTL: 24 * time.Hour,
		SigningMethod:   jwt.SigningMethodHS256,
	}
	jwtUtils.InitJWTConfig(&jwtConfig)

	for _, test := range cases {
		t.Run(test.TestName, func(t *testing.T) {
			ctx, recorder := testhelpers.PrepareContextFormTestCase(app.Server.Echo, test)

			mockStore := new(MockStore)
			MockCache := new(MockCache)
			ctx.Set("mockStore", mockStore)
			ctx.Set("mockCache", MockCache)

			if test.MockFunc != nil {
				test.MockFunc(app.Server.Echo, ctx)
			}

			userRepository := repositories.NewUserRepository(mockStore)
			authService := authService.NewService(userRepository, MockCache, app.Logger)
			handler := auth.NewHandler(authService)

			err := handler.Login(ctx)
			if test.ExpectError != nil {
				assert.Error(t, err)
				errorResponse, ok := err.(*errors.ErrorResponse)
				assert.True(t, ok)
				assert.Equal(t, test.ExpectError.HttpCode, errorResponse.ErrorHTTPCode())
				assert.Equal(t, test.ExpectError.Code, errorResponse.ErrorCode())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.Response.Status, recorder.Code)

				if test.Response.Body == nil {
					assert.Empty(t, recorder.Body.String())
				} else {
					UserLogin := &response.UserLogin{}
					err := json.Unmarshal(recorder.Body.Bytes(), UserLogin)
					assert.NoError(t, err)
					assert.JSONEq(t, *test.Response.Body, *testhelpers.ToJSONString(&UserLogin.User))
				}
			}
			mockStore.AssertExpectations(t)
		})
	}
}
