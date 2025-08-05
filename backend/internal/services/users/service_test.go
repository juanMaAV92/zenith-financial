package users

import (
	"context"
	"net/http"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/juanMaAV92/go-utils/errors"
	"github.com/juanMaAV92/zenith-financial/backend/internal/domain/request"
	"github.com/juanMaAV92/zenith-financial/backend/internal/domain/response"
	"github.com/juanMaAV92/zenith-financial/backend/internal/entities"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Create(ctx context.Context, user *entities.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}
func (m *MockRepository) GetByEmail(ctx context.Context, email string) (*entities.User, error) {
	args := m.Called(ctx, email)
	if user, ok := args.Get(0).(*entities.User); ok {
		return user, args.Error(1)
	}
	return nil, args.Error(1)
}

func Test_CreateUser(t *testing.T) {
	ctx := context.Background()

	testCase := []struct {
		name             string
		request          *request.CreateUser
		expectedResponse *response.User
		expectError      error
		mockFunc         func(*MockRepository)
	}{
		{
			name: "user repository returns error",
			request: &request.CreateUser{
				UserName: "testuser",
				Email:    "error@mail.com",
				Password: "password123",
				Currency: "USD",
			},
			expectedResponse: nil,
			expectError: errors.New(
				http.StatusInternalServerError,
				"ERROR",
				[]string{"Unable to process request"},
			),
			mockFunc: func(repo *MockRepository) {
				repo.On("GetByEmail",
					ctx,
					mock.Anything,
				).Return(nil, errors.New(
					http.StatusInternalServerError,
					"ERROR",
					[]string{"Unable to process request"},
				))
			},
		},
		{
			name: "user already exists error 409",
			request: &request.CreateUser{
				UserName: "testuser",
				Email:    "user_exists@mail.com",
				Password: "password123",
			},
			expectedResponse: nil,
			expectError: errors.New(
				http.StatusConflict,
				"USER_EXISTS",
				[]string{"User already exists"},
			),
			mockFunc: func(repo *MockRepository) {
				repo.On("GetByEmail",
					ctx,
					"user_exists@mail.com",
				).Return(&entities.User{}, nil)
			},
		},
		{
			name: "error creating user",
			request: &request.CreateUser{
				UserName: "testuser",
				Email:    "test@mail.com",
				Password: "password123",
				Currency: "USD",
			},
			expectedResponse: nil,
			expectError: errors.New(
				http.StatusInternalServerError,
				"CREATE_USER_ERROR",
				[]string{"Unable to create user"},
			),
			mockFunc: func(repo *MockRepository) {
				repo.On("GetByEmail",
					ctx,
					"test@mail.com",
				).Return(nil, nil)
				repo.On("Create",
					ctx,
					mock.AnythingOfType("*entities.User"),
				).Return(errors.New(
					http.StatusInternalServerError,
					"CREATE_USER_ERROR",
					[]string{"Unable to create user"},
				))
			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(MockRepository)
			tc.mockFunc(mockRepo)

			svc := NewService(mockRepo)
			response, err := svc.CreateUser(ctx, tc.request)

			if tc.expectError != nil {
				assert.Equal(t, tc.expectError, err)
			} else {
				if err != nil {
					t.Fatalf("Error creating user: %v", err)
				}
				assert.Equal(t, tc.expectedResponse, response)
			}
		})
	}
}
