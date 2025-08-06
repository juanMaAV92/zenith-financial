package tests

import (
	"context"
	"net/http"
	"testing"

	"github.com/juanMaAV92/go-utils/errors"
	"github.com/juanMaAV92/go-utils/testhelpers"
	"github.com/juanMaAV92/zenith-financial/backend/cmd/handlers/users"
	"github.com/juanMaAV92/zenith-financial/backend/internal/repositories"
	userService "github.com/juanMaAV92/zenith-financial/backend/internal/services/users"
	"github.com/juanMaAV92/zenith-financial/backend/tests/helpers"
	"github.com/stretchr/testify/assert"
)

func (m *MockStore) Create(ctx context.Context, destination interface{}) error {
	args := m.Called(ctx, destination)
	return args.Error(0)
}

func Test_createUser(t *testing.T) {
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
			userService := userService.NewService(userRepository)
			handler := users.NewHandler(userService)

			err := handler.CreateUser(ctx)
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
