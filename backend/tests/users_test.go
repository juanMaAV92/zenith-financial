package tests

import (
	"context"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/juanMaAV92/go-utils/errors"
	"github.com/juanMaAV92/go-utils/testhelpers"
	"github.com/juanMaAV92/zenith-financial/backend/cmd/handlers/users"
	"github.com/juanMaAV92/zenith-financial/backend/internal/domain/response"
	"github.com/juanMaAV92/zenith-financial/backend/internal/entities"
	"github.com/juanMaAV92/zenith-financial/backend/internal/repositories"
	usersService "github.com/juanMaAV92/zenith-financial/backend/internal/services/users"
	"github.com/juanMaAV92/zenith-financial/backend/tests/helpers"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockStore struct {
	mock.Mock
}

func (m *MockStore) Create(ctx context.Context, destination interface{}) error {
	args := m.Called(ctx, destination)
	return args.Error(0)
}

func (m *MockStore) FindOne(ctx context.Context, destination interface{}, conditions interface{}) (bool, error) {
	args := m.Called(ctx, destination, conditions)
	return args.Get(0).(bool), args.Error(1)
}

func Test_UsersValidateCredentials(t *testing.T) {
	path := "/users/register"
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
			},
		},
	}

	app := helpers.NewTestServer()

	for _, test := range cases {
		t.Run(test.TestName, func(t *testing.T) {
			ctx, recorder := testhelpers.PrepareContextFormTestCase(app.Server.Echo, test)

			mockStore := new(MockStore)
			ctx.Set("mockStore", mockStore)

			if test.MockFunc != nil {
				test.MockFunc(app.Server.Echo, ctx)
			}

			userRepository := repositories.NewUserRepository(mockStore)
			userService := usersService.NewService(userRepository)
			handler := users.NewHandler(userService)

			err := handler.ValidateCredentials(ctx)
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
					assert.JSONEq(t, *test.Response.Body, recorder.Body.String())
				}
			}
			mockStore.AssertExpectations(t)
		})
	}
}
